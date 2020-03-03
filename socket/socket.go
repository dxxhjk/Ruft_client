package socket

import (
	"fmt"
	"net"
)

type Udpsocket struct {
	Addr string
	Buf  [512]byte
}

func New(addr string) *Udpsocket {
	u := new(Udpsocket)
	u.Addr = addr
	return u
}

func (u *Udpsocket) Send(msg string) {
	conn, err := net.Dial("udp", u.Addr)
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("Hello! I'm client in UDP! " + msg))
}
