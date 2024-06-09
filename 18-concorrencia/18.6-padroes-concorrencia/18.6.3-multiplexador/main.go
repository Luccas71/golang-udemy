package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Ola Mundo"), escrever("Ola Go"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// junta dois ou mais canais em um canal só
func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()
	return canalSaida
}

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
