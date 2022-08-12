package main

import (
	"fmt"
	"github.com/uekiGityuto/go-practice/handler"
	"github.com/uekiGityuto/go-practice/infra/dao"
	"github.com/uekiGityuto/go-practice/usecase"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server started.")
	listen()
}

func listen() {
	db := dao.NewDB()
	userHandler := handler.NewUser(*usecase.NewUser(dao.NewUser(db)))
	http.HandleFunc("/user", userHandler.HandleUser)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
