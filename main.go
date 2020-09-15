package main

import (
	"neopets/controllers"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/price", controllers.PriceWizard)
	http.ListenAndServe(":8090", nil)
}

func main() {
	StartServer()
}
