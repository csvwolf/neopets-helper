package controllers

import (
	"encoding/json"
	"neopets/services"
	"net/http"
)

/**
PriceWizard to compare prize
*/
func PriceWizard(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := services.QuickGetPrice(item)
	value, err := json.Marshal(services.PriceWizard{
		Price: price,
		Name:  item,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
