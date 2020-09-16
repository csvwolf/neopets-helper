package services

import (
	"log"
	"neopets/common"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const ShopOfOffer = "http://www.neopets.com/shop_of_offers.phtml?slorg_payout=yes"

/**
Rich Slorg: Get 100 NP or 50 NP a day
*/
func GetShopOfOffer() {
	res, err := common.Got("GET", ShopOfOffer, nil, []*http.Cookie{})
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	text := strings.TrimSpace(doc.Find("#content > table > tbody > tr > td.content > table > tbody").First().Text())
	re := regexp.MustCompile(`\d+ Neopoints`)

	log.Print("Got " + string(re.Find([]byte(text))))
}
