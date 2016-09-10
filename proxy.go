package main

import (
	"fmt"
	"os"
)

const version = "1.0"

func main() {
	if msg, ok := CheckArgs(os.Args); !ok {
		fmt.Println(msg)
		return
	}

	Info := ServerInfo{}

	Info.Action = os.Args[1]
	Info.Model = os.Args[2]
	Info.Ip = os.Args[3]
	Info.Port = ToInt(os.Args[4])
	Info.Controller()
}
