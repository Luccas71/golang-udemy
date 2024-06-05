package main

import "fmt"

/*
	panic =>
		mata a execução do programa,
		nada é executado apos o panic,
		vai chamar todos defer antes de parar,
	recover =>
		recuperar a execução da aplicação apos um panic,
		continua a execução e retorno o zerovalue do tipo do retorno,
*/

func media(n1, n2 float64) bool {
	defer recuperarExecucao()
	fmt.Println("Aprovado?")
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	}
	if media < 6 {
		return false
	}
	panic("Exatamente igual a 6")
}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recuperada com sucesso")
	}
}

func main() {
	fmt.Println(media(6, 5))
}
