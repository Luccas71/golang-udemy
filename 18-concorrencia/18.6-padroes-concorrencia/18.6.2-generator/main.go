package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Ola mundo")
	for {
		fmt.Println(<-canal)
	}
}

// generator => encapsula uma go routine e retorna um canal
//evita declarações na func main

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}

//  <- chan = canal que recebe
//  chan <- = canal que envia
