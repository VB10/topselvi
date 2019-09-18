package cmd

import (
	"../utility"
	"io"
	"net/http"
)

// Middleware (this function) makes adding more than one layer of middleware easy
// by specifying them as a list. It will run the last specified handler first.
func Middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}

// AuthMiddleware is an example of a middleware layer. It handles the request authorization
// by checking for a key in the url.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestKey := r.Header.Get("token")
		userID := r.Header.Get("userID")

		if error := VerifyUserToken(userID); error != nil {
			utility.GenerateError(w, error, http.StatusUnauthorized, "token error")
			return
		}
		if len(requestKey) == 0 {
			// Report Unauthorized
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"error":"invalid_key"}`)
			return
		}
		next.ServeHTTP(w, r)
	})
}
