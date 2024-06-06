package main

import "fmt"

// metodos => funcoes associadas a alguma estrutura (func, struct...)

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) mostrarDados() {
	fmt.Println("Mostrando dados do usuario")
	fmt.Printf("Nome: %s, Idade: %d\n", u.nome, u.idade)
}

// usar ponteiro para alterar dados de forma permanente
func (u *usuario) aumentarIdade() {
	u.idade++
}

func main() {
	// user1 := usuario{"Lucas", 20}
	// user1.mostrarDados()

	user2 := usuario{"Carlos", 15}
	user2.mostrarDados()
	user2.aumentarIdade()
	user2.mostrarDados()
}
