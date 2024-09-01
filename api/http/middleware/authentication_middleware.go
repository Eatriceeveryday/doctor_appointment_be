package middleware

import (
	"BackendTugasAkhir/api/http/utils"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"strings"
	"time"
)

func AuthenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")

		if reqToken == "" {
			utils.JSONResponse(w, utils.Response{Msg: "Invalid Token"}, http.StatusBadRequest)
			return
		}

		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = strings.TrimSpace(splitToken[1])

		id, err := validateToken(reqToken)
		if err != nil {
			utils.JSONResponse(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateToken(token string) (string, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})

	if err != nil {
		return "", err
	}

	if !tkn.Valid {
		return "", errors.New("Invalid Token")
	}

	claims := tkn.Claims.(jwt.MapClaims)
	exp := claims["exp"].(float64)
	if exp <= float64(time.Now().Unix()) {
		return "", errors.New("Expired Token")
	}

	return claims["id"].(string), nil
}
