package webresponse

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	// "github.com/jbcc/brc-api/internal/webmiddleware"
	"github.com/jbcc/brc-api/pkg/herror"
	"github.com/jbcc/brc-api/pkg/logger"
)

type acceptedBody struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	RequestID string `json:"requestId,omitempty"`
}

type errorBody struct {
	Err *herror.HorizonError `json:"error"`
}

func WriteJSON(ctx context.Context, w http.ResponseWriter, status int, val interface{}) {
	// Guard statements
	if val == nil {
		WriteEmpty(ctx, w, status)
		return
	}

	// Preparation

	log := logger.Current(ctx).WithFields(logrus.Fields{
		"action": "writejson-body",
		"scope":  "webresponse",
	})

	// Prepare body

	var body []byte
	var err error
	switch castVal := val.(type) {
	case []byte:
		body = castVal
	default:
		body, err = json.Marshal(val)
		if err != nil {
			log.WithError(err).Error("unable to covert response value to JSON")
			WriteErrorJSON(ctx, w, err)
			return
		}
	}

	// Write headers

	headers := w.Header()
	length := len(body)
	lengthStr := strconv.FormatInt(int64(length), 10)
	headers.Set("Content-Type", "application/json")
	headers.Set("Content-Length", lengthStr)
	w.WriteHeader(status)

	// Write body

	if _, err := w.Write(body); err != nil {
		log.Printf("unable to write response body: %s", err.Error())
	}
}

func WriteAcceptedJSON(ctx context.Context, w http.ResponseWriter) {
	val := acceptedBody{
		Status:  "accepted",
		Message: "Request accepted to be processed",
	}
	// if reqID, ok := webmiddleware.RequestID(ctx); ok {
	// 	val.RequestID = reqID
	// }

	WriteJSON(ctx, w, http.StatusAccepted, val)
}

func WriteErrorJSON(ctx context.Context, w http.ResponseWriter, err error) {
	herr := herror.Wrap(err,
		herror.Code("unexpected-error"),
		herror.Status(500),
	)
	writeHorizonErrorJSON(ctx, w, &herr)
}

// Writes the HorizonError to the HTTP response. The HorizonError must not be
// nil.
func writeHorizonErrorJSON(ctx context.Context, w http.ResponseWriter, herr *herror.HorizonError) {
	// Guard statements

	if herr == nil {
		panic("illegal call to WriteJSONHorizonError() with nil HorizonError")
	}

	// Preparation

	log := logger.Current(ctx).WithFields(logrus.Fields{
		"action": "writejson-error",
		"scope":  "webresponse",
	})

	// Prepare body

	// if reqID, ok := webmiddleware.RequestID(ctx); ok {
	// 	herr.Source.RequestID = &reqID
	// }

	bodyStruct := errorBody{herr}
	bodyBinary, err := json.Marshal(bodyStruct)
	if err != nil {
		log.WithError(err).Error("unable to convert response to JSON")
		w.WriteHeader(500)
		return
	}

	// Write headers

	headers := w.Header()
	length := len(bodyBinary)
	lengthStr := strconv.FormatInt(int64(length), 10)
	headers.Set("Content-Type", "application/json")
	headers.Set("Content-Length", lengthStr)
	headers.Set("Access-Control-Allow-Origin", "*")
	headers.Set("Access-Control-Allow-Credentials", "true")

	statusCode, _ := herr.HttpStatusCode()
	w.WriteHeader(statusCode)

	// Write body

	if _, err := w.Write(bodyBinary); err != nil {
		log.WithError(err).Error("unable to write response body")
		return
	}
}
