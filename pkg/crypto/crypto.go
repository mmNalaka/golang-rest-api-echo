package crypto

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"golang-rest-api-echo/pkg/config"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, 12)

	if err != nil {
		return "", errors.Wrap(err, "Error creating password.")
	}

	return string(hashedPassword), nil
}

func ComparePasswordWithHash(hash string, password string) error {
	passwordBytes := []byte(password)
	hashBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)
	return errors.Wrap(err, "Error comparing password and hash")
}

func CreateJwtToken(userId string, cfg config.AppConfig) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	expiresIn, err := strconv.ParseInt(cfg.JWtExpiresIn, 10, 64)

	if err != nil {
		return "", errors.Wrap(err, "Error parsing int.")
	}

	expiration := time.Duration(int64(time.Minute) * expiresIn)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userId
	claims["expires_in"] = time.Now().Add(expiration).Unix()

	t, err := token.SignedString([]byte(cfg.JWtSecret))
	if err != nil {
		return "", errors.Wrap(err, "Error signing JWT")
	}

	return t, nil
}
