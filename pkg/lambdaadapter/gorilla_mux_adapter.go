// Copyright 2018 Amazon.com, Inc. or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// See:
// - https://github.com/awslabs/aws-lambda-go-api-proxy/blob/5e023998abc6065e7a1ba4efb4c66fbd5eb7163a/NOTICE
// - https://github.com/awslabs/aws-lambda-go-api-proxy/blob/5e023998abc6065e7a1ba4efb4c66fbd5eb7163a/LICENSE

package lambdaadapter

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/gorilla/mux"
)

// An adapter from a Lambda API Gateway proxy request to a gorilla/mux HTTP
// router. This implementation embeds the context provided with the Lambda
// request inside the http.Request for use by the request handler. This enables
// use of libraries that rely on information in the Lambda context such as
// X-Ray.
//
// The github.com/awslabs/aws-lambda-go-api-proxy/gorillamux adapter from AWS
// Labs was written to support a version of go before the context package was
// introduced. For X-Ray to work within Lambda, the context provided in the
// call to the Lambda handler must be used throughout the request.
type GorillaMuxAdapter struct {
	core.RequestAccessor
	router *mux.Router
}

// Create a new adapter from a Lambda API Gateway proxy request to a
// gorilla/mux router.
func NewGorillaMuxAdapter(router *mux.Router) *GorillaMuxAdapter {
	return &GorillaMuxAdapter{
		router: router,
	}
}

// Handle a Lambda API Gateway proxy request. In addition to the standard
// aws-lambda-go-api-proxy core handling, this method embeds the Lambda's
// context in the http.Request handed to the gorilla/mux router.
func (h *GorillaMuxAdapter) Proxy(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req, err := h.ProxyEventToHTTPRequest(event)
	if err != nil {
		return core.GatewayTimeout(), core.NewLoggedError("Could not convert proxy event to request: %v", err)
	}
	// Set provided Host
	if host, ok := event.Headers["Host"]; ok {
		req.Host = host
		if req.URL != nil {
			req.URL.Host = host
		}
	}
	// Set provided context
	contextualReq := req.WithContext(ctx)

	w := core.NewProxyResponseWriter()
	h.router.ServeHTTP(http.ResponseWriter(w), contextualReq)

	resp, err := w.GetProxyResponse()
	if err != nil {
		return core.GatewayTimeout(), core.NewLoggedError("Error while generating proxy response: %v", err)
	}

	return resp, nil
}
