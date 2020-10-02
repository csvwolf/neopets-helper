package services

import (
	"encoding/json"
	"io/ioutil"
	"neopets/common"
	"neopets/utils"
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

/**
GetTrudysSurprise EveryDay
*/
func GetTrudysSurprise(session string) (*SurpriseItem, error) {
	data := url.Values{}
	data.Set("action", TrudysSurpriseAction)
	res, err := common.Got(http.MethodPost, TrudysSurpriseURL, strings.NewReader(data.Encode()), []*http.Cookie{
		utils.NeopetsSession(session),
	})

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var item SurpriseItem
	if err := json.Unmarshal(bodyBytes, &item); err != nil {
		return nil, err
	}

	if item.Error != "" {
		return nil, err
	}

	return &item, nil
}
