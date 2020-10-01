package main

import (
	"neopets/common"
	"neopets/controllers"
	"neopets/controllers/daily"
	"neopets/middlewares"
)

const NeoPriceDb = 1

func StartServer() {
	server := common.WebServer{Host: "", Port: 8090}
	server.Use(middlewares.Json)
	server.Get("/price", controllers.PriceWizard)
	server.Post("/daily/grundo", daily.MagicBlueGrundo)
	server.Post("/daily/shop-of-offer", daily.ShopOfOffer)
	server.Post("/daily/trudys-surprise", daily.TrudysSurprise)
	server.Start(nil)
}

func main() {
	rdb := &common.Redis{Addr: "localhost:6379", Password: "", Db: NeoPriceDb}
	rdb.Connect()
	StartServer()
}
