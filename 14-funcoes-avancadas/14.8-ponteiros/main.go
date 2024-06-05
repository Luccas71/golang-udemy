package main

import "fmt"

// altera uma copia
func inverteSinal(numero int) int {
	numero = numero * -1
	return numero
}

// altera o valor na memoria, em todo o codigo
func inverteSinalPonteiro(numero *int) int {
	*numero = *numero * -1
	return *numero
}

func main() {
	numero := 20
	novoNumero := inverteSinal(numero) //nao altera o valor pois foi passada apenas uma copia
	fmt.Println(novoNumero)
	fmt.Println(numero)

	numero2 := 40
	fmt.Println(numero2)
	novoNumero2 := inverteSinalPonteiro(&numero2)
	fmt.Println(novoNumero2)
}
