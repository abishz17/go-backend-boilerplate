package route

import (
	"github.com/abishz17/go-backend-template/infrastructure"
	handler2 "github.com/abishz17/go-backend-template/internal/api/handler"
	repository2 "github.com/abishz17/go-backend-template/internal/api/repository"
	service2 "github.com/abishz17/go-backend-template/internal/api/service"
	"github.com/labstack/echo/v4"
)

func NewSignUpRoute(group *echo.Group, db *repository2.DataBase, env *infrastructure.Env) {
	repository := repository2.NewUserRepository(db)
	service := service2.NewUserService(repository, env)
	handler := handler2.NewUserHandler(service)
	group.POST("/sign-up", handler.Create)
	group.POST("/login", handler.Login)
}
