package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/yudistirarivaldi/technical-test-deeptech/internal/utils"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func JWTMiddleware(secret string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{
				"responseCode": "01",
				"message":      "Missing or invalid Authorization header",
			})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := utils.ParseJWT(tokenStr, secret)
		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{
				"responseCode": "01",
				"message":      "Invalid or expired token",
			})
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)

		next(w, r.WithContext(ctx))
	}
}

func GetUserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(UserIDKey).(int64)
	return userID, ok
}

func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value("role").(string)
		if !ok || role != "admin" {
			utils.WriteJSON(w, http.StatusForbidden, map[string]string{
				"responseCode": "03",
				"message":      "Forbidden: Admins only",
			})
			return
		}

		next(w, r)
	}
}
