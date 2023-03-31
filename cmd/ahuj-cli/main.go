package main

import (
	"ahuj/internal/application/service"
	"ahuj/internal/infrastructure/adder/internaladder"
	"ahuj/internal/infrastructure/database"
	"ahuj/internal/infrastructure/database/repository"
	"context"
	"flag"
	"fmt"
)

func main() {
	x := flag.Int("x", 0, "x number")
	y := flag.Int("y", 0, "y number")
	flag.Parse()

	c := NewContainer()

	sum, err := c.CalculatorServie.Add(context.Background(), *x, *y)
	if err != nil {
		fmt.Printf("calculator servie add error: %s\n", err)
	}

	fmt.Printf("sum is: %d\n", sum)
}

type Container struct {
	CalculatorServie *service.Calculator
}

func NewContainer() *Container {
	adderRepository := &internaladder.Adder{}
	resultRepository := repository.NewMySQLResultRepository(database.NewMySQLConnection())

	calculator := &service.Calculator{
		AdderRepository:  adderRepository,
		ResultRepository: resultRepository,
	}

	return &Container{
		CalculatorServie: calculator,
	}
}
