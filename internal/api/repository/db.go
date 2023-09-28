package repository

import (
	"github.com/abishz17/go-backend-template/bootstrap"
)

type DataBase struct {
	conn bootstrap.DBConn
}

func NewDataBase(conn bootstrap.DBConn) *DataBase {
	return &DataBase{
		conn: conn,
	}
}
