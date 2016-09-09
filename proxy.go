package main

import (
	"fmt"
	"os"

	"github.com/cho61183/proxy/model"
	"github.com/cho61183/proxy/tools"
)

func main() {
	if msg, ok := tools.CheckArgs(os.Args); !ok {
		fmt.Println(msg)
		return
	}

	Info := model.ServerInfo{}

	Info.Action = os.Args[1]
	Info.Model = os.Args[2]
	Info.Ip = os.Args[3]
	Info.Port = tools.ToInt(os.Args[4])
	Info.Controller()
}
