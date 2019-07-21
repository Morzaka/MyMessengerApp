package main

import (
    "log"
    "net/http"
)

type Channel struct {
    Id string `json:"id"`
    Name string `json:"name"`
}

func main() {
    router := newRouter()

    router.Handle("channel add", addChannel)

    http.Handle("/", router)
    log.Fatal(http.ListenAndServe(":4000", nil))
}