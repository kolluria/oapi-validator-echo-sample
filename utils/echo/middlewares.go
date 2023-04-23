package echo

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oapi-validator-echo-sample/utils"
)

// RequestIDWithConfig returns a middleware that injects a request ID into the context of each request.
// It also sets the TraceID key in context's store which will be used by the logger to trace the request.
func RequestIDWithConfig() echo.MiddlewareFunc {
	return middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		RequestIDHandler: func(e echo.Context, s string) {
			e.Set(utils.TraceIDKey, s)
		},
		Generator: func() string {
			return uuid.New().String()
		},
		TargetHeader: echo.HeaderXRequestID,
	})
}
