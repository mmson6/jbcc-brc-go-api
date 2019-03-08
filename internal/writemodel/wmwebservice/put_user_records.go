package wmwebservice

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/jbcc/brc-api/internal/brc"
	"github.com/jbcc/brc-api/internal/models"
	"github.com/jbcc/brc-api/internal/readmodel/rmrepository"
	"github.com/jbcc/brc-api/internal/webresponse"
	"github.com/jbcc/brc-api/internal/writemodel/wmrepository"
	"github.com/jbcc/brc-api/pkg/brcapiv1"
	"github.com/jbcc/brc-api/pkg/brchttpv1/commandbody"
	"github.com/jbcc/brc-api/pkg/logger"
)

type PutUserRecordResponse struct {
	ChapterCount int    `json:"chapterCount,omitempty"`
	DisplayName  string `json:"displayName"`
	Group        string `json:"group"`
	ID           string `json:"id"`
	VerseCount   int    `json:"verseCount,omitempty"`
}

func PutUserRecords(w http.ResponseWriter, r *http.Request) {
	// Common setup
	ctx := r.Context()
	log := logger.Current(ctx).WithFields(logrus.Fields{
		"func":    "PutUserRecords",
		"package": "wmwebservice",
	})

	// Read  request body
	body, err := readBody(ctx, r)
	if err != nil {
		log.WithError(err).Error("unable to read request body")
		webresponse.WriteErrorJSON(ctx, w, err)
		return
	}

	// Parse request body
	bodyRef, err := parseRequestBodyForUpdateUserRecord(ctx, body)
	if err != nil {
		log.WithError(err).Error("unable to parse request body")
		webresponse.WriteErrorJSON(ctx, w, err)
		return
	}

	userRecord := bodyRef.Data
	updatedUserRecordRef, err := populateUserRecord(ctx, userRecord)
	if err != nil {
		log.WithError(err).Error("unable to populate user record")
		webresponse.WriteErrorJSON(ctx, w, err)
		return
	}

	err = updateUserRecord(ctx, *updatedUserRecordRef)
	if err != nil {
		log.WithError(err).Error("unable to update user record")
		webresponse.WriteErrorJSON(ctx, w, err)
		return
	}

	responseObj := PutUserRecordResponse{
		ChapterCount: updatedUserRecordRef.ChapterCount,
		DisplayName:  updatedUserRecordRef.DisplayName,
		Group:        updatedUserRecordRef.Group,
		ID:           updatedUserRecordRef.ID,
		VerseCount:   updatedUserRecordRef.VerseCount,
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
	// // Parse request body
	// profileRef, err := parseRequestBodyForOrgProfile(body)
	// if err != nil {
	// 	log.WithError(err).Error("unable to parse request body")
	// 	webresponse.WriteErrorJSON(ctx, w, err)
	// 	return
	// }

	// // Update the organization
	// uoOrgRef, err := populateOrganizationWithUpdatedProfile(ctx, orgID, *profileRef)
	// if err != nil {
	// 	log.WithError(err).Error("unable to read organization")
	// 	webresponse.WriteErrorJSON(ctx, w, err)
	// 	return
	// } else if uoOrgRef == nil {
	// 	log.WithError(err).Error("unable to find organization")
	// 	webresponse.WriteJSON(ctx, w, http.StatusNotFound, nil)
	// 	return
	// }

	// // Publish uo.organization.updated
	// if err = wmeventpublisher.PublishOrganizationUpdated(ctx, *uoOrgRef); err != nil {
	// 	log.WithError(err).Error("unable to publish uo.organization.updated event")
	// 	webresponse.WriteErrorJSON(ctx, w, err)
	// 	return
	// }
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS

func calculateChaptersAndVerses(ctx context.Context, record brcapiv1.Record) brc.RecordCalculationOutput {
	return brc.CalculateTotal(record)
}

func populateUserRecord(ctx context.Context, userRecord brcapiv1.UserRecord) (*models.UserRecord, error) {
	repo := rmrepository.Build(ctx)
	model, err := repo.ReadUserRecordByUserID(ctx, userRecord.ID)
	if err != nil {
		return nil, err
	}

	output := calculateChaptersAndVerses(ctx, userRecord.Record)

	updatedModel := models.UserRecord{
		ChapterCount: output.Chapters,
		DisplayName:  model.DisplayName,
		Group:        model.Group,
		ID:           userRecord.ID,
		Record:       userRecord.Record,
		VerseCount:   output.Verses,
	}

	return &updatedModel, nil
}

func updateUserRecord(ctx context.Context, userRecord models.UserRecord) error {
	repo := wmrepository.Build(ctx)

	err := repo.UpdateUserRecord(ctx, userRecord)
	if err != nil {
		return err
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// HELPER FUNCTIONS

func parseRequestBodyForUpdateUserRecord(ctx context.Context, body []byte) (*commandbody.UpdateUserRecord, error) {
	var reqStruct commandbody.UpdateUserRecord
	err := json.Unmarshal(body, &reqStruct)
	if err != nil {
		return nil, err
	}

	return &reqStruct, nil
}
