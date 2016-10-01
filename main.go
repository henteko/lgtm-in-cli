package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "lgtm"
	app.Version = "1.0.0"
	app.Usage = "get http://lgtm.in/g image url"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	app.Run(os.Args)
}
