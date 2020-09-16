package main

import (
	"neopets/common"
	"neopets/controllers"
	"net/http"
)

const NeoPriceDb = 1

func StartServer() {
	http.HandleFunc("/price", controllers.PriceWizard)
	http.ListenAndServe(":8090", nil)
}

func main() {
	rdb := &common.Redis{Addr: "localhost:6379", Password: "", Db: NeoPriceDb}
	rdb.Connect()
	StartServer()
}
