package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name:        "ihp",
		Description: "helps me to save 10 mins per year",
		Version:     "v0.0.1",
		Commands:    listCmds(),
		Compiled:    time.Now(),
		Authors:     []*cli.Author{{Name: "zerospiel", Email: "ww@bk.ru"}},
		Copyright:   "🤷‍♂️",
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
