package main

import (
	"ahuj/internal/application/service"
	"ahuj/internal/infrastructure/adder/internaladder"
	"ahuj/internal/infrastructure/database"
	"ahuj/internal/infrastructure/database/repository"
	"ahuj/internal/infrastructure/httpserver/handler"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
)

func main() {
	c := NewContainer()

	e := echo.New()
	e.HideBanner = true
	e.Debug = true
	e.Validator = &CustomValidator{validator: validator.New()}

	api := e.Group("/api")

	calculator := api.Group("/calculator")
	calculator.GET("/add", c.CalculateHTTPHandler.Handle)

	api.GET("/history", c.HistoryHTTPHandler.Handle)

	log.Println("AHUJ started!")

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))))
}

type Container struct {
	CalculateHTTPHandler *handler.CalculateHTTPHandler
	HistoryHTTPHandler   *handler.HistoryHTTPHandler
}

func NewContainer() *Container {
	adderRepository := &internaladder.Adder{}
	//adderRepository := &restadder.Adder{}
	resultRepository := repository.NewMySQLResultRepository(database.NewMySQLConnection())
	//resultRepository := repository.NewInMemory()

	calculator := &service.Calculator{
		AdderRepository:  adderRepository,
		ResultRepository: resultRepository,
	}

	calculateHTTPHandler := &handler.CalculateHTTPHandler{
		CalculatorService: calculator,
	}

	resultsHTTPHandler := &handler.HistoryHTTPHandler{
		CalculatorService: calculator,
	}

	return &Container{
		CalculateHTTPHandler: calculateHTTPHandler,
		HistoryHTTPHandler:   resultsHTTPHandler,
	}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
