package queue

import (
	"fmt"
)

type InMemoryQueue struct {
	Queue chan string
}

func NewInMemoryQueue() *InMemoryQueue {
	return &InMemoryQueue{
		Queue: make(chan string),
	}
}

func (i *InMemoryQueue) Consume() {
	select {
	case msg := <-i.Queue:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
}

func (i *InMemoryQueue) Push(content string) {
	msg := "hi"

	select {
	case i.Queue <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
}
