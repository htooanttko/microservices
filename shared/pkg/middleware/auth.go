package middleware

import (
	"net/http"
	"strings"

	"github.com/htooanttko/microservices/shared/pkg/responses"
	"github.com/htooanttko/microservices/shared/pkg/utils"
)

func AuthMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			responses.WithError(rw, http.StatusUnauthorized, "missing token")
			return
		}

		t := strings.TrimPrefix(auth, "Bearer ")
		em, err := utils.VerifyAccessToken(t)
		if err != nil {
			responses.WithError(rw, http.StatusUnauthorized, "invalid token")
			return
		}

		ctx := utils.WithUserEmail(r.Context(), em)
		n.ServeHTTP(rw, r.WithContext(ctx))
	})
}
