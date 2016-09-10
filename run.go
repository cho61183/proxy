package main

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

func init() {
	fmt.Println("version : ", version)
	goversion, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Go	  :", string(goversion))

	loadServer()
	checkServer()
	reportServer()
	bingingServer()
}

//加载所有服务
func loadServer() {
	fmt.Println("Load server :", "ok")
}

//检测服务状态
func checkServer() {
	fmt.Println("Check server:", "ok")
}

//报告服务状态
func reportServer() {
	fmt.Println("Reprot server:", "error:0", "ok:5")
}

//绑定端口开始接受查询等相关操作
func bingingServer() {
	fmt.Println("Bingding port:", 1234)
	fmt.Println("Bingding server:", "ok")
}

func StartServer() {
	listener, err := net.Listen("tcp", "localhost:7777")
	CheckErr(err)
	fmt.Println("proxy server: start ok")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		CheckErr(err)
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	fmt.Println("connection is connected from ...", conn.RemoteAddr().String())

	defer conn.Close()
	for {

		buf := make([]byte, 512)
		lenght, err := conn.Read(buf)
		err_flg := CheckErr(err)
		if err_flg == 0 {
			break
		}

		content := strings.Trim(string(buf[:lenght-1]), "\r")
		if content == "1" {
			fmt.Println("123123")
		}
		fmt.Println(content)
	}
}
