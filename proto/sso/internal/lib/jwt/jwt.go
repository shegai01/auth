package jwt

import (
	"auth/internal/domain/models"
	"auth/internal/lib/jwt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func NewToken(usr models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = usr.ID
	claims["email"] = usr.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["app_id"] = app.ID

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
