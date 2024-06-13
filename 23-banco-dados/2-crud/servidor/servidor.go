package servidor

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"net/http"
)

type usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter JSON"))
		return
	}

	db, err := db.Conectar()
	if err != nil {
		w.Write([]byte("Erro ao conectar ao banco de dados"))
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into usuarios (name, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Erro ao inserir dados no banco"))
	}
	defer stmt.Close()

	insercao, err := stmt.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar statement"))
	}

	//retornando o id do usuario inserido
	lastID, err := insercao.LastInsertId()
	if err != nil {
		w.Write([]byte("Não foi possível retornar o ID"))
	}

	//STATUS CODE
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", lastID)))
}
