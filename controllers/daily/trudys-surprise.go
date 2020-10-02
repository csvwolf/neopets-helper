package daily

import (
	"encoding/json"
	"neopets/services"
	"neopets/types"
	"net/http"
)

func TrudysSurprise(w http.ResponseWriter, r *http.Request) {
	var body types.NeopetsSession

	resultChan := make(chan services.SurpriseItem)
	errorChan := make(chan error)

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go services.GetTrudysSurprise(body.Session, resultChan, errorChan)

	select {
	case item := <-resultChan:
		value, err := json.Marshal(item)

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
