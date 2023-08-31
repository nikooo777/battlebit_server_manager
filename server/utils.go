package http

import (
	"database/sql"
	"strconv"
	"sync"

	"battlebit_admin_panel/model"

	errors "github.com/pludderio/errs"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var concurrentSteamIDCache = sync.Map{}

func (s *Server) GetAdmin(steamID int64) (*model.Admin, error) {
	cachedVal, found := concurrentSteamIDCache.Load(steamID)
	if found {
		return cachedVal.(*model.Admin), nil
	}

	admin, err := model.Admins(qm.Where(model.AdminColumns.SteamID+"=?", steamID)).One(boil.GetDB())
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, errors.Err(err)
	}

	concurrentSteamIDCache.Store(steamID, admin)
	return admin, nil
}

var ErrIntTooSmall = errors.Err("64-bit steamid int should be bigger than 76561197960265728")

// SteamInt64ToString takes a 64-bit integer and converts it to a steamid string format "STEAM_X:Y:Z"
// The argument must be bigger than 76561197960265728, or it will return an error.
func SteamInt64ToString(steamInt int64) (string, error) {
	if steamInt <= 76561197960265728 {
		return string(""), ErrIntTooSmall
	}

	steamInt = steamInt - 76561197960265728
	remainder := steamInt % 2
	steamInt = steamInt / 2
	return "STEAM_0:" + strconv.FormatInt(remainder, 10) + ":" + strconv.FormatInt(steamInt, 10), nil

}
