package http

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"battlebit_admin_panel/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	errors "github.com/pludderio/errs"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type chatLog struct {
	SteamID64 string    `json:"steamid_64"`
	Message   string    `json:"message"`
	UserName  string    `json:"username"`
	Timestamp time.Time `json:"timestamp"`
}

type chatListResponse struct {
	Logs       []chatLog `json:"logs"`
	PrevPageID int       `json:"prev_page_id"`
	NextPageID int       `json:"next_page_id"`
}

func (s *Server) ListChat(c *gin.Context) {
	const pageLimit = 50
	curPos := 0
	nextFrom := c.Query("next_from")
	prevFrom := c.Query("prev_from")
	playerID := c.Query("player_id")
	var err error
	if nextFrom != "" {
		curPos, err = strconv.Atoi(nextFrom)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}
	if prevFrom != "" {
		curPos, err = strconv.Atoi(prevFrom)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	queryFilters := make([]qm.QueryMod, 0)
	if curPos != 0 {
		if nextFrom != "" {
			queryFilters = append(queryFilters, qm.Where(model.ChatLogColumns.ID+" <= ?", curPos), qm.OrderBy(model.ChatLogColumns.ID+" DESC"))
		} else {
			queryFilters = append(queryFilters, qm.Where(model.ChatLogColumns.ID+" >= ?", curPos), qm.OrderBy(model.ChatLogColumns.ID+" ASC"))
		}
	} else {
		queryFilters = append(queryFilters, qm.OrderBy(model.ChatLogColumns.ID+" DESC"))
	}
	var filteredPlayer *model.Player
	if playerID != "" {
		filteredPlayer, err = model.Players(qm.Where(model.PlayerColumns.SteamID+" = ?", playerID)).One(boil.GetDB())
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSONP(http.StatusOK, []chatLog{})
				return
			}
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		queryFilters = append(queryFilters, qm.Where(model.ChatLogColumns.PlayerID+" = ?", filteredPlayer.ID))
	}
	queryFilters = append(queryFilters, qm.Limit(pageLimit), qm.Load("Player"))
	chatLogs, err := model.ChatLogs(queryFilters...).All(boil.GetDB())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSONP(http.StatusOK, []chatLog{})
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var logs []chatLog
	lastID := 0
	firstId := -1
	sort.Slice(chatLogs, func(i, j int) bool {
		return chatLogs[i].ID > chatLogs[j].ID
	})
	for _, c := range chatLogs {
		logs = append(logs, chatLog{
			Message:   c.Message,
			UserName:  c.R.Player.Name,
			SteamID64: fmt.Sprintf("%d", c.R.Player.SteamID),
			Timestamp: c.Timestamp,
		})
		lastID = c.ID
		if firstId == -1 {
			firstId = c.ID
		}
	}
	nextFilters := []qm.QueryMod{qm.Where(model.ChatLogColumns.ID+" < ?", lastID), qm.OrderBy(model.ChatLogColumns.ID + " DESC"), qm.Limit(1)}
	prevFilters := []qm.QueryMod{qm.Where(model.ChatLogColumns.ID+" > ?", firstId), qm.OrderBy(model.ChatLogColumns.ID + " ASC"), qm.Limit(1)}
	if filteredPlayer != nil {
		playerFilter := qm.Where(model.ChatLogColumns.PlayerID+" = ?", filteredPlayer.ID)
		nextFilters = append(nextFilters, playerFilter)
		prevFilters = append(prevFilters, playerFilter)
	}

	nextPage, err := model.ChatLogs(nextFilters...).One(boil.GetDB())
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	previousPage, err := model.ChatLogs(prevFilters...).One(boil.GetDB())
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	nextPageID := -1
	if nextPage != nil {
		nextPageID = nextPage.ID
	}
	previousPageID := -1
	if previousPage != nil {
		previousPageID = previousPage.ID
	}
	c.JSONP(http.StatusOK, chatListResponse{
		Logs:       logs,
		PrevPageID: previousPageID,
		NextPageID: nextPageID,
	})
}

func (s *Server) ChatPage(c *gin.Context) {
	session := sessions.Default(c)
	steamID64 := session.Get("steamid_64")
	if steamID64 == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/steam/begin")
		return
	}
	admin, err := s.GetAdmin(steamID64.(int64))
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to check Steam ID")
		return
	}
	if admin == nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	logrus.Infof("User %s is logged in, rendering the home page", admin.Name)
	bytes, err := os.ReadFile("assets/chatlog.html")
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Err(err))
		return
	}
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, string(bytes))
}

func (s *Server) HomePage(c *gin.Context) {
	session := sessions.Default(c)
	steamID64 := session.Get("steamid_64")
	if steamID64 == nil {
		c.Redirect(http.StatusTemporaryRedirect, "/auth/steam/begin")
		return
	}
	admin, err := s.GetAdmin(steamID64.(int64))
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to check Steam ID")
		return
	}
	if admin == nil {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return
	}
	logrus.Infof("User %s is logged in, rendering the home page", admin.Name)
	bytes, err := os.ReadFile("assets/index.html")
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Err(err))
		return
	}
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, string(bytes))
}

func (s *Server) recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(500, gin.H{
		"title": "Error",
		"err":   err,
	})
}

func (s *Server) errorHandler(c *gin.Context) {
	c.Next()
	err := c.Errors.Last()
	if err == nil {
		return
	}
	logrus.Errorln(errors.FullTrace(err))
	c.String(-1, err.Error())
}
