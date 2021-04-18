package main

import (
	"awesomeProject1/internal/handler"
	"awesomeProject1/internal/service"

	"net/http"
)

func main() {
	nameService := service.InitNameService("redis:6379")

	handler.InitHandler(nameService)

	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		panic(err)
	}
}
