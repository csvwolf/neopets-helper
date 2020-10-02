package daily

import (
	"encoding/json"
	"neopets/services"
	"neopets/types"
	"net/http"
)

func ShopOfOffer(w http.ResponseWriter, r *http.Request) {
	var body types.NeopetsSession

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	np, err := services.GetShopOfOffer(body.Session)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	value, err := json.Marshal(services.OfferNp{Np: np})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = w.Write(value)
}
