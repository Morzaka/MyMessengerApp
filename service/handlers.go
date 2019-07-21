package main

import (
	"github.com/mitchellh/mapstructure"
	"log"
)

func addChannel(client *Client, data interface{}) {
	var channel Channel
	var message Message

	err := mapstructure.Decode(data, &channel)
	if err != nil {
		log.Println(err.Error())
	}
	// TODO: insert into RethinkDB

	channel.Id = "ABC123"
	message.Name = "channel add"
	message.Data = channel
	client.send <- message
}
