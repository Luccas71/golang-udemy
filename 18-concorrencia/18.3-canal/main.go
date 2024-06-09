package main

import (
	"fmt"
	"time"
)

// CANAL => usado para enviar e receber dados de um mesmo tipo (string - string)
// DEADLOCK => o canal fica esperando um dado que nunca vai chegar

func main() {
	canal := make(chan string) //criando um canal que vai receber/enviar string
	go escrever("Ola mundo!", canal)

	// for {
	// 	mensagem, aberto := <-canal // canal recebendo o texto, o programa sempre vai esperar receber para continuar
	// 	if !aberto {                //condição para encerrar o loop apos fechar o canal
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim da execuação")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // enviando o texto para o canal
		time.Sleep(time.Second)
	}
	close(canal) // fechando o canal
}
