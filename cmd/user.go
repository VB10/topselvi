package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Users struct {
	IsFirstLogin bool   `json:"isFirstLogin"`
	Mail         string `json:"mail"`
	Username     string `json:"username"`
	Wallet       int    `json:"wallet"`
	UserID       string `json:"userID"`
	Picture       string `json:"picture"`
}

type CustomToken struct {
	Token             string `json:"token"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
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
func GetUserData(userToken string) (*Users, error) {
	var ctx = context.Background()
	app := FBInstance()

	database, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	claims, err := JWTParser(userToken)
	if err != nil {
		return nil, err
	}
	userID := fmt.Sprintf("%v", claims[FbUid])
	document, err := database.Collection(FirestoreUsers).Doc(userID).Get(ctx)
	if err != nil {
		return nil, err
	}
	var user Users
	if err := document.DataTo(&user); err != nil {
		return nil, err
	}

	user.UserID = userID
	return &user, nil
}

//GetUserData function take UserInfo in the firebase db.
func UpdateUserData(user Users) error {

	var ctx = context.Background()
	app := FBInstance()

	database, err := app.Firestore(ctx)
	if err != nil {
		return err
	}

	//document, err := database.Collection(FirestoreUsers) .Update(ctx, []firestore.Update{{Path: "wallet", Value: user.Wallet}} )
	document, err := database.Collection(FirestoreUsers).Doc("u2wevTndeKRHGai0b9KeHXKDXU32").Set(ctx, user)
	if err != nil {
		return err
	}
	print(document)
	return nil
}

//RefreshUserToken function take UserInfo in the firebase db.
func RefreshUserToken(token string) (string, error) {
	var ctx = context.Background()
	app := FBInstance()

	jwtParsed, err := JWTParser(token)
	if len(jwtParsed) == 0 {
		return "", errors.New("error")
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return "", err
	}
	userID := fmt.Sprintf("%v", jwtParsed[FbUid])

	customToken, err := client.CustomToken(ctx, userID)
	if err != nil {
		return "", err
	}

	var customTokenModel CustomToken
	customTokenModel.Token = customToken
	customTokenModel.ReturnSecureToken = true

	var x, _ = json.Marshal(customTokenModel)

	resp, err := http.Post(FirebaseAuthSigninCustomToken, "application/json", bytes.NewBuffer(x))
	if err != nil {
		return "", nil
	}
	var result map[string]interface{}

	_ = json.NewDecoder(resp.Body).Decode(&result)
	idToken := result[QueryIDToken].(string)
	return idToken, nil
}

// CreateDefaultUser create new user data
func CreateDefaultUser(userToken string)  (*Users, error) {
	var ctx = context.Background()
	app := FBInstance()

	database, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	claims, err := JWTParser(userToken)
	if err != nil {
		return nil, err
	}
	userID := fmt.Sprintf("%v", claims[FbUid])
	email := fmt.Sprintf("%v", claims[FBEmail])
	name := fmt.Sprintf("%v", claims[FBName])
	picture := fmt.Sprintf("%v", claims[FBPicture])


	var newUser = Users{
		IsFirstLogin: false,
		Mail:         email,
		Username:     name,
		Wallet:       NewUserWallet,
		UserID:       userID,
		Picture:      picture,
	}

	_ , err = database.Collection(FirestoreUsers).Doc(userID).Set(ctx, newUser)
	if err != nil {
		return nil,err
	}

	return &newUser, nil
}
