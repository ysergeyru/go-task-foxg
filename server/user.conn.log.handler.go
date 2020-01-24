package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/matryer/respond"

	"github.com/ysergeyru/go-task-foxg/pg"
	"github.com/ysergeyru/go-task-foxg/types"
	"github.com/ysergeyru/go-task-foxg/utils"
)

// handleUserLogDuplicatesCheck returns true if there are at least two matching ip addresses for given user IDs
func (s *Server) HandleUserLogDuplicatesCheck(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("handleUserLogDuplicatesCheck")
	muxFirstUserID := mux.Vars(r)["first_user_id"]
	muxSecondUserID := mux.Vars(r)["second_user_id"]

	firstUserID, err := strconv.Atoi(muxFirstUserID)
	if err != nil {
		err = fmt.Errorf("Invalid 'first_user_id' value: %s", muxFirstUserID)
		s.logger.Error(err)
		respond.With(w, r, http.StatusBadRequest, err)
	}
	secondUserID, err := strconv.Atoi(muxSecondUserID)
	if err != nil {
		err = fmt.Errorf("Invalid 'second_user_id' value: %s", muxSecondUserID)
		s.logger.Error(err)
		respond.With(w, r, http.StatusBadRequest, err)
	}
	// Get connections per user from Postgres
	usersConnections, err := pg.GetConnectionsByUserIDs(firstUserID, secondUserID)
	if err != nil {
		s.logger.Error(err)
		respond.With(w, r, http.StatusInternalServerError, err)
		return
	}

	if len(usersConnections) == 0 {
		s.logger.Debugf("No connections found for user IDs %d and %d", firstUserID, secondUserID)
		respond.With(w, r, http.StatusOK, types.DuplicateRespond{Duplicate: false})
		return
	}

	if isDuplicate := utils.CheckForDuplicateIP(usersConnections); !isDuplicate {
		s.logger.Debugf("No duplicates found for user IDs %d and %d", firstUserID, secondUserID)
		respond.With(w, r, http.StatusOK, types.DuplicateRespond{Duplicate: false})
		return
	} else {
		s.logger.Debugf("Found duplicate IPs for user IDs %d and %d", firstUserID, secondUserID)
		respond.With(w, r, http.StatusOK, types.DuplicateRespond{Duplicate: true})
	}
}
