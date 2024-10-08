package util

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fleimkeipa/challengers-api/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// retrieve JWT key from .env file
var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

// generate JWT token
func GenerateJWT(user model.User) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID,
		"role": user.RoleID,
		"iat":  time.Now().Unix(),
		"eat":  time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString(privateKey)
}

// validate JWT token
func ValidateJWT(c echo.Context) error {
	token, err := getToken(c)
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if ok {
		return nil
	}

	return errors.New("invalid token provided")
}

// validate Admin role
func ValidateAdminRoleJWT(context echo.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token claims")
	}

	userRole := uint(claims["role"].(float64))
	if userRole == model.AdminRole {
		return nil
	}

	return errors.New("invalid admin token provided")
}

// validate Challenger role
func ValidateChallengerRoleJWT(context echo.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token claims")
	}

	var userRole = uint(claims["role"].(float64))
	if userRole == model.ChallengerRole || userRole == model.AdminRole {
		return nil
	}

	return errors.New("invalid challenger or admin token provided")
}

// check token validity
func getToken(context echo.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})

	return token, err
}

// extract token from request Authorization header
func getTokenFromRequest(c echo.Context) string {
	bearerToken := c.Request().Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
