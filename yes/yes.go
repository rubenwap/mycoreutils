package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func sayYes(expletive string, ch chan string) {
	ch <- expletive
}

func yes() *cli.App {
	app := cli.NewApp()
	app.Name = "yes"
	app.Usage = "yes outputs expletive, or, by default, ``y'', forever"
	app.Action = func(c *cli.Context) error {

		yesChan := make(chan string)

		expletive := c.Args().Get(0)
		if expletive == "" {
			expletive = "y"
		}

		for {
			go func() {
				sayYes(expletive, yesChan)
			}()
			fmt.Println(<-yesChan)
		}

	}
	return app
}

func main() {
	app := yes()
	app.Run(os.Args)
}
