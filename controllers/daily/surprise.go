package daily

import (
	"encoding/json"
	"neopets/services"
	"neopets/types"
	"net/http"
)

func TrudysSurprise(w http.ResponseWriter, r *http.Request) {
	var body types.NeopetsSession

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, err := services.GetTrudysSurprise(body.Session)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value, err := json.Marshal(item)
	_, err = w.Write(value)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
