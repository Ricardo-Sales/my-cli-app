package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}
	app := cli.NewApp()
	app.Name = "My First Cli"
	app.Usage = "Search for IP and server freom internet"
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search IP from internet",
			Flags:  flags,
			Action: SearchIp,
		},
		{
			Name:   "server",
			Usage:  "Search server from internet",
			Flags:  flags,
			Action: SearchServer,
		},
	}
	return app
}
func SearchIp(ctx *cli.Context) {
	host := ctx.String("host")
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}

}
func SearchServer(ctx *cli.Context) {
	host := ctx.String("host")
	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}
	for _, server := range servers {
		fmt.Println(server.Host)
	}

}
