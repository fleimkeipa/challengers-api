package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/fleimkeipa/challengers-api/model"
	"github.com/fleimkeipa/challengers-api/uc"
	"github.com/fleimkeipa/challengers-api/util"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandlers struct {
	userUC *uc.UserUC
}

func NewUserHandlers(uc *uc.UserUC) *UserHandlers {
	return &UserHandlers{
		userUC: uc,
	}
}

// Register user
func (rc *UserHandlers) Register(c echo.Context) error {
	var input model.Register

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var user = model.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
		RoleID:   input.RoleID,
	}

	_, err := rc.userUC.Create(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"user": user.Username})
}

// User Login
func (rc *UserHandlers) Login(c echo.Context) error {
	var input model.Login

	if err := c.Bind(&input); err != nil {
		var errorMessage string
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			validationError := validationErrors[0]
			if validationError.Tag() == "required" {
				errorMessage = fmt.Sprintf("%s not provided", validationError.Field())
			}
		}

		return c.JSON(http.StatusBadRequest, echo.Map{"error": errorMessage})
	}

	user, err := rc.userUC.GetUserByUsername(c.Request().Context(), input.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	if err := model.ValidateUserPassword(user.Password, input.Password); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	jwt, err := util.GenerateJWT(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": jwt, "username": input.Username, "message": "Successfully logged in"})
}

func (rc *UserHandlers) Get(c echo.Context) error {
	var opts = getUserFindOpts(c)

	challenges, err := rc.userUC.Get(c.Request().Context(), opts)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"total": len(challenges),
		"pagination": echo.Map{
			"limit": opts.Limit,
			"skip":  opts.Skip,
		},
		"data": challenges,
	})
}

func getUserFindOpts(c echo.Context) model.UserFindOpts {
	return model.UserFindOpts{
		PaginationOpts: getPagination(c),
		RoleID:         getFilter(c, "role_id"),
		Username:       getFilter(c, "username"),
		Email:          getFilter(c, "email"),
	}
}
