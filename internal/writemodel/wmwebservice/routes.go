package wmwebservice

import (
	"github.com/gorilla/mux"
)

func AddRoutes(r *mux.Router) {
	// Write model commands
	rCmds := r.NewRoute().Subrouter()
	// rCmds.Use(webmiddleware.AuthenticateRequestJWT)
	rCmds.HandleFunc("/user", PutUserProfile)
	rCmds.HandleFunc("/records", PutUserRecords)
}
