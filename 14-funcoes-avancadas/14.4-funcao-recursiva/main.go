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

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

	// fatorial(3) chama fatorial(2).
	// fatorial(2) chama fatorial(1).
	// fatorial(1) chama fatorial(0).
	// fatorial(0) retorna 1 (caso base).
	// fatorial(1) retorna 1 * 1 = 1.
	// fatorial(2) retorna 2 * 1 = 2.
	// fatorial(3) retorna 3 * 2 = 6.
}
func main() {
	// posicao := fibonacci(7)
	fac := factorial(3)
	fmt.Println(fac)
}
