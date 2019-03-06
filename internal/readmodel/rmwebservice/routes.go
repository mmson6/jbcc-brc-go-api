package rmwebservice

import (
	"github.com/gorilla/mux"

	"github.com/jbcc/brc-api/internal/webmiddleware"
)

func AddRoutes(r *mux.Router) {
	// GET requests to the read model
	rGets := r.Methods("GET").Subrouter()
	rGets.Use(webmiddleware.ResponseExpiresImmediately)

	// Health GET requests
	rGets.HandleFunc("/health", GetHealth)
	rGets.HandleFunc("/me/{user_id}", GetMeIdentity)
	rGets.HandleFunc("/records/{user_id}", GetUserRecords)
	// Authenticated GET requests
	// rGetsAuth := rGets.NewRoute().Subrouter()
	// rGetsAuth.Use(webmiddleware.AuthenticateRequestJWT)
	// rGetsAuth.HandleFunc("/me", GetMe)
	// rGetsAuth.HandleFunc("/me/organizations", GetMeOrganizations)
	// rGetsAuth.HandleFunc("/organizations", GetOrganizations)
	// rGetsAuth.HandleFunc("/organizations/{org_id}", GetOrganizationsID)
	// rGetsAuth.HandleFunc("/organizations/{org_id}/memberships", GetOrganizationsIDMemberships)
	// rGetsAuth.HandleFunc("/organizations/{org_id}/roles", GetOrganizationsIDRoles)
	// rGetsAuth.HandleFunc("/organizations/{org_id}/roles/{product}", GetOrganizationsIDRolesProduct)
	// rGetsAuth.HandleFunc("/users", GetUsers)
	// rGetsAuth.HandleFunc("/users/{user_id}", GetUsersID)

	// Read model event processor
	// rEvts := r.Path("/rm/events").Methods("POST").Subrouter()
	// rEvts.HandleFunc("", PostEventsNotification).Headers("x-amz-sns-message-type", "Notification")
	// rEvts.HandleFunc("", PostEventsSubscriptionConfirmation).Headers("x-amz-sns-message-type", "SubscriptionConfirmation")
	// rEvts.HandleFunc("", PostEventsUnsubscribeConfirmation).Headers("x-amz-sns-message-type", "UnsubscribeConfirmation")
}
