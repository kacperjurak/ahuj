package handler

import (
	"ahuj/internal/domain/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type request struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type result struct {
	Sum int `json:"sum""`
}

type CalculateHTTPHandler struct {
	CalculatorService service.Calculator
}

func (h *CalculateHTTPHandler) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	req := &request{}

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	sum, err := h.CalculatorService.Add(ctx, req.X, req.Y)
	if err != nil {
		log.Errorf("calculator servie add error: %s", err)
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, result{
		Sum: sum,
	})
}
