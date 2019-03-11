package rmwebservice

import (
	// "context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	// "github.com/jbcc/brc-api/internal/readmodel/rmrepository"
	"github.com/jbcc/brc-api/internal/webresponse"
	"github.com/jbcc/brc-api/pkg/logger"
)

type GetUserIdentityResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	UserID  string `json:"userId"`
}

func GetUserIdentity(w http.ResponseWriter, r *http.Request) {
	// Extract the request variables

	vars := mux.Vars(r)
	userID := vars["user_id"]

	// Common setup
	ctx := r.Context()
	log := logger.Current(ctx).WithFields(logrus.Fields{
		"func":    "GetUserIdentity",
		"package": "rmwebservice",
	})

	// Guard statements
	if userID == "" {
		err := errors.New("user ID not found")
		webresponse.WriteErrorJSON(ctx, w, err)
		return
	}

	// findUserInfoByUserID(ctx, userID)

	version := "v1"

	responseObj := GetUserIdentityResponse{
		Status:  "UP",
		Version: version,
		UserID:  userID,
	}

	jsonBin, err := json.Marshal(responseObj)
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

// func findUserInfoByUserID(ctx context.Context, userID string) error {
// 	repo := rmrepository.Build(ctx)
// 	_, err := repo.ReadMyAccount(ctx, userID)

// 	return err
// }
