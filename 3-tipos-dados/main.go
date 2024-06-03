package main

import (
	"errors"
	"fmt"
)

func main() {

	// INT
	// int, int8, int16, int32, int64 => diferença está na quantidade de bits

	//uint => nao suporta numeros negativos
	// alias = apelido
	// rune = int32
	// byte = int8

	//FLOAT
	// float = numero real => 1.25

	// STRING
	var str string = "bola"
	fmt.Println(str)

	str2 := "Botao"
	fmt.Println(str2)

	//O char irá retornar o numero correspondente na tabela ASCII
	char := '^'
	fmt.Println(char)

	//valor zero => valor atribuido a uma variavel de forma padrão
	// float e int = 0
	// bool = false
	// string = ""
	// erro = nil (nulo)

	//BOOL
	booleano := true
	fmt.Println(booleano)

	//ERROR
	var err error = errors.New("Erro interno")
	fmt.Println(err)
}
