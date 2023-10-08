package main

import (
	"github.com/abishz17/go-backend-template/bootstrap"
	"github.com/abishz17/go-backend-template/internal/api/route"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
	route.Setup(&app, e)
	err := e.Start(":" + env.ServerPort)
	if err != nil {
		log.Fatal("Cannot spin up the server", err)
	}
}
