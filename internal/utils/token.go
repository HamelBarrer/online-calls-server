package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

func CreationToken(userId int) (string, error) {
	claims := &JwtCustomClaims{
		strconv.Itoa(userId),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidateToken(t string) (*JwtCustomClaims, bool, error) {
	token, err := jwt.ParseWithClaims(t, &JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	}, jwt.WithLeeway(5*time.Second))
	if err != nil {
		return nil, false, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if ok && token.Valid {
		return claims, true, nil
	}

	return nil, false, nil
}
