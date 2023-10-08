package route

import (
	"github.com/abishz17/go-backend-template/bootstrap"
	"github.com/labstack/echo/v4"
)

func Setup(app *bootstrap.Application, echo *echo.Echo) {
	groups := echo.Group("/api")
	NewSignUpRoute(groups, app.DataBase, app.Env)
}
