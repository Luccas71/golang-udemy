package main

import (
	"crud/servidor"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

	fmt.Println("Rodando na porta 8000")
	http.ListenAndServe(":8000", router)
}
