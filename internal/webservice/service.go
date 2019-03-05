package webservice

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gorilla/mux"

	"github.com/jbcc/brc-api/pkg/lambdaadapter"
	"github.com/jbcc/brc-api/pkg/logger"
)

type Service struct {
	Router *mux.Router

	gorillaAdapter *lambdaadapter.GorillaMuxAdapter
	server         *http.Server
}

func New() *Service {
	router := mux.NewRouter()

	return &Service{
		Router: router,
	}
}

func (s *Service) Start() {
	lambda.Start(s.lambdaHandler)
}

////////////////////////////////////////////////////////////////////////////////
// PRIVATE FUNCTIONS

func (s *Service) lambdaHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log := logger.Current(ctx)

	// It is safe to initialize the gorillaAdapter by checking nil each time
	// lambdaHandler() is invoked because Lambda guarantees that a Lambda
	// instance will only handle one request at a time; i.e., there are no
	// threading concerns. Checking for nil is much faster than using
	// sync.Once{}.Do().
	if s.gorillaAdapter == nil {
		log.Info("lambda cold start")
		s.gorillaAdapter = lambdaadapter.NewGorillaMuxAdapter(s.Router)
	}

	return s.gorillaAdapter.Proxy(ctx, req)
}
