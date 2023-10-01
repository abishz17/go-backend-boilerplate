package main

import (
	"github.com/abishz17/go-backend-template/bootstrap"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	e := echo.New()
	err := e.Start(":" + env.ServerPort)
	if err != nil {
		log.Fatal("Cannot spin up the server", err)
	}
}
