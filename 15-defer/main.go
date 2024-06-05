package main

import "fmt"

/*
	adia a execução até o ultimo momento possivel
	dentro de um bloco, será executado antes do return
*/

func funcao1() {
	fmt.Println("Função 1")
}
func funcao2() {
	fmt.Println("Função 2")
}
func media(n1, n2 float64) bool {
	defer fmt.Println("Retornando situação")
	fmt.Println("Calculando média")

	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	// defer funcao1()
	// funcao2()

	fmt.Println(media(4, 6))
}
