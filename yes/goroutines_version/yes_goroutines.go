/*

This version was done as a proof of concept, since goroutines are not helping in the use case of this
tool, because of the serial nature of the IO operations. This was more to see how semaphore works in
high throughput operations

*/


package main

import (
	"github.com/urfave/cli"
	"os"
	"runtime"
	"context"
	"log"
	"golang.org/x/sync/semaphore"
	"bytes"
)

var (
    maxWorkers = runtime.GOMAXPROCS(0)
    sem        = semaphore.NewWeighted(int64(maxWorkers))
	yesChan    = make(chan []byte)
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
				 yesChan <- bytes.Repeat([]byte(expletive), 100000000)
			 }()
			 go func() {
			 	defer sem.Release(2)
			 	os.Stdout.Write(<-yesChan)
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
