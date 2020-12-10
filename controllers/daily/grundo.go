package daily

import (
	"encoding/json"
	"neopets/services"
	"neopets/types"
	"net/http"
)

/**
MagincBlueGrundo API
*/
func MagicBlueGrundo(w http.ResponseWriter, r *http.Request) {
	var body types.NeopetsSession

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	grundo, err := services.GetMagicBlueGrundoGift(body.Session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	value, err := json.Marshal(grundo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if grundo.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	_, err = w.Write(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
