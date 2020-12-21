package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/zerospiel/ihp/internal/cmds/ing"
)

func listCmds() []*cli.Command {
	ig := &cli.Command{
		Name:    "girl",
		Aliases: []string{"gg"}, // get girl
		Usage:   "put an `URL` to a post, will open the raw photo :)",
		Action: func(ctx *cli.Context) error {
			if len(os.Args) != 3 {
				return fmt.Errorf("expected 'ihp girl URL', got '%+v'", os.Args)
			}
			return ing.OpenLink(os.Args[2])
		},
	}
	return []*cli.Command{ig}
}
