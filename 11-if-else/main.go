package main

import "fmt"

func main() {
	numero := 13
	if numero < 7 {
		fmt.Println("Menor do que 7")
	} else {
		fmt.Println("Maior que 7")
	}

	//if init => declara uma variavel com if
	//a variavel só existe dentro do if
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("O numero é maior que zero")
	}

	// else if
	if numero > 0 {
		fmt.Println("Numero maior que zero")
	} else if numero == 0 {
		fmt.Println("Numero é igual a zero")
	} else {
		fmt.Println("Numero é menor que zero")
	}

}
