package main

import (
	"fmt"
)

//	"fmt"
//	"os"
//	"os/signal"

const version = "2.0"

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
	fmt.Printf("初始化 Mysql 连接池，连接数：%d \n", 10)
	pool := &ConnPool{
		MaxActive: 10,
		//Dial: func() (*autorc.Conn, error) {
		Dial: func() (interface{}, error) {
			//conn := autorc.New("tcp", "", "localhost:3306", "root", "root", "test")
			//conn.Register("set names utf8")
			db, err := RedisDb(6379)
			return db, err
		},
	}
	fmt.Println(pool.Get())
}

func main() {

}
