package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/urfave/cli"
)

var wg sync.WaitGroup

func byteCounts(text string) int {
	return len(text)
}

func lineCounts(text string) int {
	return len(strings.Split(text, "\n"))
}

func characterCounts(text string) int {
	return utf8.RuneCountInString(text)
}

func wordCounts(text string) int {
	return len(strings.Fields(text))
}

func wc() *cli.App {
	app := cli.NewApp()
	app.Name = "wc"
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

		bytesChan := make(chan int)
		linesChan := make(chan int)
		charactersChan := make(chan int)
		wordChan := make(chan int)

		wg.Add(4)
		go func(text string) {
			bytesChan <- byteCounts(text)
			wg.Done()
		}(buf.String())
		go func(text string) {
			linesChan <- lineCounts(text)
			wg.Done()
		}(buf.String())
		go func(text string) {
			charactersChan <- characterCounts(text)
			wg.Done()
		}(buf.String())
		go func(text string) {
			wordChan <- wordCounts(text)
			wg.Done()
		}(buf.String())

		for i := 0; i < 4; i++ {
			select {
			case msg1 := <-bytesChan:
				m["c"] = msg1
				m["m"] = 0
			case msg2 := <-linesChan:
				m["l"] = msg2
			case msg3 := <-charactersChan:
				m["m"] = msg3
				m["c"] = 0
			case msg4 := <-wordChan:
				m["w"] = msg4
			}
		}

		wg.Wait()

		for key, value := range m {
			if c.Bool(key) && value != 0 {
				fmt.Print("\t", key, ": ", value)
			}
		}
		fmt.Println("")

		return nil
	}

	return app
}

func main() {
	app := wc()
	app.Run(os.Args)
}
