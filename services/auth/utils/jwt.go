package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var as = []byte("access-secret")
var rs = []byte("refresh-secret")

func GenerateTokens(em string) (string, string, error) {
	a := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": em,
		"exp":   time.Now().Add(15 * time.Minute).Unix(),
	})
	r := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": em,
		"exp":   time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	at, err := a.SignedString(as)
	if err != nil {
		return "", "", err
	}

	rt, err := r.SignedString(rs)
	if err != nil {
		return "", "", err
	}

	return at, rt, nil
}

func VerifyRefreshToken(t string) (string, error) {
	tok, err := jwt.Parse(t, func(jt *jwt.Token) (interface{}, error) {
		return rs, nil
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
