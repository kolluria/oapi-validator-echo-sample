package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	servergen "github.com/oapi-validator-echo-sample/server-gen"
	"github.com/oapi-validator-echo-sample/utils"
)

// ErrorHandler is a custom error handler for the Server
func ErrorHandler(err error, c echo.Context) {
	e := c.Echo()

	// If the response has already been written to the client, then the error might have been returned by one of the middlewares.
	// In that case, we don't have to do anything.
	// This idea is derived from echo.DefaultHTTPErrorHandler.
	if c.Response().Committed {
		e.Logger.Info("Response already committed. Ignoring the error - ", err)
		return
	}

	code := int32(http.StatusInternalServerError)
	msg := "Internal Server Error"
	respType := "error"

	// TODO: unwrap the error and set the code, message and type accordingly
	err = errors.Unwrap(err)
	if t, ok := err.(utils.Error); ok {
		code = t.Code()
		msg = t.Error()
	}

	err = c.JSON(int(code), servergen.ApiResponse{
		Code:    &code,
		Message: &msg,
		Type:    &respType,
	})
	if err != nil {
		e.Logger.Error(err, "Error writing error response")
	}
}
