package mq

import (
	"github.com/nats-io/nats.go"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect("192.168.3.6:4222")
	if err != nil {
		panic(err)
	}
}
