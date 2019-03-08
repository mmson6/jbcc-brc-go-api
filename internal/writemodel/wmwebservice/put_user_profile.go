package wmwebservice

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/jbcc/brc-api/internal/readmodel/rmrepository"
	"github.com/jbcc/brc-api/internal/webresponse"
	"github.com/jbcc/brc-api/internal/writemodel/wmrepository"
	"github.com/jbcc/brc-api/pkg/brcapiv1"
	"github.com/jbcc/brc-api/pkg/brchttpv1/commandbody"
	"github.com/jbcc/brc-api/pkg/logger"
)

type PutUserProfileResponse struct {
	DisplayName string `json:"displayName"`
	Group       string `json:"group"`
	ID          string `json:"id"`
}

func PutUserProfile(w http.ResponseWriter, r *http.Request) {
	// Common setup
	ctx := r.Context()
	log := logger.Current(ctx).WithFields(logrus.Fields{
		"func":    "PutUserProfile",
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
	bodyRef, err := parseRequestBodyForCreateUserProfile(ctx, body)
	if err != nil {
		log.WithError(err).Error("unable to parse request body")
		webresponse.WriteErrorJSON(ctx, w, err)
		return
	}

	userProfile := bodyRef.Data

	// Generate 6-digit randome ID for the new user and write to DB
	id := generateRandomID(ctx)
	displayName := userProfile.DisplayName

	// Check if the new user's display name and generated ID is unique
	uniqueDN, uniqueID, err := checkForUniqueIDAndDisplayName(ctx, id, displayName)
	if err != nil {
		webresponse.WriteErrorJSON(ctx, w, err)
		return
	} else if !uniqueDN || !uniqueID {
		if !uniqueDN {
			uErr := errors.New("duplicate display name already exists")
			webresponse.WriteErrorJSON(ctx, w, uErr)
			return
		}
		if !uniqueID {
			uErr := errors.New("duplicate ID generated. try again")
			webresponse.WriteErrorJSON(ctx, w, uErr)
			return
		}
	}

	// Write the new user's profile to DB
	err = createNewUser(ctx, id, userProfile)

	// Write initial record to DB for the new user
	err = createInitialRecord(ctx, id, userProfile)

	responseObj := PutUserProfileResponse{
		DisplayName: userProfile.DisplayName,
		Group:       userProfile.Group,
		ID:          id,
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

func createInitialRecord(ctx context.Context, id string, userProfile brcapiv1.UserProfile) error {
	repo := wmrepository.Build(ctx)

	return repo.CreateInitialRecordForNewUser(ctx, id, userProfile)
}

func createNewUser(ctx context.Context, id string, userProfile brcapiv1.UserProfile) error {
	repo := wmrepository.Build(ctx)

	return repo.CreateNewUser(ctx, id, userProfile)
}

func checkForUniqueIDAndDisplayName(ctx context.Context, id, displayName string) (bool, bool, error) {
	repo := rmrepository.Build(ctx)

	var uniqueDN bool
	var uniqueDNErr error
	var uniqueID bool
	var uniqueIDErr error
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		uniqueDN, uniqueDNErr = repo.CheckForUniqueDisplayName(ctx, displayName)
		wg.Done()
	}()
	go func() {
		uniqueID, uniqueIDErr = repo.CheckForUniqueID(ctx, id)
		wg.Done()
	}()
	wg.Wait()

	if uniqueDNErr != nil {
		logger.Current(ctx).Error("unable to check for unique display name")
		return false, false, uniqueDNErr
	} else if uniqueIDErr != nil {
		logger.Current(ctx).Error("unable to check for unique ID")
		return false, false, uniqueIDErr
	}

	return uniqueDN, uniqueID, nil
}

////////////////////////////////////////////////////////////////////////////////
// HELPER FUNCTIONS

func generateRandomID(ctx context.Context) string {
	// Randomly generate 6-digit int ID
	max := 1000000
	min := 100000
	rand.Seed(time.Now().UTC().UnixNano())
	intID := min + rand.Intn(max-min)

	id := strconv.Itoa(intID)
	return id
}

func parseRequestBodyForCreateUserProfile(ctx context.Context, body []byte) (*commandbody.CreateUserProfile, error) {
	var reqStruct commandbody.CreateUserProfile
	err := json.Unmarshal(body, &reqStruct)
	if err != nil {
		return nil, err
	}

	return &reqStruct, nil
}
