package abi

import (
	"github.com/urfave/cli"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"encoding/hex"
)

func getEncodeDataJson(c *cli.Context) (error, string) {
	if c.Args().First() == "" {
		return errors.New("abi: params decode data required"), ""
	}
	return nil, c.Args().First()
}
func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}
func getDecodeData(c *cli.Context) (error, []decodedCallData) {
	var (
		err  error
		dds []decodedCallData
		encodeDataBytes []byte
	)
	err, abiJsonString := getAbiJson(c)
	if err != nil {
		return err, nil
	}
	err, encodeJsonString := getEncodeDataJson(c)
	if err != nil {
		return err, nil
	}
	if has0xPrefix(encodeJsonString) {
		encodeDataBytes, err = hexutil.Decode(encodeJsonString)
	}else if encodeDataBytes, err = hex.DecodeString(encodeJsonString); err != nil {
		// it is likely a json string
		encodeDataBytes = []byte(encodeJsonString)
		err = nil
	}
	if err != nil {
		return err, nil
	}
	if dds, err = parseCallData(encodeDataBytes, abiJsonString); err == nil {
		return nil, dds
	}else if dds, err = parseEventData(encodeDataBytes, abiJsonString); err == nil {
		return nil, dds
	}
	return errors.New("abi: faild to decode"), nil
}

func decode(c *cli.Context) error {
	err, data := getDecodeData(c)
	if err != nil {
		return err
	}
	if len(data) > 0 {
		fmt.Println("abi decode successfully!")
		if len(data) == 1 {
			fmt.Println(data[0])
		}else {
			fmt.Println(data)
		}
	}else {
		return errors.New("abi: decode get empty list, fail")
	}

	return nil
}