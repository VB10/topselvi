package cmd

import (
	"github.com/VB10/topselvi/utility"
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
		apiKey := r.Header.Get(QueryApiKey)
		userToken := r.Header.Get(QueryUserToken)

		//TODO: FIX API KEY
		if len(apiKey) == 0 {
			// Report Unauthorized
			utility.GenerateError(w, nil, http.StatusNotFound, ApiKeyEmptyError)
			return
		}

		if err := VerifyUserToken(userToken)
			err != nil {
			utility.GenerateError(w, err, http.StatusUnauthorized, "token error")
			return
		}

		next.ServeHTTP(w, r)
	})
}
