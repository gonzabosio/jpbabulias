package token

import (
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth

func NewTokenAuth() {
	TokenAuth = jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
}

var ATMaxAgeStd = 1 * 60
var RTMaxAgeStd = 2 * 60

func GenerateAccessToken(userID string) (accessToken string, err error) {
	accessClaims := map[string]interface{}{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(ATMaxAgeStd)).Unix(),
	}
	_, accessToken, err = TokenAuth.Encode(accessClaims)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func GenerateRefreshToken(userID string) (refreshToken string, err error) {
	refreshClaims := map[string]interface{}{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(RTMaxAgeStd)).Unix(),
	}
	_, refreshToken, err = TokenAuth.Encode(refreshClaims)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}
