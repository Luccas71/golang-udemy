package main

import "fmt"

func main() {
	canal := make(chan string, 2) //tamanho do buffer
	// só vai bloquear quando atingir a cap do buffer
	// sem buffer, sempre vai bloquear
	canal <- "Ola mundo!"
	mensagem := <-canal
	fmt.Println(mensagem)
}
