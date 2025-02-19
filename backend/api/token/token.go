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

var ATMaxAgeStd = int(15 * time.Minute / time.Second)
var RTMaxAgeStd = int(30 * time.Minute / time.Second)

func GenerateAccessToken(userID string) (accessToken string, err error) {
	accessClaims := map[string]interface{}{
		"user_id": userID,
		"exp":     time.Now().Add(time.Duration(ATMaxAgeStd) * time.Second),
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
		"exp":     time.Now().Add(time.Duration(RTMaxAgeStd) * time.Second),
	}
	_, refreshToken, err = TokenAuth.Encode(refreshClaims)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}
