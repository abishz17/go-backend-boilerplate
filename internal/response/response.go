package response

import (
	"errors"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type Response struct {
	ErrorMessage string `json:"message,omitempty"`
	Data         any    `json:"data,omitempty"`
	Error        bool   `json:"error"`
}

func SuccessResponse(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, Response{
		Data:  data,
		Error: false,
	})
}

func ServerErrorResponse(ctx echo.Context, err error) error {
	return ctx.JSON(http.StatusInternalServerError, Response{
		ErrorMessage: "Something unusual happened",
		Error:        true,
	})
}

func BadRequestResponse(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusBadRequest, Response{
		ErrorMessage: message,
		Error:        true,
	})
}

func NotFoundResponse(ctx echo.Context, message string) error {
	return ctx.JSON(http.StatusNotFound, Response{
		Error:        true,
		ErrorMessage: message,
	})
}

func ErrorResponse(ctx echo.Context, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NotFoundResponse(ctx, "record not found")
	} else if errors.Is(err, gorm.ErrDuplicatedKey) {
		return BadRequestResponse(ctx, "duplicate data")
	}
	if errors.As(err, &AppError{}) {
		return BadRequestResponse(ctx, err.Error())
	}
	return ServerErrorResponse(ctx, err)
}
