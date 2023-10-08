package handler

import (
	"github.com/abishz17/go-backend-template/internal/api/service"
	"github.com/abishz17/go-backend-template/internal/response"
	"github.com/abishz17/go-backend-template/internal/view"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler interface {
	Create(ctx echo.Context) error
	Login(ctx echo.Context) error
}

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{
		userService: service,
	}
}

func (u UserHandler) Create(ctx echo.Context) error {
	var userView view.UserCreateView
	err := ctx.Bind(&userView)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}
	user, err := u.userService.CreateUser(ctx, userView)
	if err != nil {
		//return echo.NewHTTPError(500, "Something unusual occured.")
		return ctx.JSON(http.StatusInternalServerError, "Something unusual happened")
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"user": user,
	})
}

func (u UserHandler) Login(ctx echo.Context) error {
	var userView view.UserLoginView
	err := ctx.Bind(&userView)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}
	tokens, err := u.userService.UserLogin(ctx, userView)
	if err != nil {
		return response.ErrorResponse(ctx, err)
	}
	return response.SuccessResponse(ctx, map[string]interface{}{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}
