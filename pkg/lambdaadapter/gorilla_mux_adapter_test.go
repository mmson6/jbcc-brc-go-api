package lambdaadapter

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_GorillaMuxAdapter_SimplePingRequests_ProxiesEventCorrectly(t *testing.T) {
	//////////
	// SETUP

	// Context
	ctx := context.Background()

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("unfortunately-required-header", "")
		_, _ = fmt.Fprintf(w, "Home Page")
	})
	router.HandleFunc("/products", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("unfortunately-required-header", "")
		_, _ = fmt.Fprintf(w, "Products Page")
	})

	// Adapter
	adapter := NewGorillaMuxAdapter(router)

	//////////
	// ACTIONS

	// Need to test calling separate URLs to make sure the Lambda proxy is
	// properly routing requests instead of sending all requests to one
	// handler.

	// GET /
	homePageReq := events.APIGatewayProxyRequest{
		Path:       "/",
		HTTPMethod: "GET",
	}
	homePageResp, homePageReqErr := adapter.Proxy(ctx, homePageReq)

	// GET /products
	productsPageReq := events.APIGatewayProxyRequest{
		Path:       "/products",
		HTTPMethod: "GET",
	}
	productsPageResp, productsPageReqErr := adapter.Proxy(ctx, productsPageReq)

	//////////
	// ASSERTIONS

	require.NoError(t, homePageReqErr)
	assert.Equal(t, 200, homePageResp.StatusCode)
	assert.Equal(t, "Home Page", homePageResp.Body)

	require.NoError(t, productsPageReqErr)
	assert.Equal(t, 200, productsPageResp.StatusCode)
	assert.Equal(t, "Products Page", productsPageResp.Body)
}
