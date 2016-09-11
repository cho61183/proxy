package main

import (
	"fmt"
	"time"
)

var nowFunc = time.Now()

type ConnPool struct {
	//db 连接类
	Dial func() (interface{}, error)
	//最大连接数.
	MaxIdle int
	//最大请求数
	MaxActive int
	//定义连接类型
	idle chan interface{}
}

type idleConn struct {
	c interface{}
	t time.Time
}

func (this *ConnPool) InitPool() error {
	this.idle = make(chan interface{}, this.MaxIdle)
	for i := 0; i < this.MaxIdle; i++ {
		db, err := this.Dial
		if err != nil {
			return err
		}
		this.idle <- idleConn{t: nowFunc(), c: db}
	}
	return nil
}

func (this *ConnPool) Release(conn interface{}) {
	this.idle <- idleConn{t: nowFunc(), c: conn}
}

func (this *ConnPool) Get() error {
	if this.idle == nil {
		this.InitPool()
	}
	ic := this.idle
	conn := ic.(idleConn).c
	defer this.Release(conn)
	return conn
}

func init() {
	fmt.Println(nowFunc)
}
