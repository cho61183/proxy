package main

import (
	"fmt"
	"os"

	"github.com/cho61183/proxy/tools"
)

/*
	定义接收的参数.
	action 动作
	model 模块
	ip ip地址
	port 端口
*/
type ServerInfo struct {
	Action string
	Model  string
	Ip     string
	Port   int
}

func main() {

	if msg, ok := tools.CheckArgs(os.Args); !ok {
		fmt.Println(msg)
		return
	}

	Info := ServerInfo{}

	Info.Action = os.Args[1]
	Info.Model = os.Args[2]
	Info.Ip = os.Args[3]
	Info.Port = tools.ToInt(os.Args[4])

	fmt.Println(Info)
	//	switch ServerInfo.Action {

	//	}
}
