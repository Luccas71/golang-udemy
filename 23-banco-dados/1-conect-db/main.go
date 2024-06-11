package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//string de conexao com db
	stringConnect := "golang:golan@/go-udemy?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConnect)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//verificando se a conexão está aberta
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão aberta!")
}
