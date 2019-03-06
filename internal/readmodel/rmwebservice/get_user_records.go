package rmwebservice

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/jbcc/brc-api/internal/webresponse"
	"github.com/jbcc/brc-api/pkg/logger"
)

type GetUserRecordsResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	UserID  string `json:"userId"`
}

func GetUserRecords(w http.ResponseWriter, r *http.Request) {
	// Extract the request variables

	vars := mux.Vars(r)
	userID := vars["user_id"]

	ctx := r.Context()
	log := logger.Current(ctx).WithFields(logrus.Fields{
		"func":    "GetUserRecords",
		"package": "rmwebservice",
	})
	log.Info(`getting system health`)
	log.Info(`getting system health hahaha mike check`)

	version := "v1"

	responseObj := GetUserRecordsResponse{
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
	w.WriteHeader(200)

	_, _ = w.Write(jsonBin)
}
