package services

import (
	"io/ioutil"
	"log"
	"neopets/common"
	"net/http"
	"net/url"
	"strings"
)

const TrudysSurpriseURL = "http://www.neopets.com/trudydaily/ajax/claimprize.php"
const TrudysSurpriseAction = "beginroll"

func GetTrudysSurprise() {
	data := url.Values{}
	data.Set("action", TrudysSurpriseAction)
	res, err := common.Got(http.MethodPost, TrudysSurpriseURL, strings.NewReader(data.Encode()), []*http.Cookie{})

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Print(bodyString)
}
