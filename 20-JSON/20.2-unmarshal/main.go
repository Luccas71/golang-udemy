package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` // `json:"-"` ignora o campo
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Rex","raca":"vira-lata","idade":3}`

	//JSON em struct
	var c cachorro
	if err := json.Unmarshal([]byte(cachorroEmJSON), &c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	//JSON em map
	cachorroEmJSON2 := `{"nome":"Alex","raca":"fila"}`
	c2 := make(map[string]string)
	if err := json.Unmarshal([]byte(cachorroEmJSON2), &c2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)

}
