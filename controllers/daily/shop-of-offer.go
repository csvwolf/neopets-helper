package daily

import (
	"encoding/json"
	"neopets/services"
	"neopets/types"
	"net/http"
)

func ShopOfOffer(w http.ResponseWriter, r *http.Request) {
	var body types.NeopetsSession

	offerChan := make(chan string)
	errorChan := make(chan error)
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go services.GetShopOfOffer(body.Session, offerChan, errorChan)

	select {
	case np := <-offerChan:
		value, err := json.Marshal(services.OfferNp{Np: np})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = w.Write(value)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case err := <-errorChan:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
