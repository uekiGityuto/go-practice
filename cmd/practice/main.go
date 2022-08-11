package main

import (
	"fmt"
	"github.com/uekiGityuto/go-practice/handler"
)

func main() {
	fmt.Println("Server started.")
	handler.Listen()
}
