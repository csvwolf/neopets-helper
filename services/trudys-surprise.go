package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"neopets/common"
	"net/http"
	"net/url"
	"strings"
)

const TrudysSurpriseURL = "http://www.neopets.com/trudydaily/ajax/claimprize.php"
const TrudysSurpriseAction = "beginroll"

type SurprisePrize struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Value string `json:"value"`
}

type SurpriseItem struct {
	AdjustedNp int             `json:"adjustedNp"`
	BadLuck    string          `json:"badLuck"`
	GameState  int             `json:"gameState"`
	Prizes     []SurprisePrize `json:"prizes"`
	Error      string          `json:"error"`
}

func GetTrudysSurprise(session string, resultChan chan SurpriseItem, errorChan chan error) {
	data := url.Values{}
	data.Set("action", TrudysSurpriseAction)
	res, err := common.Got(http.MethodPost, TrudysSurpriseURL, strings.NewReader(data.Encode()), []*http.Cookie{
		{
			Name:   "neologin",
			Value:  session,
			Domain: ".neopets.com",
		},
	})

	if err != nil {
		errorChan <- err
		return
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		errorChan <- err
		return
	}

	var item SurpriseItem
	if err := json.Unmarshal(bodyBytes, &item); err != nil {
		errorChan <- err
		return
	}

	if item.Error != "" {
		errorChan <- errors.New(item.Error)
		return
	}

	resultChan <- item
}
