package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

const (
	SUPER_SECRET_KEY    = "ANDRUHA CHMO EBANOE"
	TokenDuration       = 15 * time.Minute
	ErrMsgNotAuthorized = "Missing login data"
	ContextKeyUserID    = "user_id"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId int
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if parts[0] != "Bearer" && len(parts) != 2 {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]
		userID, err := getUserLogin(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyUserID, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
func GenerateJWT(id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenDuration)),
		},
		UserId: id,
	})

	tokenString, err := token.SignedString([]byte(SUPER_SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func getUserLogin(jwtString string) (int, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(jwtString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(SUPER_SECRET_KEY), nil
	})
	if err != nil {
		return -1, err
	}
	if !token.Valid {
		return -1, err
	}
	return claims.UserId, nil
}
