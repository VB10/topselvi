package cmd

import (
	"context"
	"errors"
	"log"
)

type Users struct {
	IsFirstLogin bool   `json:"isFirstLogin"`
	Mail         string `json:"mail"`
	Username     string `json:"username"`
	Wallet       int    `json:"wallet"`
	UserID       string `json:"userID"`
}

// VerifyUserToken method control userid in the firebase.
func VerifyUserToken(userID string) error {

	ctx := context.Background()
	app := FBInstance()
	client, error := app.Auth(ctx)
	if error != nil {
		return error
	}

	verifyToken, error := client.VerifyIDToken(ctx, userID)
	if error != nil {
		log.Print(verifyToken)
		return error
	}
	return nil
}

//GetUserData function take UserInfo in the firebase db.
func GetUserData(userID string) (*Users, error) {
	if len(userID) == 0 {
		var error = errors.New("User token must be required");
		return nil, error;
	}

	var ctx = context.Background()
	app := FBInstance()

	client, error := app.Auth(ctx)
	if error != nil {
		return nil, error
	}

	verifyToken, error := client.VerifyIDToken(ctx, userID)
	if error != nil {
		return nil, error
	}

	database, error := app.Firestore(ctx)
	if error != nil {
		return nil, error
	}

	document, error := database.Collection(FIRESTORE_USERS).Doc(verifyToken.UID).Get(ctx)
	if error != nil {
		return nil, error
	}

	var user Users

	if err := document.DataTo(&user); err != nil {
		return nil, error
	}

	user.UserID = verifyToken.UID
	return &user, nil
}

//RefreshUserToken function take UserInfo in the firebase db.
func RefreshUserToken(userID string) error {
	// TODO : JWT PARSER ADDED
	//eyJhbGciOiJSUzI1NiIsImtpZCI6ImVlMjc0MWQ0MWY5ZDQzZmFiMWU2MjhhODVlZmI0MmE4OGVmMzIyOWYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20veW91Mndpbi0zYjlkOSIsImF1ZCI6InlvdTJ3aW4tM2I5ZDkiLCJhdXRoX3RpbWUiOjE1Njg4NDgzMjUsInVzZXJfaWQiOiJ1MndldlRuZGVLUkhHYWkwYjlLZUhYS0RYVTMyIiwic3ViIjoidTJ3ZXZUbmRlS1JIR2FpMGI5S2VIWEtEWFUzMiIsImlhdCI6MTU2ODg0ODMyNSwiZXhwIjoxNTY4ODUxOTI1LCJlbWFpbCI6InZlbGlAdGVzdC5jb20iLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZW1haWwiOlsidmVsaUB0ZXN0LmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6InBhc3N3b3JkIn19.yXpRYd-e4YLqy-Zlk6jPpNklX0Ml-BEF1Q0MgFX7DQUPCU78oa-DJunVbBeTqBAo0iJsqnkaV3LJE5GtYyJ2kA-94cQPE8xuh0BlL9MeHQQpKdg3_uhGI6RM9PSWc51_RnP7RV-V4220IDkTyJpN3MjDH9Aqx35Tn3l9VqGFcV1KpFFkHghClGRGCGWVcJCvShMeAgLWyzp-EB622MQczlpFgFaBP1dAD3glWswPIHwErVVhqPOGesBuxBKix3khe6RVxJTuVf5IPaNLK0k84arE1A6u_P_ImVayW9BNo4hYmQ9Dcw_YmY-oCj1ylwgZHC0oMmLXFDOG_zLdiAwoqw
	var ctx = context.Background()
	app := FBInstance()

	client, error := app.Auth(ctx)
	if error != nil {
		return error
	}

	error = client.RevokeRefreshTokens(ctx, userID)
	if error != nil {
		return error
	}

	return nil
}
