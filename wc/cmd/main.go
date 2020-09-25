package main

import (
	util "github.com/rubenwap/mycoreutils/wc/util"
	"os"
)

func main() {
	app := util.Wc()
	app.Run(os.Args)
}
