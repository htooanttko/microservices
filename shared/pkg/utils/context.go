package utils

import (
	"context"
	"net/http"
)

type contextkey string

const userEmailKey = contextkey("user_email")

func WithUserEmail(ctx context.Context, em string) context.Context {
	return context.WithValue(ctx, userEmailKey, em)
}

func GetEmailFromToken(r *http.Request) (string, bool) {
	em, ok := r.Context().Value(userEmailKey).(string)
	return em, ok
}
