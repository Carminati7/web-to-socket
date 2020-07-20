package main

import (
	"fmt"
	"log"
	"time"
	"print-app/socket"
	"encoding/json"
	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func main() {
	connect()
	for {
	}
}

type Message struct {
	Address string
	Data string
}

type Data struct{
	Type string
	Data []uint8
}

func connect() {
	opts := &socketio_client.Options{
		Transport: "websocket",
		Query:     make(map[string]string),
	}
	uri := "http://localhost:3001"
	client, err := socketio_client.NewClient(uri, opts)
	if err != nil {
		log.Printf("NewClient error:%v\n", err)
	 	time.Sleep(2 * time.Second)
		connect()
		return
	}
	client.On("connection", func() {
		log.Printf("on connect\n")
	})
	client.On("hello", func(msg Message) {
		log.Printf("on message:%v\n", msg.Data)
		print( msg.Address , msg.Data)
	})
	client.On("disconnection", func() {
		connect()
	})
}

func print( address string , data string) {
	var byte Data
	var err = json.Unmarshal([]uint8(data) , &byte )
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("data:", byte.Data)
	socket.Write(byte.Data, address)
}
