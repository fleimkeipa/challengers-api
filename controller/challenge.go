package controller

import (
	"net/http"

	"github.com/fleimkeipa/challengers-api/model"
	"github.com/fleimkeipa/challengers-api/uc"

	"github.com/labstack/echo/v4"
)

type ChallengeHandlers struct {
	chUC *uc.ChallengeUC
}

func NewChallengeHandlers(uc *uc.ChallengeUC) *ChallengeHandlers {
	return &ChallengeHandlers{
		chUC: uc,
	}
}

func (rc *ChallengeHandlers) Create(c echo.Context) error {
	var input model.ChallengeRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var challenge = model.Challenge{
		ID:   input.ID,
		Name: input.Name,
	}

	_, err := rc.chUC.Create(c.Request().Context(), challenge)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"challenge": challenge.Name})
}

func (rc *ChallengeHandlers) Update(c echo.Context) error {
	var input model.ChallengeRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	var challenge = model.Challenge{
		ID:   input.ID,
		Name: input.Name,
	}

	_, err := rc.chUC.Update(c.Request().Context(), challenge)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"challenge": challenge.Name})
}

func (rc *ChallengeHandlers) Delete(c echo.Context) error {
	var id = c.QueryParam("id")

	if err := rc.chUC.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"challenge": "deleted succesfully"})
}

func (rc *ChallengeHandlers) GetByID(c echo.Context) error {
	var id = c.QueryParam("id")

	challenge, err := rc.chUC.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"data": challenge})
}
