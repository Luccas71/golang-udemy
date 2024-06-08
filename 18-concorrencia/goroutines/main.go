package main

import (
	"fmt"
	"time"
)

//CONCORRENCIA != PARALELISMO
//A goroutine nao vai esperar que a execução de uma funcao acabe para chamar a outra funcao

func main() {
	go escrever("Ola mundo") //goroutine
	escrever("Ola Go!")
	//executar tudo com goroutine não irá executar nada!
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
