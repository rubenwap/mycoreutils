package main

import (
	"bytes"
	"github.com/urfave/cli"
	"os"
	"bufio"
)

func yes() *cli.App {
	app := cli.NewApp()
	app.Name = "yes"
	app.Usage = "yes outputs expletive, or, by default, ``y'', forever"
	app.Action = func(c *cli.Context) error {
    expletive := c.Args().Get(0)
		if expletive == "" {
			expletive = "y"
		}

		for {
			f := bufio.NewWriterSize(os.Stdout, 8192)
			f.Write(bytes.Repeat([]byte(expletive+"\n"), 100000000))
		}
	}
	return app
}

func main() {
	app := yes()
	app.Run(os.Args)
}
