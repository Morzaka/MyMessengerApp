package main

import (
	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
	"log"
)

type User struct {
	Id string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:32769",
		Database: "rtsupport",
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	//response, err := r.Table("user").Insert(user).RunWrite(session)
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}

	//user := User{
	//	Name: "Monika Macedonia",
	//}
	//
	//response, _ := r.Table("user").
	//	Get("067c991d-1394-4709-ac9a-e66a7b1db832").
	//	Update(user).
	//	RunWrite(session)
	//log.Println(response)

	cursor, _ := r.Table("user").
		Changes(r.ChangesOpts{IncludeInitial: true}).
		Run(session)
	var changeResponse r.ChangeResponse
	for cursor.Next(&changeResponse) {
		log.Println(changeResponse)
	}
}