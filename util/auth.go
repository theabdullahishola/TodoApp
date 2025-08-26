package util

import (
	"errors"

	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateRefreshToken(id, email string) (string, error) {

	var refresh_secret = os.Getenv("Refresh_secret")
	refreshClaims := jwt.MapClaims{
		"userID": id,
		"email":  email,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256,
		refreshClaims).SignedString([]byte(refresh_secret))
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}
func GenerateAccessToken(id, email string) (string, error) {
	var access_secret = os.Getenv("SECRET_KEY")

	accessClaims := jwt.MapClaims{
		"userID": id,
		"email":  email,
		"exp":    time.Now().Add(15 * time.Minute).Unix(),
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(access_secret))

	if err != nil {
		return "", err
	}

	return accessToken, nil
}
func VerifyAccessToken(token string) (string, error) {
	var access_secret = os.Getenv("SECRET_KEY")
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invaild signing methos")
		}
		return []byte(access_secret), nil
	})

	if err != nil {
		return "", errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return "", errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("cant get claim")
	}
	userID, ok := claims["userID"].(string)
	if !ok {
		return "", errors.New("userID not found in token")
	}
	return userID, nil
}

func VerifyRefreshToken(token string) (string, error) {
	var refresh_secret = os.Getenv("Refresh_secret")
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invaild signing methos")
		}
		return []byte(refresh_secret), nil
	})

	if err != nil {
		return "", errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return "", errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("cant get claim")
	}
	userID, ok := claims["userID"].(string)
	if !ok {
		return "", errors.New("userID not found in token")
	}
	return userID, nil
}
