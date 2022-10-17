package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	Secret []byte
}

func (j *JWT) Sign(exp time.Time, data any) (string, error) {
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": data,
		"exp":  exp,
	}).SignedString(j.Secret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
