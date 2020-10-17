package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"runtime"
	"context"
	"log"
	"strings"
	"golang.org/x/sync/semaphore"
)


var (
    maxWorkers = runtime.GOMAXPROCS(0)
    sem        = semaphore.NewWeighted(int64(maxWorkers))
    yesChan    = make(chan string)
)

func yes() *cli.App {
	ctx := context.TODO()
	app := cli.NewApp()
	app.Name = "yes"
	app.Usage = "yes outputs expletive, or, by default, ``y'', forever"
	app.Action = func(c *cli.Context) error {

		expletive := c.Args().Get(0)
		if expletive == "" {
			expletive = "y"
		}

		for {

			if err := sem.Acquire(ctx, 4); err != nil {
				log.Printf("Failed to acquire semaphore: %v", err)
				break
			}

			go func() {
				defer sem.Release(2)
				yesChan <- strings.Repeat(expletive+"\n", 100000000)
			}()
			go func() {
				defer sem.Release(2)
				fmt.Fprint(os.Stdout, <-yesChan)
			}()
		}
		return nil
	}
	return app
}

func main() {
	app := yes()
	app.Run(os.Args)
}
