package engine

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtPayload struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
}

func HashJwtClaims(payload JwtPayload, secret string) (string, error) {
	claims := jwt.MapClaims{
		"userId": payload.UserId,
		"role":   payload.Role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyJwtClaims(token string, secret string) (*JwtPayload, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return &JwtPayload{
		UserId: claims["userId"].(string),
		Role:   claims["role"].(string),
	}, nil
}
