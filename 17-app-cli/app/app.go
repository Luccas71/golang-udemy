package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IP e nomes de servidores na internet"
	//separar flags numa variavel, pois sera reutilizada
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			//valor padrao se nada for passado
			Value: "google.com.br",
		},
	}
	app.Commands = []cli.Command{
		{
			//comando de buscar ip
			Name:  "ip",
			Usage: "Buscar IPs de endereços na internt",
			Flags: flags,
			// recebe uma funcao
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "buscar servidor onde site está hospedado",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	//retornar o valor da flag host
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
