package main

import "fmt"

// função que chama ela mesma
// pode ocorrer estouro de pilha (stack overflow)
// precisa de um condição de parada
// nao utilizar com numeros muito grandes

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
	// 1 1 2 3 5 8 13 21 34 55
}
func main() {
	posicao := fibonacci(7)
	fmt.Println(posicao)
}
