package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "vira-lata", 3}
	cachorro_JSON, err := json.Marshal(c) // transforma struct/map em JSON (retorna []byte)
	if err != nil {
		log.Fatal(err)
	}
	c2 := map[string]string{"Nome": "Bob", "Raca": "buldog"}
	cachorro_JSON2, err := json.Marshal(c2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cachorro_JSON, cachorro_JSON2)
	fmt.Println(bytes.NewBuffer(cachorro_JSON))  // retorna em formato JSON
	fmt.Println(bytes.NewBuffer(cachorro_JSON2)) // retorna em formato JSON
}
