package main

import "fmt"

func main() {

	var variavel1 int = 10
	var variavel2 int = variavel1

	//alterar o valor de variavel1 nao vai alterar a variavel2, pois ela é apenas uma cópia

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//ponteiro é uma referencia na memoria
	var variavel3 int = 100

	//ponteiro recebe o espaço na memoria da variavel3, agora os dois compartilham o mesmo valor e nao mais uma copia
	var ponteiro *int = &variavel3
	//o & faz referencia ao espaço na memoria

	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	//o * mostra o valor que está armazenado neste espaço de memoria
	fmt.Println(variavel3, *ponteiro)
}
