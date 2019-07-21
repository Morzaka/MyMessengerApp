package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type FindHandler func(string) (Handler, bool)

type Message struct {
	Name string `json:"name"`
	Data interface{} `json:"data"`
}

type Client struct {
	send chan Message
	socket *websocket.Conn
	findHandler FindHandler
}

func (client *Client) Read() {
	var message Message
	for {
		if err := client.socket.ReadJSON(&message); err != nil {
			break
		}
		if handler, found := client.findHandler(message.Name); found {
			handler(client, message.Data)
		}
		err := client.socket.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func (client *Client) Write() {
	for msg := range client.send{
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	err := client.socket.Close()
	if err != nil {
		log.Println(err.Error())
	}
}


func NewClient(socket *websocket.Conn, findHandler FindHandler) *Client{
	return &Client{
		send: make(chan Message),
		socket: socket,
		findHandler: findHandler,
	}
}