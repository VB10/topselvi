package cmd

import (
	"context"
	"net/http"

	firebase "firebase.google.com/go"
)

func init() {}

//BaseError return you4win base error model.
type BaseError struct {
	User    string `json:"user,omitempty"`
	Message string `json:"message,omitempty"`
}

func verifyUser(_app *firebase.App, userID string, w http.ResponseWriter) (string, BaseError) {
	ctx := context.Background()
	client, err := _app.Auth(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", BaseError{userID, err.Error()}
	}

	token, err := client.GetUser(ctx, userID)
	if err != nil {

		http.Error(w, err.Error(), http.StatusNotFound)
		return "", BaseError{userID, err.Error()}
	}

	return token.UID, BaseError{}
}
