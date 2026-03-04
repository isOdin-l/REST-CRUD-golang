package errors

import (
	"fmt"
	"isOdin/RestApi/pkg/api"
	"net/http"

	"github.com/labstack/echo/v5"
)

type AppError struct {
	HttpCode int
	Message  string
	Err      error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewAppError(httpCode int, message string, err error) *AppError {
	return &AppError{HttpCode: httpCode, Message: message, Err: err}
}

var (
	ErrBadRequest          = NewAppError(http.StatusBadRequest, "bad request", nil)
	ErrUnauthorized        = NewAppError(http.StatusUnauthorized, "unauthorized access", nil)
	ErrForbidden           = NewAppError(http.StatusForbidden, "access forbidden", nil)
	ErrNotFound            = NewAppError(http.StatusNotFound, "resource not found", nil)
	ErrConflict            = NewAppError(http.StatusConflict, "resource conflict", nil)
	ErrInternalServerError = NewAppError(http.StatusInternalServerError, "internal server error", nil)
	ErrValidation          = NewAppError(http.StatusBadRequest, "validation failed", nil)
)

func NewValidationError(err error) *AppError {
	return NewAppError(http.StatusBadRequest, "validation failed", err)
}

func NewInternalError(err error) *AppError {
	return NewAppError(http.StatusInternalServerError, "internal server error", err)
}

func FromToErrorApi(err *AppError) *api.ResponseError {
	return &api.ResponseError{
		Error: struct {
			HttpCode int    "json:\"code\""
			Message  string "json:\"message\""
		}{
			HttpCode: err.HttpCode,
			Message:  err.Message,
		},
	}
}

func ResponseError(c *echo.Context, err *AppError) error {
	return c.JSON(err.HttpCode, FromToErrorApi(err))
}
