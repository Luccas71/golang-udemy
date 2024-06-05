package main

// mantem as variaveis criadas na funçao permitindo que sejam usadas
import "fmt"

func closure() func() {
	texto := "Dentro da funcao closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}
func main() {
	texto := "Alo"
	fmt.Println(texto)

	funcClosure := closure()
	funcClosure()
}
