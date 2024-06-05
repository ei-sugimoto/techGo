package pkg

import (
	"log/slog"

	"github.com/dgrijalva/jwt-go"
)

func EncodeJwt(userID string) string {
	claims := jwt.MapClaims{
		"user_id": userID,
	}
	logger := NewLogger().With(slog.String("path", "pkg/"))

	tokenMap := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenMap.SignedString([]byte("secret"))
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}

	return token
}

func DecodeJwt(token string, needColumn string) (string, error) {
	logger := NewLogger().With(slog.String("path", "pkg/"))

	tokenMap, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	claims, ok := tokenMap.Claims.(jwt.MapClaims)
	if !ok {
		logger.Error("failed to get claims")
		return "", err
	}

	decodeString := claims[needColumn].(string)
	return decodeString, nil
}
