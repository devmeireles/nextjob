package middlewares

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/devmeireles/nextjob/utils"
	"github.com/dgrijalva/jwt-go"
)

func AuthJwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		missingToken := errors.New("Missing authorization token")
		invalidToken := errors.New("Invalid authorization token")

		var header = r.Header.Get("Authorization")
		header = strings.TrimSpace(header)

		if header == "" {
			utils.ResErr(w, missingToken, http.StatusInternalServerError)
			return
		}

		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			utils.ResErr(w, invalidToken, http.StatusInternalServerError)
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)

		ctx := context.WithValue(r.Context(), "userID", claims["userID"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
