package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// check for valid admin token
func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := ValidateJWT(c); err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Authentication required"})
		}

		if err := ValidateAdminRoleJWT(c); err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Only Administrator is allowed to perform this action"})
		}

		return next(c)
	}
}

// check for valid challanger token
func JWTAuthChallenger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := ValidateJWT(c); err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Authentication required"})
		}

		if err := ValidateChallengerRoleJWT(c); err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Only registered Challengers are allowed to perform this action"})
		}

		return next(c)
	}
}
