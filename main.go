package main

import (
	"github.com/urfave/cli"
	"github.com/sanguohot/ethereum-abi/abi"
	"log"
	"fmt"
	"os"
	"time"
)

func InitApp() error {
	app := cli.NewApp()
	app.Name = "ethereum-abi"
	app.Usage = "command line for ethereum-abi!"
	app.Version = "1.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Sanguohot",
			Email: "hw535431@163.com",
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Printf("%s-%s", app.Name, app.Version)
		fmt.Printf("\n%s", app.Usage)
		return nil
	}
	app.Commands = *abi.AbiCommands

	return app.Run(os.Args)
}

func main() {
	err := InitApp()
	if err != nil {
		log.Fatal(err)
	}
}