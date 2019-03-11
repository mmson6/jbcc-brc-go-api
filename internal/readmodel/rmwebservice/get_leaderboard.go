package rmwebservice

import (
	"context"
	"encoding/json"
	// "errors"
	"net/http"
	"strconv"

	// "github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/jbcc/brc-api/internal/models"
	"github.com/jbcc/brc-api/internal/readmodel/rmrepository"
	"github.com/jbcc/brc-api/internal/webresponse"
	"github.com/jbcc/brc-api/pkg/logger"
)

type GetLeaderboardResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	UserID  string `json:"userId"`
}

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	// Common setup
	ctx := r.Context()
	log := logger.Current(ctx).WithFields(logrus.Fields{
		"func":    "GetLeaderboard",
		"package": "rmwebservice",
	})

	// Find user record by userID
	leaderboardRef, err := readBRCLeaderboard(ctx)
	if err != nil {
		log.WithError(err).Error("unable to find user record")
		webresponse.WriteErrorJSON(ctx, w, err)
		return
	}

	jsonBin, err := json.Marshal(leaderboardRef)
	if err != nil {
		log.WithError(err).Error("unable to convert response object to JSON")
		webresponse.WriteErrorJSON(ctx, w, err)
		return
	}

	// Send the response
	contentLengthInt := len(jsonBin)
	contentLength := strconv.FormatInt(int64(contentLengthInt), 10)

	w.Header().Set("Content-Length", contentLength)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(200)

	_, _ = w.Write(jsonBin)
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS

func readBRCLeaderboard(ctx context.Context) (*models.Leaderboard, error) {
	repo := rmrepository.Build(ctx)
	return repo.ReadLeaderboard(ctx)
}
