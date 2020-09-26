package daily

import (
	"encoding/json"
	"neopets/services"
	"neopets/types"
	"net/http"
)

func MagicBlueGrundo(w http.ResponseWriter, r *http.Request) {
	var body types.NeopetsSession

	gundoChan := make(chan string)
	errorChan := make(chan error)
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	go services.GetMagicBlueGrundoGift(body.Session, gundoChan, errorChan)

	select {
	case gift := <-gundoChan:
		value, err := json.Marshal(services.BlueGrundoGift{
			Gift: gift,
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
	case err := <-errorChan:
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
