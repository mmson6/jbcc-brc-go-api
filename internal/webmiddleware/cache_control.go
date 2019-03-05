package webmiddleware

import "net/http"

func ResponseExpiresImmediately(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Set("Cache-Control", "max-age=0, must-revalidate")
		next.ServeHTTP(w, r)
	})
}
