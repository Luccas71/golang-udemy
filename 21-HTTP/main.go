package main

import (
	"log"
	"net/http"
)

func main() {

	//rotas
	http.HandleFunc("/home", home)
	http.HandleFunc("/users", users)
	//iniciando servidor
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola mundo!"))
}
func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar usuários..."))
}
