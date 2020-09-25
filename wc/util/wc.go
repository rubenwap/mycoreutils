package util

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
)

// Wc is the main function that triggers the counts
func Wc() *cli.App {
	app := cli.NewApp()
	app.Name = "WC"
	app.Usage = "The wc utility displays the number of lines, words, and bytes contained in each input file"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "c",
			Usage: "The number of bytes in each input file is written to the standard output",
		},
		cli.BoolFlag{
			Name:  "l",
			Usage: "The number of lines in each input file is written to the standard output",
		},
		cli.BoolFlag{
			Name:  "m",
			Usage: "The number of characters in each input file is written to the standard output",
		},
		cli.BoolFlag{
			Name:  "w",
			Usage: "The number of words in each input file is written to the standard output",
		},
	}
	app.Action = func(c *cli.Context) error {

		var buf bytes.Buffer
		var m map[string]int
		m = make(map[string]int)

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			buf.WriteString(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
		}

		for i := range c.Args() {
			fmt.Print(i)
			content, err := ioutil.ReadFile(c.Args().Get(i))
			if err != nil {
				log.Fatal(err)
			}
			buf.WriteString(string(content))
		}

		if c.Bool("c") {
			m["clen"] = byteCounts(buf.String())
		}

		if c.Bool("l") {
			m["llen"] = lineCounts(buf.String())
		}

		if c.Bool("m") {
			m["mlen"] = characterCounts(buf.String())
		}

		if c.Bool("w") {
			m["wlen"] = wordCounts(buf.String())
		}

		if c.Bool("c") && c.Bool("m") {
			m["mlen"] = 0
		}

		for _, value := range m {
			if value != 0 {
				fmt.Print("\t", value)
			}
		}
		fmt.Println("")

		return nil
	}
	return app
}
