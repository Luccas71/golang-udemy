package main

import "fmt"

func main() {
	soma := 1 + 8
	sub := 2 - 7
	div := 12.0 / 5
	mult := 15 * 4
	restoDaDiv := 15 % 4

	fmt.Println(soma, sub, div, mult, restoDaDiv)

	// não é permitido nenhuma operação com tipos diferentes ente si
	// var numero1 int32 = 15
	// var numero2 int64 = 15
	// soma := numero1 + numero2 !!! NÃO

	//Operadores de atribuição
	var numero1 int = 12
	numero2 := 1.2
	fmt.Println(numero1, numero2)
	fmt.Println("-------------------------")

	//Operadores Unários
	numero1 = 0
	numero1++     // numero1 = numero1 + 1
	numero1 += 23 // numero1 = numero1 + 23
	numero1--     // numero1 = numero1 - 1
	numero1 -= 45 // numero1 = numero1 - 45
	//menos utilizados
	numero1 %= 3 // numero1 = numero1 % 2
	numero1 /= 2 // numero1 = numero1 / 2
	numero1 *= 3 // numero1 = numero1 * 3
	fmt.Println(numero1)
	fmt.Println("---------------------")

	//Operadores logicos => sempre retornam bool
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println("-----------------------------")

	//Operadores de Comparação => E (&&), OU (||), NOT (!)
	caso1 := false
	caso2 := true
	fmt.Println(caso1 && caso2) //match se os dois valores forem true
	fmt.Println(caso1 || caso2) // match se um dos valores for true
	fmt.Println(!caso1)         // inverte o valor booleano (true => false)
}
