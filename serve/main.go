package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/websocket"
    "github.com/mitchellh/mapstructure"
)

type Message struct {
    Name string `json:"name"`
    Data interface{} `json:"data"`
}

type Channel struct {
    Id string
    Name string
}

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {return true},
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":4000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hello Server")
    socket, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    for {
        //msgType, msg, err := socket.ReadMessage()
        //if err != nil {
        //    log.Println(err)
        //    return
        //}
        var inMessage Message
        if err := socket.ReadJSON(&inMessage); err != nil {
            log.Println(err)
            break
        }
        fmt.Printf("%#v\n", inMessage)
        switch inMessage.Name {
        case "channel add":
            err := addChannel(inMessage.Data)
            if err != nil {
                outMessage = Message{"error", err}
                if err := socket.WriteJSON(outMessage); err != nil {
                    log.Println(err)
                    break
                }
            }
        case "channel subscribe":
            go subscribeChannel(socket)
        }
        //fmt.Println(string(msg))
        //if err = socket.WriteMessage(msgType, msg); err != nil {
        //    log.Println(err)
        //    return
        //}
    }
}

func addChannel(data interface{}) (Channel, error) {
    var channel Channel

    err := mapstructure.Decode(data, &channel)
    if err != nil {
        return Channel{}, err
    }
    channel.Id = "1"
    log.Println("added channel")
    return channel, nil
}

func subscribeChannel(socket *websocket.Conn) {
    //TODO: rethinkDB Query / changelog
    for {
        time.Sleep(time.Second * 1)
        message := Message{"channel add",
            Channel{"1", "Software Support"}}
        socket.WriteJSON(message)
        fmt.Println("sent new channel")
    }
}