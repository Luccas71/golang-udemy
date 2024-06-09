package main

import (
	"fmt"
	"teste-intro/endereco"
)

func main() {
	tipoEndereco := endereco.TipoDeEndereco("bairro dos anjos")
	fmt.Println(tipoEndereco)
}
