package cmd

import "github.com/dgrijalva/jwt-go"

func JWTParser(key string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(key, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	if token == nil {
		return nil, err
	}

	if err != nil && err.Error() != JWTFirebaseKeyError {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}