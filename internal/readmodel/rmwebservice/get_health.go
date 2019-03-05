package rmwebservice

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/jbcc/brc-api/internal/webresponse"
	"github.com/jbcc/brc-api/pkg/logger"
)

type GetHealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.Current(ctx).WithFields(logrus.Fields{
		"func":    "GetHealth",
		"package": "rmwebservice",
	})
	log.Info(`getting system health`)
	log.Info(`getting system health hahaha mike check`)

	version := "v1"

	responseObj := GetHealthResponse{
		Status:  "UP",
		Version: version,
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
