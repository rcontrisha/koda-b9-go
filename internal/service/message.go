package service

import (
	"fmt"
	"time"
)

type Message struct {
	Name string
	Msg  string
}

func Run() {
	channel := make(chan Message)

	go Board(channel)

	messages := []Message{
		{Name: "Gipen", Msg: "Tes"},
		{Name: "Niko", Msg: "Hidup Jowoki"},
		{Name: "Alfan", Msg: "P"},
		{Name: "Pajar", Msg: "Wahuy"},
		{Name: "Karlos", Msg: "Test"},
	}

	for _, i := range messages {
		SendMsg(channel, i)
	}
}

func SendMsg(schan chan<- Message, msg Message) {
	schan <- msg
}

func Board(bchan <-chan Message) {
	for msg := range bchan {
		fmt.Printf("Message from %s: %s\n", msg.Name, msg.Msg)
		time.Sleep(100 * time.Millisecond)
	}
}
