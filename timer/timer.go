package timer

import (
	"time"
)

//工厂模式，只能通过New函数创建
//出于client结构体声明需要，这里目前大写，之后如何处理待定
type MsgTimer struct {
	Ch    chan string
	Delay time.Duration
}

func New(ch chan string, delay time.Duration) *MsgTimer {
	p := new(MsgTimer)
	p.Ch = ch
	p.Delay = delay
	return p
}

func (k *MsgTimer) Start() {
	t := time.NewTimer(k.Delay)
	select {
	case <-t.C:
		k.Ch <- "timer fired"
	}
}

//func (k *MsgTimer) Init(delay time.Duration) {
//	go k.Start(delay)
//
//	for {
//		select {
//		case msg = <-k.Ch:
//			fmt.Println(msg)
//			go k.start(delay)
//		}
//	}
//}
