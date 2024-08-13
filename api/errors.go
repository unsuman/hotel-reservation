package api

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(Error); ok {
		return c.Status(apiError.Code).JSON(apiError)
	}
	apiError := NewError(fiber.StatusInternalServerError, err.Error())
	return c.Status(apiError.Code).JSON(apiError)
}

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

// Error implements the Error interface
func (e Error) Error() string {
	return e.Err
}

func NewError(code int, err string) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

func ErrUnAuthorized() Error {
	return Error{
		Code: fiber.StatusUnauthorized,
		Err:  "unauthorized request",
	}
}

func ErrNotResourceNotFound(res string) Error {
	return Error{
		Code: fiber.StatusNotFound,
		Err:  res + " resource not found",
	}
}

func ErrBadRequest() Error {
	return Error{
		Code: fiber.StatusBadRequest,
		Err:  "invalid JSON request",
	}
}

func ErrInvalidID() Error {
	return Error{
		Code: fiber.StatusBadRequest,
		Err:  "invalid id given",
	}
}
