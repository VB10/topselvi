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

// VerifyUserToken method control userToken in the firebase.
func VerifyUserToken(userToken string) error {
	app := FBInstance()
	client, err := app.Auth(context.Background())
	if err != nil {
		return err
	}
	verifyToken, err := client.VerifyIDToken(context.Background(), userToken)
	if err != nil {
		log.Print(verifyToken)
		return err
	}
	return nil
}

//GetUserData function take UserInfo in the firebase db.
func GetUserData(userID string) (*Users, error) {
	if len(userID) == 0 {
		var err = errors.New("User token must be required");
		return nil, err
	}

	var ctx = context.Background()
	app := FBInstance()

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	verifyToken, err := client.VerifyIDToken(ctx, userID)
	if err != nil {
		return nil, err
	}

	database, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	document, err := database.Collection(FirestoreUsers).Doc(verifyToken.UID).Get(ctx)
	if err != nil {
		return nil, err
	}

	var user Users
	if err := document.DataTo(&user); err != nil {
		return nil, err
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
	userID := fmt.Sprintf("%v", jwtParsed[FbUid])

	 customToken, err := client.CustomToken(ctx, userID);
	 if err != nil {
		return err
	}

	print(customToken)


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
