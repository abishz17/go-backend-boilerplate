package bootstrap

import "github.com/abishz17/go-backend-template/internal/api/repository"

type Application struct {
	Env      *Env
	DataBase *repository.DataBase
}

func App() Application {
	app := &Application{}
	env := NewEnv()
	dbConn := NewDBConn(env)
	app.DataBase = repository.NewDataBase(dbConn)
	return *app
}
