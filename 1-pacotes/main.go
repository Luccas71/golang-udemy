package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo da função main")
	auxiliar.Escrever()
	err := checkmail.ValidateFormat("email")
	fmt.Println(err)
}
