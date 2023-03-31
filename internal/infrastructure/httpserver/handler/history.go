package handler

import (
	"ahuj/internal/domain/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strconv"
)

type historyResult struct {
	X   int `json:"x"`
	Y   int `json:"y"`
	Sum int `json:"sum"`
}

type HistoryHTTPHandler struct {
	CalculatorService service.Calculator
}

func (h *HistoryHTTPHandler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.FormValue("id")

	validate := validator.New()

	err := validate.Var(id, "required,numeric")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "id is required and should be an integer number")
	}

	atoi, err := strconv.Atoi(id)
	if err != nil {
		log.Errorf("strint to int error: %s", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	m, err := h.CalculatorService.History(ctx, atoi)
	if err != nil {
		log.Errorf("calculator service history error: %s", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, historyResult{
		X:   m.X,
		Y:   m.Y,
		Sum: m.Sum,
	})
}
