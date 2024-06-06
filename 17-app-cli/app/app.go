package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IP e nomes de servidores na internet"
	app.Commands = []cli.Command{
		{
			//comando de buscar ip
			Name:  "ip",
			Usage: "Buscar IPs de endereços na internt",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					//valor padrao se nada for passado
					Value: "devbook.com.;br",
				},
			},
			// recebe uma funcao
			Action: buscarIps,
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
		println(ip)
	}
}
