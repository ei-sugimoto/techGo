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
