package middleware

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(tokenStr string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	klem := token.Claims.(*JwtClaims)
	fmt.Println(klem.Username)

	return klem, nil
}
