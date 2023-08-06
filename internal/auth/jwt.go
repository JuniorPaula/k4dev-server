package auth

import (
	"errors"
	"fmt"
	"knowledge-api/internal/config"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TokenGenerator(userID int64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 24).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

func TokenValidate(r *http.Request) error {
	strToken := getToken(r)
	token, err := jwt.Parse(strToken, getValidKeyAuthentication)
	if err != nil {
		return errors.New("unauthorized")
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("invalid token")
}

func CheckToken(token string) bool {
	jwtToken, err := jwt.Parse(token, getValidKeyAuthentication)
	if err != nil {
		return false
	}

	if _, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		return true
	}
	return false
}

func GetUserID(r *http.Request) (int64, error) {
	strToken := getToken(r)
	token, err := jwt.Parse(strToken, getValidKeyAuthentication)
	if err != nil {
		return 0, errors.New("unauthorized")
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseInt(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if err != nil {
			return 0, errors.New("unauthorized")
		}
		return userID, nil
	}

	return 0, errors.New("invalid token")
}

func getValidKeyAuthentication(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil,
			fmt.Errorf("invalid assignature method! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}
