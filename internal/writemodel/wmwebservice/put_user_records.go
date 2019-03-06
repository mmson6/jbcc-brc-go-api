package wmwebservice

import (
	// "context"
	// "encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"

	// "code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/auth0"
	// authrepo "code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/authorization/repository"
	// "code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/constants"
	// "code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/events/uoevents"
	// "code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/uoerror"
	"github.com/jbcc/brc-api/internal/webresponse"
	// wmrepo "code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/wm/repository"
	// "code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/writemodel/wmeventpublisher"
	// "code.siemens.com/horizon/platform-verticals/user/uo-api-go/pkg/herror"
	"github.com/jbcc/brc-api/pkg/logger"
	// "code.siemens.com/horizon/platform-verticals/user/uo-api-go/pkg/uohttpv1/commandbody"
)

func PutUserRecords(w http.ResponseWriter, r *http.Request) {
	// Extract the request variables
	vars := mux.Vars(r)

	// Common setup
	ctx := r.Context()
	log := logger.Current(ctx).WithFields(logrus.Fields{
		"func":    "PutUserRecords",
		"package": "wmwebservice",
	})

	log.Info("Mike check vars in PutUserRecords: ", vars)

	// // Guard statements
	// if orgID == "" {
	// 	err := uoerror.NewNotFoundError()
	// 	log.WithError(err).Error("organization ID not found")
	// 	webresponse.WriteErrorJSON(ctx, w, err)
	// 	return
	// }

	// // Authorize request
	// if err := authorizeAccessForOrgProfileUpdate(ctx, orgID); err != nil {
	// 	log.WithError(err).Error("unable to authorize access")
	// 	webresponse.WriteErrorJSON(ctx, w, err)
	// 	return
	// }

	// Read  request body
	body, err := readBody(ctx, r)
	if err != nil {
		log.WithError(err).Error("unable to read request body")
		webresponse.WriteErrorJSON(ctx, w, err)
	}

	log.Info("mike test with data : ", body)
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

	webresponse.WriteAcceptedJSON(ctx, w)
}

////////////////////////////////////////////////////////////////////////////////
// HELPER FUNCTIONS

// func authorizeAccessForOrgProfileUpdate(ctx context.Context, orgID string) error {
// 	// Get the external ID of the user
// 	token, _ := auth0.LoadToken(ctx)
// 	externalID := token.Subject()

// 	repo := authrepo.Build(ctx)
// 	acl, err := repo.ReadAccessControlListByExternalID(ctx, externalID)
// 	if err != nil {
// 		return err
// 	} else if acl == nil {
// 		herr := uoerror.NewNotFoundError(
// 			herror.Detail("Unable to find access controls"),
// 		)
// 		return herr
// 	}

// 	permission := constants.PermissionOrganizationProfileUpdate
// 	systemAccess := acl.HasSystemAccess(permission)
// 	targetAccess := acl.HasTargetAccess(permission, uoevents.TypeOrganization, orgID)

// 	if !(systemAccess || targetAccess) {
// 		return uoerror.NewForbiddenError()
// 	}

// 	return nil
// }

// func organizationProfileEventModel(profile commandbody.UpdateOrganizationProfile) uoevents.OrganizationProfile {
// 	var evtModel uoevents.OrganizationProfile
// 	mapstructure.Decode(profile.Data, &evtModel)
// 	return evtModel
// }

// func parseRequestBodyForOrgProfile(body []byte) (*commandbody.UpdateOrganizationProfile, error) {
// 	var reqStruct commandbody.UpdateOrganizationProfile
// 	err := json.Unmarshal(body, &reqStruct)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &reqStruct, nil
// }

// func populateOrganizationWithUpdatedProfile(
// 	ctx context.Context,
// 	orgID string,
// 	profile commandbody.UpdateOrganizationProfile,
// ) (*uoevents.Organization, error) {
// 	repo := wmrepo.Build(ctx)

// 	// Read the user from the database
// 	uoOrgRef, err := repo.ReadOrganizationByID(ctx, orgID)
// 	if err != nil {
// 		return nil, err
// 	} else if uoOrgRef == nil {
// 		herr := uoerror.NewNotFoundError(
// 			herror.Detail("Unable to find organization with ID %s", orgID),
// 		)
// 		return nil, herr
// 	}

// 	// Update the org's profile
// 	uoProfile := organizationProfileEventModel(profile)
// 	uoOrgRef.Profile = uoProfile

// 	return uoOrgRef, nil
// }
