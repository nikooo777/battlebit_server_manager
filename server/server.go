package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/steam"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Server is an instance of a peer server that houses the listener and store.
type Server struct {
	grp *sync.WaitGroup
}

// NewServer returns an initialized Server pointer.
func NewServer(stopGrp *sync.WaitGroup) *Server {
	return &Server{
		grp: stopGrp,
	}
}

// Start starts the server listener to handle connections.
func (s *Server) Start(address string) error {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(s.errorHandler)
	router.Use(nice.Recovery(s.recoveryHandler))
	store := cookie.NewStore([]byte(viper.GetString("cookie_secret")))
	gothic.Store = store
	goth.UseProviders(
		steam.New(viper.GetString("steam_api_key"), fmt.Sprintf("http://%s%s", viper.GetString("domain"), "/auth/steam/callback")),
	)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(sessions.Sessions("mysession", store), cors.New(config))

	router.GET("/", s.HomePage)
	router.GET("/chat", s.ChatPage)
	router.GET("/charts", s.ChartsPage)
	router.GET("/reports", s.ReportsPage)
	router.GET("/api/chat", s.ListChat)
	router.GET("/api/reports", s.ListReports)
	router.GET("/api/player_history", s.PlayerHistory)
	router.GET("/auth/:provider/begin", func(c *gin.Context) {
		fn := gothic.GetProviderName
		gothic.GetProviderName = func(req *http.Request) (string, error) {
			provider := c.Params.ByName("provider")
			if provider == "" {
				return fn(req)
			}
			return provider, nil
		}
		c.Param("provider")
		gothic.BeginAuthHandler(c.Writer, c.Request)
	})

	router.GET("/auth/steam/callback", func(c *gin.Context) {
		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			c.String(http.StatusBadRequest, "Failed to complete authentication")
			return
		}
		// Extract the steam ID of the user
		steamID := user.UserID

		// Convert the steamID to various formats
		steamID64, err := strconv.ParseInt(steamID, 10, 64)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to convert Steam ID to SteamID64")
			return
		}

		log.Println("Steam ID 64:", steamID64)
		admin, err := s.GetAdmin(steamID64)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to check Steam ID")
			return
		}
		if admin == nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			return
		}

		session := sessions.Default(c)
		// Save the steam ID in the session
		session.Set("steamid_64", steamID64)
		err = session.Save()
		if err != nil {
			c.String(http.StatusBadRequest, "Failed to save session")
			return
		}
		c.Redirect(http.StatusFound, "/")
	})

	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}
	log.Println("HTTP server listening on " + address)
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s", err)
		}
	}()
	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
	return nil
}

func (s *Server) listenForShutdown(listener *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)
	<-interruptChan
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := listener.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Infof("shut down!")
}
