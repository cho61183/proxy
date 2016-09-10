package main

import (
	"fmt"
	"strings"
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

/*
 对服务的分发动作处理。
add
*/
func (s ServerInfo) Controller() (string, bool) {
	fmt.Println("This is Model.Servier msg :")
	switch strings.ToLower(s.Action) {
	case "add":
		return add(s)
	case "del":
		return del(s)
	}
	return "ok", true

}

/*
添加动作
*/
func add(s ServerInfo) (string, bool) {
	fmt.Println("add")
	return "This add:", true
}

/*
删除动作
*/
func del(s ServerInfo) (string, bool) {
	fmt.Println("del")
	return "This del:", true
}
