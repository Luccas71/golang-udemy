package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
		// fallthrough => avança direto para a proxima condição
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Numero Invalido"
	}

	//passando a condição dentro do case:
	// case numero == 1: todo

	// Go NÃO exige a utilização do break no switch
}

func main() {
	fmt.Println(diaDaSemana(69))
}
