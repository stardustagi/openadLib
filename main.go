package main

import (
	"adApi/core/logger"
	"github.com/envoyproxy/protoc-gen-validate/validate"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"sync"
)

func runServer(e *echo.Echo, wg *sync.WaitGroup) {
	defer wg.Done()
	e.Validator = &validate.CustomValidator{Validator: validator.New()}

	e.Logger.Fatal(e.Start(":1323"))
}

func main() {
	logger.NewLogger()
}
