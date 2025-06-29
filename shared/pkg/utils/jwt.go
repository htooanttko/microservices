package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var as = []byte("access-secret")

func VerifyAccessToken(t string) (string, error) {
	tok, err := jwt.Parse(t, func(jt *jwt.Token) (interface{}, error) {
		return as, nil
	})
	if err != nil || !tok.Valid {
		return "", nil
	}

	c, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	em := c["email"].(string)
	return em, nil
}
