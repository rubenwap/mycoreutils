package util

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"io/ioutil"
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
		
		var buf bytes.Buffer
		var clen int
		var llen int
		var mlen int
		var wlen int
		
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		buf.WriteString(text)

		for i := range c.Args() {
			fmt.Print(i)
			content, err := ioutil.ReadFile(c.Args().Get(i))
			if err != nil {
				log.Fatal(err)
			}
		    buf.WriteString(string(content))
		}


		if c.Bool("c") {

			clen = byteCounts(buf.String())

		}

		if c.Bool("l") {

		}

		if c.Bool("m") {

			mlen = characterCounts(buf.String())

			
		}

		if c.Bool("w") {
			
		}

		if c.Bool("c") && c.Bool("m") {
			mlen = 0
		}
		
		fmt.Print(fmt.Sprintf("\t %d \t %d \t %d \t %d", clen, llen, mlen, wlen))
		return nil
	}
	return app
}