package main

import "fmt"

//funcao nao nomeada
//pode receber paramatros
//pode ter retorno armazenado em uma var
// func () {} ()

func main() {

	func() {
		fmt.Println("Ola mundo")
	}()

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Lucas")
	fmt.Println(retorno)
}
