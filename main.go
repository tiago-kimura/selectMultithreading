package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type messsage struct {
	id  int
	msg string
}

func main() {
	chanRMQ := make(chan messsage)
	chanKFK := make(chan messsage)
	var id int64
	go func() {
		for {
			chanRMQ <- messsage{
				id:  int(atomic.AddInt64(&id, 1)),
				msg: "message from RabbitMQ",
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			chanKFK <- messsage{
				id:  int(atomic.AddInt64(&id, 1)),
				msg: "message from Kafka",
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case msgRMQ := <-chanRMQ:
			fmt.Println("Received: ", msgRMQ)
		case msgKFK := <-chanKFK:
			fmt.Println("Received: ", msgKFK)
		case <-time.After(time.Second * 3):
			fmt.Print("Timeout")
		}
	}
}
