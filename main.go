package main

import (
	"neopets/common"
	"neopets/controllers"
)

const NeoPriceDb = 1

func StartServer() {
	server := common.WebServer{Host: "", Port: 8090}
	server.Get("/price", controllers.PriceWizard)
	server.Start(nil)
}

func main() {
	rdb := &common.Redis{Addr: "localhost:6379", Password: "", Db: NeoPriceDb}
	rdb.Connect()
	StartServer()
}
