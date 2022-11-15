package main

import (
	"log"
	"os"

	"github.com/Ricardo-Sales/my-cli-app/app"
)

func main() {
	aplication := app.Generate()
	if err := aplication.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
