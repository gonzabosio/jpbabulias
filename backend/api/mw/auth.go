package mw

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/gonzabosio/jpbabulias/api/handlers"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			handlers.WriteJSON(w, map[string]string{
				"message":       "Invalid authentication",
				"error_details": err.Error(),
			}, http.StatusUnauthorized)
			return
		}
		if token == nil {
			handlers.WriteJSON(w, map[string]string{
				"message": "Access token is missing",
			}, http.StatusUnauthorized)
			return
		}

		rtCk, err := r.Cookie("refresh_token")
		if err != nil {
			if http.ErrNoCookie == err {
				handlers.WriteJSON(w, map[string]string{
					"message":       "Refresh token is missing",
					"error_details": err.Error(),
				}, http.StatusBadRequest)
				return
			} else {
				handlers.WriteJSON(w, map[string]string{
					"message":       "Refresh token expired",
					"error_details": err.Error(),
				}, http.StatusUnauthorized)
				return
			}
		}
		fmt.Println("Refresh Token:", rtCk.Value)

		// accTknCk, err := r.Cookie("access_token")
		// if err != nil {
		// 	if http.ErrNoCookie == err {
		// 		handlers.WriteJSON(w, map[string]string{
		// 			"message":       "Access token is missing",
		// 			"error_details": err.Error(),
		// 		}, http.StatusBadRequest)
		// 		return
		// 	} else {
		// 		handlers.WriteJSON(w, map[string]string{
		// 			"message":       "Access token expired",
		// 			"error_details": err.Error(),
		// 		}, http.StatusUnauthorized)
		// 		return
		// 	}
		// }
		// fmt.Println("Access Token:", accTknCk.Value)
		ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
