package main

import (
	"log"
	"os"

	"github.com/Luccas71/golang-udemy/17-app-cli/app"
)

func main() {
	err := app.Gerar().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
