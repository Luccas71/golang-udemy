package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	p1 := pessoa{"Lucas", 21}
	e1 := estudante{p1, "Eng", "redentor"}
	fmt.Println(e1)

	e2 := estudante{pessoa{"Carlos", 22}, "Med", "ufv"}
	fmt.Println(e2)

}
