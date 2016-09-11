package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

// init func 初始化函数
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
	fmt.Println("Bingding port:", 7777)
	listener, err := net.Listen("tcp", "localhost:7777")
	CheckErr(err)
	fmt.Println("proxy server: start ok")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		CheckErr(err)
		go doServerStuff(conn)
	}
	fmt.Println("Bingding server:", "ok")
}

//开启监听函数. 这里发现的是服务端如何关闭线程,其实就是退出for循环体,让goroutine 自动结束.
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
		// 通过json转码,可以做到客户端传递过来的信息,包含了特俗字符,方便进行字符串过滤.
		//tmp, _ := json.Marshal(string(buf[:lenght]))
		//fmt.Println(string(tmp))

		content := strings.Trim(string(buf[:lenght]), "\r\n")
		if content == "quit" {
			fmt.Println("bye")
			break
		}
		//fmt.Println(content)
	}
}
