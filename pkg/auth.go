package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JwtPayload struct {
	Sub string `json:"sub"`
}

func ComparePassword(currentPass, truePass string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(truePass), []byte(currentPass)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, fmt.Errorf("wrong password")
		}
		return false, err
	}
	return true, nil
}

func GenerateToken(secret string, expirationInHour int, jwtPayload JwtPayload) (string, map[string]any, error) {
	claims := jwt.MapClaims{
		"iss": "dealls",
		"sub": jwtPayload.Sub,
		"exp": time.Now().Add(time.Hour * time.Duration(expirationInHour)).Unix(),
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, claims, err
}

func GetJwtPayload(ctx context.Context) (JwtPayload, error) {
	_, claims, _ := jwtauth.FromContext(ctx)
	sub, ok := claims["sub"]
	if !ok {
		return JwtPayload{}, fmt.Errorf("sub not found")
	}

	subString, ok := sub.(string)
	if !ok {
		return JwtPayload{}, fmt.Errorf("failed to cast sub to string")
	}

	return JwtPayload{
		Sub: subString,
	}, nil
}
