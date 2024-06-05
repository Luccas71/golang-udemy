package main

import "fmt"

// init => vai ser a primeira função executada, antes mesmo da main
// posso ter uma init por arquivo
//muito usada para setup

func init() {
	fmt.Println("Executando funcao init")
}

func main() {
	fmt.Println("Executando funcao main")
}
