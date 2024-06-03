package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint
}

func main() {
	//"instanciando" uma struct
	var u usuario
	u.nome = "Lucas"
	u.idade = 12
	fmt.Println(u)

	//inferencia de tipo
	u2 := usuario{"Carlos", 45, endereco{"Rua 1", 12}}
	fmt.Println(u2)

	//ignorando algum valor
	u3 := usuario{nome: "Andre"}
	fmt.Println(u3)

	// uma struct dentro de outra
}
