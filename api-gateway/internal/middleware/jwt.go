// package middleware

// import (
// 	"context"
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// var jwtSecret = []byte("supersecretkey")

// type Claims struct {
// 	UserID string `json:"user_id"`
// 	Email  string `json:"email"`
// 	jwt.RegisteredClaims
// }

// func GenerateJWT(userID, email string) (string, error) {
// 	claims := &Claims{
// 		UserID: userID,
// 		Email:  email,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			Issuer:    "api-gateway",
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(jwtSecret)
// }

// func JWTAuth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authHeader := r.Header.Get("Authorization")
// 		if authHeader == "" {
// 			http.Error(w, "missing authorization header", http.StatusUnauthorized)
// 			return
// 		}

// 		parts := strings.Split(authHeader, " ")
// 		if len(parts) != 2 || parts[0] != "Bearer" {
// 			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
// 			return
// 		}

// 		tokenStr := parts[1]
// 		claims := &Claims{}

// 		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
// 			return jwtSecret, nil
// 		})

// 		if err != nil || !token.Valid {
// 			http.Error(w, "invalid or expired token", http.StatusUnauthorized)
// 			return
// 		}

// 		// Store user info in context
// 		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
// 		ctx = context.WithValue(ctx, "email", claims.Email)

// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
