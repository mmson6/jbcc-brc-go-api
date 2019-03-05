package webresponse

import (
	"context"
	"net/http"
)

func WriteEmpty(ctx context.Context, w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}
