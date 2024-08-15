package middleware

import (
	"jwt-try/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte("ayam")

type Claimsasdasda struct {
	UserID   int    `json:"userID"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJwt(user domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, Claimsasdasda{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "DIDI KEREN!",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Hour)),
		},
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
