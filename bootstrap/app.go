package bootstrap

import (
	"github.com/abishz17/go-backend-template/infrastructure"
	"github.com/abishz17/go-backend-template/internal/api/repository"
)

type Application struct {
	Env      *infrastructure.Env
	DataBase *repository.DataBase
}

func App() Application {
	app := &Application{}
	env := infrastructure.NewEnv()
	dbConn := infrastructure.NewDBConn(env)
	app.DataBase = repository.NewDataBase(dbConn)
	app.Env = &env
	return *app
}
