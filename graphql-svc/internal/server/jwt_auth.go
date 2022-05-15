package server

import (
	"context"
	"net/http"
)

func JwtAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !validateJwt(r.Header.Get("Authorization")) {
			w.WriteHeader(http.StatusForbidden)
		} else {
			r = r.WithContext(context.WithValue(r.Context(), "user_id", "1"))
			next.ServeHTTP(w, r)
		}
	}

	return http.HandlerFunc(fn)
}

func validateJwt(auth string) bool {
	return true
}
