package util

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

// Wc is the main function
func Wc() *cli.App {
	app := cli.NewApp()
	app.Name = "WC"
	app.Usage = "The wc utility displays the number of lines, words, and bytes contained in each input file"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "c",
			Usage: "The number of bytes in each input file is written to the standard output",
		},
		cli.BoolFlag{
			Name: "l",
			Usage: "The number of lines in each input file is written to the standard output",
		},
		cli.BoolFlag{
			Name: "m",
			Usage: "The number of characters in each input file is written to the standard output",
		},
		cli.BoolFlag{
			Name: "w",
			Usage: "The number of words in each input file is written to the standard output",
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println(os.Args)
		return nil
	}
	return app
}