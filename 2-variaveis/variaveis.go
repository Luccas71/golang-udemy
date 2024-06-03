package main

import "fmt"

func main() {
	//formas de declarar variaveis
	var variavel1 string = "variavel 1"
	fmt.Println(variavel1)

	//inferencia de tipos
	//define o tipo da variavel baseado no valor que ela recebe
	variavel2 := 12
	fmt.Println(variavel2)

	//declarando multiplas variavel]is
	variavel3, variavel4 := 12, "batata"
	fmt.Println(variavel3, variavel4)

	//O mesmo é válido para constante, porem seu valor é imutável

	//trocando o valor das variaveis
	n1, n2 := 12, 23
	fmt.Printf("valor da n1: %d e valor de n2: %d\n", n1, n2)
	n1, n2 = n2, n1
	fmt.Printf("valor da n1: %d e valor de n2: %d\n", n1, n2)
}
