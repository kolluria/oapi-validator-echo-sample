package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoutils "github.com/oapi-validator-echo-sample/utils/echo"
)

func main() {
	ctx := context.TODO()
	s := NewEchoServer(ctx,
		echoutils.RequestIDWithConfig(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}),
	)
	s.Logger.Fatal(s.Start(":8080"))
}
