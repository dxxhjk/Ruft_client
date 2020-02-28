package main

import (
	"./socket"
	"./timer"
	"fmt"
	"time"
)

type client struct {
	msgtimer  timer.MsgTimer
	udpsocket socket.Udpsocket
}

func (c *client) ini() {
	go c.startMsgHandler()
	go c.msgtimer.Start()
}

func newClient() *client {
	c := new(client)
	msgch := make(chan string)
	c.msgtimer = *timer.New(msgch, 1*time.Second)
	c.udpsocket = *socket.New("127.0.0.1:8888")
	return c
}

func (c *client) startMsgHandler() {
	var msg string
	for {
		select {
		case msg = <-c.msgtimer.Ch:
			fmt.Println(msg)
			go c.msgtimer.Start()
			go c.udpsocket.Send()
		}
	}
}

func main() {
	c := newClient()
	c.ini()
	fmt.Println("client")
	time.Sleep(10 * time.Second)
	fmt.Println("client")
}
