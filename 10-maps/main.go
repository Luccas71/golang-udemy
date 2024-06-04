package main

import "fmt"

func main() {
	//Maps => estrutura de chave-valor
	// map1 = map[tipo da chave]tipo do valor{}

	map1 := map[string]int{
		"nota1": 7,
		"nota2": 3,
	}
	fmt.Println(map1)

	//um map pode receber outro map
	map2 := map[int]map[string]string{
		1: {
			"nome":      "Lucas",
			"sobrenome": "Maciel",
		},
		2: {
			"nome":      "Valeria",
			"sobrenome": "Maciel",
		},
	}
	fmt.Println(map2)
	//removendo uma chave do map
	delete(map2, 1)
	fmt.Println(map2)

	//posso adicionar infos, desde que respeite a estrutura do map
	map2[3] = map[string]string{"signo": "gemeos"}
	fmt.Println(map2)
}
