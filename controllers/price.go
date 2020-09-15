package controllers

import (
	"encoding/json"
	"neopets/services"
	"net/http"
)

func PriceWizard(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := services.QuickGetPrice(item)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	value, err := json.Marshal(services.PriceWizard{
		Price: price,
		Name:  item,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(value)
}
