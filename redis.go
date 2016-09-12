package main

import (
	"fmt"
)

type Conn struct {
	Port int
}

func RedisDb() (*Conn, error) {
	c := &Conn{Port: 6666}
	return c, nil
}

func (this *Conn) GetPort() int {
	return this.Port
}

func (this *Conn) Get() {
	fmt.Println("This is get:", this.Port)
}
