package main

import (
	"log"
	"main/api/controllers"
	"main/internal/data"
	"net/http"
)

func requestHandler() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func routeHandler() {
	controllers.InitReceiptsController()
	controllers.InitPingController()
}

func main() {
	data.InitDb()
	routeHandler()
	requestHandler()
}
