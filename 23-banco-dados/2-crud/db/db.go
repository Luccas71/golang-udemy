package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/go-udemy?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConexao)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
