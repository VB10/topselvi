package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
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
func RefreshUserToken(token string) error {
	var ctx = context.Background()
	app := FBInstance()

	jwtParsed, error := JWTParser(token)

	if len(jwtParsed) == 0 {
		return errors.New("error")
	}

	client, error := app.Auth(ctx)
	if error != nil {
		return error
	}
	userID := fmt.Sprintf("%v", jwtParsed[FB_UID])

	error = client.RevokeRefreshTokens(ctx, userID)
	if error != nil {
		return error
	}
	return nil
}

func JWTParser(key string) (jwt.MapClaims, error) {

	token, error := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	});
	claims := token.Claims.(jwt.MapClaims)
	return claims, error
}
