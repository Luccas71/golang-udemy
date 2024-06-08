package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//executando 2 ou mais goroutines em concorrencia sem quebar o programa

	var waitGroup sync.WaitGroup // modo menos utilizado

	//quant.de routines em espera
	waitGroup.Add(4)

	go func() {
		escrever("Goroutine 1")
		waitGroup.Done()
	}()
	go func() {
		escrever("Goroutine 2")
		waitGroup.Done()
	}()
	go func() {
		escrever("Goroutine 3")
		waitGroup.Done()
	}()
	go func() {
		escrever("Goroutine 4")
		waitGroup.Done()
	}()

	//espera as routines serem executadas
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
