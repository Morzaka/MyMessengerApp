package main

import (
	"log"
	"net/http"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

type Channel struct {
	Id   string `json:"id" gorethink:"id,omitempty"`
	Name string `json:"name" gorethink:"name"`
}

type User struct {
	Id   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:32769",
		Database: "rtsupport",
	})

	if err != nil {
		log.Panic(err.Error())
		return
	}
	router := NewRouter(session)

	router.Handle("channel add", addChannel)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
