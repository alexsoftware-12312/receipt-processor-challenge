package main

import (
	"log"
	"main/api/controllers"
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
	routeHandler()
	requestHandler()
}
