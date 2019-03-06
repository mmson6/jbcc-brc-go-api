package wmwebservice

import (
	"github.com/gorilla/mux"
	// "code.siemens.com/horizon/platform-verticals/user/uo-api-go/internal/webmiddleware"
)

func AddRoutes(r *mux.Router) {
	// Write model commands
	rCmds := r.NewRoute().Subrouter()
	// rCmds.Use(webmiddleware.AuthenticateRequestJWT)
	rCmds.HandleFunc("/me/identity", PutMeIdentity)
	rCmds.HandleFunc("/records/{user_id}", PutUserRecords)
	// rCmds.HandleFunc("/me/invitations/{inv_key}/state", PutMeInvitationsState)
	// rCmds.HandleFunc("/organizations/{org_id}/invitations", PostOrganizationsInvitations)
	// rCmds.HandleFunc("/organizations/{org_id}/memberships/{ms_id}/roles", PutOrganizationMembershipRoles)
	// rCmds.HandleFunc("/organizations/{org_id}/profile", PutOrganizationsProfile)
	// rCmds.HandleFunc("/roles/{role_technical_name}", PutRoles)

	// // Write model event processor
	// rEvts := r.Path("/wm/events").Methods("POST").Subrouter()
	// rEvts.HandleFunc("", PostEventsNotification).Headers("x-amz-sns-message-type", "Notification")
	// rEvts.HandleFunc("", PostEventsSubscriptionConfirmation).Headers("x-amz-sns-message-type", "SubscriptionConfirmation")
	// rEvts.HandleFunc("", PostEventsUnsubscribeConfirmation).Headers("x-amz-sns-message-type", "UnsubscribeConfirmation")
}
