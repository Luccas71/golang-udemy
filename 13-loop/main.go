package main

import (
	"fmt"
)

func main() {
	//FOR

	//"while"
	// i := 0
	// for i <= 10 {
	// 	time.Sleep(time.Second)
	// 	fmt.Println(i)
	// 	i += 2
	// }

	//for normal kk
	// for i := 0; i <= 6; i += 2 {
	// 	time.Sleep(time.Second)
	// 	fmt.Println(i)
	// }

	//range => iterando por slice/array/map (forEach)
	nomes := []string{"Dan", "Carlos", "Adao"}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	//iterando sobre uma string (sempre vai retornar ref tabela ASCII)
	for _, letra := range "Lucas" {
		fmt.Print(string(letra) + " ")
	}

}
