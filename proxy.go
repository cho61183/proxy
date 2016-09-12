package main

import (
	"fmt"
)

//	"fmt"
//	"os"
//	"os/signal"

const version = "2.0"

func newRedis() *ConnPool {
	fmt.Printf("初始化 Mysql 连接池，连接数：%d \n", 10)
	return &ConnPool{
		MaxActive: 10,
		//Dial: func() (*autorc.Conn, error) {
		Dial: func() (interface{}, error) {
			//conn := autorc.New("tcp", "", "localhost:3306", "root", "root", "test")
			//conn.Register("set names utf8")
			db, err := RedisDb()
			return db, err
		},
	}
}

func init() {
	//	fmt.Println("Test pool")
	//	redisPool := &ConnPool{
	//		MaxIdle:   80,
	//		MaxActive: 10,
	//		Dbclass: func() (Conn, error) {
	//			c, err := RedisDb(6379)
	//			return c, err
	//		},
	//	}
	//	fmt.Println(redisPool.Get())
	//poolNum, _ := strconv.Atoi(conf.GetValue("pool", "mysql"))

	var conn = newRedis()
	tmp := conn.Get()
	tmp.(*Conn).Get()
	fmt.Println("end")
}

func main() {

}
