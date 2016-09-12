package main

type Conn struct {
	Port int
}

func RedisDb(port int) (Conn, error) {

	c := Conn{Port: port}
	return c, nil
}

func (this *Conn) GetPort() int {
	return this.Port
}
