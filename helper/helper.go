package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
	// "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"finalProject/model"
)

func GenerateID() string {
	return uuid.New().String()
}

func Hash(plain string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plain), 8)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func GenerateTime() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

func IsHashValid(hash string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))

	return err == nil
}

func GenerateAccessToken(userID string, email string) (string, error) {
	claims := jwt.MapClaims{
		"email":   email,
		"user_id": userID,
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
	}

	return generateToken(claims, "4mAn___80SS")
}

func generateToken(claims jwt.MapClaims, secret string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := jwtToken.SignedString([]byte(secret))
	return tokenString, err
}

func VerifyAccessToken(token string) (*jwt.Token, error) {

	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, model.ErrorInvalidToken
		}

		return []byte("4mAn___80SS"), nil
	})

	return jwtToken, err
}
