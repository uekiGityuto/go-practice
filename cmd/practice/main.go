package main

import (
	"fmt"
	"github.com/uekiGityuto/go-practice/handler"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server started.")
	listen()
}

func listen() {
	http.HandleFunc("/user", handler.UserHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
