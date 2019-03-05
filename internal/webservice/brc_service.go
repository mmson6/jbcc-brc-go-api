package webservice

import (
	"context"

	"github.com/gorilla/mux"

	"github.com/jbcc/brc-api/internal/readmodel/rmwebservice"
	"github.com/jbcc/brc-api/pkg/logger"
)

func BRC(ctx context.Context) *Service {
	svc := New()
	apiRoute := createAPIRoute(ctx, svc.Router)
	addReadModel(ctx, apiRoute)
	addWriteModel(ctx, apiRoute)
	return svc
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS

func createAPIRoute(ctx context.Context, router *mux.Router) *mux.Router {
	// Create the web service and prepare API endpoints to be attached under
	// /api/v1 (or whatever API_VERSION specifies).
	apiVer := "v1"
	apiPath := "/api/" + apiVer
	return router.PathPrefix(apiPath).Subrouter()
}

func addReadModel(ctx context.Context, router *mux.Router) {
	// Attach the read model endpoints
	log := logger.Current(ctx)
	log.Info("attaching read model endpoints")
	rmwebservice.AddRoutes(router)
}

func addWriteModel(ctx context.Context, router *mux.Router) {
	// Attach the read model endpoints
	log := logger.Current(ctx)
	log.Info("attaching write model endpoints")
	// wmwebservice.AddRoutes(router)
}
