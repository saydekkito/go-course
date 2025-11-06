package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/saydekkito/go-course/utils"
)

var jwtKey = []byte(utils.MustGetEnv("JWT_SECRET"))

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "Отсутствует заголовок Authorization", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Неверный формат токена", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Недействительный или просроченный токен", http.StatusUnauthorized)
			return
		}

		username, _ := claims["username"].(string)
		role, _ := claims["role"].(string)

		r.Header.Set("X-User", username)
		r.Header.Set("X-Role", role)

		next.ServeHTTP(w, r)
	}
}

func RequireRole(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Role") != role {
			http.Error(w, "Запрещено: недостаточные права", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}
