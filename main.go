package main

import (
	"fmt"
	"github.com/sanguohot/ethereum-abi/abi"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func InitApp() error {
	app := cli.NewApp()
	app.Name = "ethereum-abi"
	app.Usage = "command line for ethereum-abi!"
	app.Version = "1.0.1"
	app.Compiled = time.Now()
	app.Flags = abi.AbiFlags
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