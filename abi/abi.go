package abi

import (
	"io/ioutil"
	"fmt"
	"os"
	"github.com/urfave/cli"
	"errors"
	"path"
)

var abiFlags = []cli.Flag{
	cli.StringFlag{
		Name: "abi-json",
		Value: "",
		Usage: "abi json string",
	},
	cli.StringFlag{
		Name: "abi-file",
		Value: "",
		Usage: "abi file path, should be absolute path",
	},
}

func getAbiJson(c *cli.Context) (error, string) {
	var (
		abiJsonString string
		abiFilePath string
	)
	if c.String("abi-json") != "" {
		abiJsonString = c.String("abi-json")
	}else if c.String("abi-file") != "" {
		abiFilePath = c.String("abi-file")
	}else if os.Getenv("abi-file") != "" {
		// abi-json to os env is too big, only support file path
		abiFilePath = os.Getenv("abi-file")
	}
	if abiFilePath != "" {
		//abiFilePath = "E:/evan/goland/src/medichain/contracts/medi/build/EasyCns.abi"
		if file, err := ioutil.ReadFile(path.Join(abiFilePath)); err != nil {
			return err, "";
		}else {
			abiJsonString = string(file)
		}
	}
	if abiJsonString == "" && abiFilePath == "" {
		return errors.New("abi: params abi-file or abi-json required"), ""
	}
	return nil, abiJsonString
}
var AbiCommands  = &[]cli.Command {
	{
		Name:        "abi",
		Usage:       "solidity abi management",
		Flags: 		 abiFlags,
		Subcommands: []cli.Command{
			{
				Name:  "encode",
				Aliases:     []string{"e", "en"},
				Usage: "encode the solidity's method and params from terminal input",
				Action: func(c *cli.Context) error {
					fmt.Println("new task template: ", c.Args().First())
					return nil
				},
			},
			{
				Name:  "decode",
				Aliases:     []string{"d", "de"},
				Usage: "decode the solidity's method and params from json",
				Action: decode,
			},
		},
	},
}

