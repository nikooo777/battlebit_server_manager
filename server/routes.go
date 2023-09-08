package http

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
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

type pagination struct {
	PrevPageID int `json:"prev_page_id"`
	NextPageID int `json:"next_page_id"`
}
type chatListResponse struct {
	Logs []chatLog `json:"logs"`
	pagination
}

type playerCount struct {
	Count int       `json:"count"`
	Date  time.Time `json:"date"`
}
type playerHistoryResponse struct {
	StartDate int64         `json:"start_date"`
	EndDate   int64         `json:"end_date"`
	History   []playerCount `json:"history"`
}

func (s *Server) PlayerHistory(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	startDate := time.Now().AddDate(0, 0, -1)
	endDate := time.Now()

	if startDateStr != "" {
		startDateInt, err := strconv.Atoi(startDateStr)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		startDate = time.Unix(int64(startDateInt), 0)
	}
	if endDateStr != "" {
		endDateInt, err := strconv.Atoi(endDateStr)
		if err != nil {
			_ = c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		endDate = time.Unix(int64(endDateInt), 0)
		if endDate.Before(startDate) {
			_ = c.AbortWithError(http.StatusBadRequest, fmt.Errorf("end date is before start date"))
			return
		}
		if endDate.After(time.Now()) {
			endDate = time.Now()
		}
	}

	statistics, err := model.Stats(qm.Where(model.StatColumns.CreatedAt+" >= ?", startDate), qm.Where(model.StatColumns.CreatedAt+" <= ?", endDate)).All(boil.GetDB())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSONP(http.StatusOK, playerHistoryResponse{
				StartDate: startDate.Unix(),
				EndDate:   endDate.Unix(),
				History:   []playerCount{},
			})
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	history := make([]playerCount, 0, len(statistics))
	for _, stat := range statistics {
		history = append(history, playerCount{
			Count: stat.PlayerCount,
			Date:  stat.CreatedAt,
		})
	}
	c.JSONP(http.StatusOK, playerHistoryResponse{
		StartDate: startDate.Unix(),
		EndDate:   endDate.Unix(),
		History:   history,
	})
}

type listReportResponse struct {
	Reports []report `json:"reports"`
	pagination
}

type reportPlayerDetails struct {
	SteamID64 string `json:"steamid_64"`
	Name      string `json:"name"`
}
type report struct {
	ReportedPlayer  reportPlayerDetails `json:"reported_player"`
	ReporterPlayer  reportPlayerDetails `json:"reporter_player"`
	Reason          string              `json:"reason"`
	ReportedAt      time.Time           `json:"reported_at"`
	RelatedChatLogs []string            `json:"related_chat_logs"`
	Status          string              `json:"status"`
}

func (s *Server) ListReports(c *gin.Context) {
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
			queryFilters = append(queryFilters, qm.Where(model.PlayerReportColumns.ID+" <= ?", curPos), qm.OrderBy(model.PlayerReportColumns.ID+" DESC"))
		} else {
			queryFilters = append(queryFilters, qm.Where(model.PlayerReportColumns.ID+" >= ?", curPos), qm.OrderBy(model.PlayerReportColumns.ID+" ASC"))
		}
	} else {
		queryFilters = append(queryFilters, qm.OrderBy(model.PlayerReportColumns.ID+" DESC"))
	}
	var reportedPlayerFilter *model.Player
	if playerID != "" {
		reportedPlayerFilter, err = model.Players(qm.Where(model.PlayerColumns.SteamID+" = ?", playerID)).One(boil.GetDB())
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				c.JSONP(http.StatusOK, listReportResponse{})
				return
			}
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		queryFilters = append(queryFilters, qm.Where(model.PlayerReportColumns.ReportedPlayerID+" = ?", reportedPlayerFilter.ID))
	}
	queryFilters = append(queryFilters, qm.Limit(pageLimit), qm.Load("ReportedPlayer"), qm.Load("Reporter"))
	reports, err := model.PlayerReports(queryFilters...).All(boil.GetDB())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSONP(http.StatusOK, listReportResponse{})
			return
		}
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	var reportsList []report

	lastID := 0
	firstId := -1
	sort.Slice(reports, func(i, j int) bool {
		return reports[i].ID > reports[j].ID
	})
	for _, r := range reports {
		relatedChat := make([]string, 0)
		if strings.Contains(r.Reason, "_Text") {
			chat, err := r.R.ReportedPlayer.ChatLogs(qm.Where(model.ChatLogColumns.Timestamp+" BETWEEN ? AND ?", r.Timestamp.Add(-time.Minute*5), r.Timestamp.Add(time.Minute*5))).All(boil.GetDB())
			if err != nil {
				if !errors.Is(err, sql.ErrNoRows) {
					_ = c.AbortWithError(http.StatusInternalServerError, err)
					return
				}
			}
			for _, c := range chat {
				relatedChat = append(relatedChat, c.Message)
			}

		}
		reportsList = append(reportsList, report{
			ReportedPlayer: reportPlayerDetails{
				SteamID64: fmt.Sprintf("%d", r.R.ReportedPlayer.SteamID),
				Name:      r.R.ReportedPlayer.Name,
			},
			ReporterPlayer: reportPlayerDetails{
				SteamID64: fmt.Sprintf("%d", r.R.Reporter.SteamID),
				Name:      r.R.Reporter.Name,
			},
			Reason:          r.Reason,
			ReportedAt:      r.Timestamp,
			RelatedChatLogs: relatedChat,
			Status:          r.Status,
		})
		lastID = r.ID
		if firstId == -1 {
			firstId = r.ID
		}
	}
	nextFilters := []qm.QueryMod{qm.Where(model.PlayerReportColumns.ID+" < ?", lastID), qm.OrderBy(model.PlayerReportColumns.ID + " DESC"), qm.Limit(1)}
	prevFilters := []qm.QueryMod{qm.Where(model.PlayerReportColumns.ID+" > ?", firstId), qm.OrderBy(model.PlayerReportColumns.ID + " ASC"), qm.Limit(1)}
	if reportedPlayerFilter != nil {
		playerFilter := qm.Where(model.PlayerReportColumns.ReportedPlayerID+" = ?", reportedPlayerFilter.ID)
		nextFilters = append(nextFilters, playerFilter)
		prevFilters = append(prevFilters, playerFilter)
	}

	nextPage, err := model.PlayerReports(nextFilters...).One(boil.GetDB())
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	previousPage, err := model.PlayerReports(prevFilters...).One(boil.GetDB())
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
	c.JSONP(http.StatusOK, listReportResponse{
		Reports: reportsList,
		pagination: pagination{
			PrevPageID: previousPageID,
			NextPageID: nextPageID,
		},
	})
}

func (s *Server) ListChat(c *gin.Context) {
	const pageLimit = 20
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
				c.JSONP(http.StatusOK, chatListResponse{})
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
		Logs: logs,
		pagination: pagination{
			PrevPageID: previousPageID,
			NextPageID: nextPageID,
		},
	})
}

func (s *Server) ChartsPage(c *gin.Context) {
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
	bytes, err := os.ReadFile("assets/player_history.html")
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Err(err))
		return
	}
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, string(bytes))
}

func (s *Server) ReportsPage(c *gin.Context) {
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
	bytes, err := os.ReadFile("assets/reports.html")
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, errors.Err(err))
		return
	}
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, string(bytes))
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
