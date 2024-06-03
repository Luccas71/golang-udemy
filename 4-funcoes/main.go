package main

import "fmt"

func main() {
	soma := chamarFuncSomar(1, 5)
	fmt.Println(soma)

	// usar _ para ignorar algum retorno da função
	// resultSoma, _ := calculos(12, 23)

	resultSoma, resultSub := calculos(12, 23)
	fmt.Println(resultSoma, resultSub)
}

// declaração de uma função
func somar(n1, n2 int) int {
	return n1 + n2
}

func chamarFuncSomar(n1, n2 int) int {
	return somar(n1, n2)
}

// retornando mais de um valor
func calculos(n1, n2 int) (int, int) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}
