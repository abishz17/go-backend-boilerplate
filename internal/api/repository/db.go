package repository

import (
	"github.com/abishz17/go-backend-template/infrastructure"
)

type DataBase struct {
	conn infrastructure.DBConn
}

func NewDataBase(conn infrastructure.DBConn) *DataBase {
	return &DataBase{
		conn: conn,
	}
}
