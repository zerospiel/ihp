package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/urfave/cli/v2"
	"github.com/zerospiel/ihp/internal/cmds/git"
	"github.com/zerospiel/ihp/internal/cmds/ing"
)

// source: https://www.w3.org/TR/2016/REC-html51-20161101/sec-forms.html#email-state-typeemail
const emailReStr = `^[a-zA-Z0-9.!#$%&'*+\/=?^_\x60{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`

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
	git := &cli.Command{
		Name:    "git",
		Aliases: []string{"g"},
		Usage:   "init git if not initiaziled and setup local config with a given uname and email",
		Action: func(ctx *cli.Context) error {
			var (
				emailRe = regexp.MustCompile(emailReStr)

				email, uname string
			)
			if len(os.Args) > 2 {
				for _, arg := range os.Args[2:] {
					if email != "" && uname != "" {
						break
					}
					c := checkMail(emailRe, arg)
					if email == "" && c {
						email = arg
						continue
					}
					if uname == "" && !c {
						uname = arg
					}
				}
			} else {
				email = "morgoevm@gmail.com"
				uname = "zerospiel"
			}
			return git.Setup(email, uname)
		},
	}
	return []*cli.Command{ig, git}
}

func checkMail(re *regexp.Regexp, s string) bool {
	if len(s) < 3 || len(s) > 254 {
		return false
	}
	return re.MatchString(s)
}
