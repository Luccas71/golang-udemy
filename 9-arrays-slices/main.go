package main

import "fmt"

func main() {
	//arrays tem seu tamanho definido na sua declaração e este tamanho é imutável
	var array1 [2]int
	array1[0] = 3
	fmt.Println(array1)

	//inferencia
	array2 := [2]int{12, 9}
	fmt.Println(array2)

	//tamanho baseado na quant de itens passados na declaração
	//nao torna o array dinamico
	array3 := [...]int{1, 2, 3}
	fmt.Println(array3)
	// array3[4] = 12 nao é permitido
}
