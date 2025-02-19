package mw

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gonzabosio/jpbabulias/api/handlers"
	"github.com/gonzabosio/jpbabulias/api/token"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// REFRESH TOKEN
		rtCk, err := r.Cookie("refresh_token")
		if err != nil {
			if http.ErrNoCookie == err {
				handlers.WriteJSON(w, map[string]string{
					"message":    "Refresh token is missing",
					"error_dets": err.Error(),
				}, http.StatusUnauthorized)
				return
			} else {
				handlers.WriteJSON(w, map[string]string{
					"message":    "Refresh token expired",
					"error_dets": err.Error(),
				}, http.StatusUnauthorized)
				return
			}
		}
		rToken, err := token.TokenAuth.Decode(rtCk.Value)
		if err != nil {
			handlers.WriteJSON(w, map[string]string{
				"message":    "Invalid authentication",
				"error_dets": fmt.Sprintf("error in refresh token validation instance: %v", err.Error()),
			}, http.StatusUnauthorized)
			return
		}
		if rToken == nil {
			handlers.WriteJSON(w, map[string]string{
				"message":    "Invalid token",
				"error_dets": "refresh token decodification: no data from token",
			}, http.StatusUnauthorized)
			return
		}
		userId := rToken.PrivateClaims()["user_id"].(string)

		// ACCESS TOKEN
		accTknCk, err := r.Cookie("access_token")
		if err != nil {
			aToken, err := token.GenerateAccessToken(userId)
			if err != nil {
				handlers.WriteJSON(w, map[string]string{
					"message":    "Failed to generate new access token",
					"error_dets": fmt.Sprintf("in middleware: %v", err.Error()),
				}, http.StatusInternalServerError)
				return
			}
			handlers.SetCookie(w, "access_token", aToken, token.ATMaxAgeStd)
			return
		}

		aToken, err := token.TokenAuth.Decode(accTknCk.Value)
		if err != nil {
			handlers.WriteJSON(w, map[string]string{
				"message":    "Invalid authentication",
				"error_dets": fmt.Sprintf("error in access token validation instance: %v", err.Error()),
			}, http.StatusUnauthorized)
			return
		}
		if aToken == nil {
			handlers.WriteJSON(w, map[string]string{
				"message":    "Invalid token",
				"error_dets": "access token decodification: no data from token",
			}, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
