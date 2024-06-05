package main

import "fmt"

// pode receber n argumentos
// só permite um parametro variatico por func
// deve obrigatoriamente ser o ultimo parametro
func somar(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(somar(1, 2, 3, 4, 10))

	escrever("Oi", 1, 2, 3)
}
