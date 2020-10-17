package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/user"
)

func whoami() *cli.App {
	app := cli.NewApp()
	app.Name = "whoami"
	app.Usage = "display effective user id"
	app.Action = func(c *cli.Context) error {

		if usr, err := user.Current(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(usr.Username)
		}
		return nil
	}
	return app
}

func main() {
	app := whoami()
	app.Run(os.Args)
}
