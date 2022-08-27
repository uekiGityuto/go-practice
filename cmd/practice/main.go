package main

import (
	"github.com/uekiGityuto/go-practice/infra/dao"
	"github.com/uekiGityuto/go-practice/ui/handler"
	"github.com/uekiGityuto/go-practice/ui/middleware"
	"github.com/uekiGityuto/go-practice/ui/validator"
	"github.com/uekiGityuto/go-practice/usecase"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := validator.RegisterCustomValidator(); err != nil {
		log.Printf("%+v\n", err)
		os.Exit(1)
	}
	db := dao.NewDB()
	userHandler := handler.NewUser(*usecase.NewUser(dao.NewUser(), db))
	http.Handle("/user", middleware.Logger(userHandler))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("サーバの起動に失敗しました。%+v\n", err)
		os.Exit(1)
	}
}
