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

type OfferNp struct {
	Np string
}

/**
Rich Slorg: Get 100 NP or 50 NP a day
*/
func GetShopOfOffer(session string, resultChan chan string, errorChan chan error) {
	res, err := common.Got("GET", ShopOfOffer, nil, []*http.Cookie{
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

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		errorChan <- err
		return
	}
	text := strings.TrimSpace(doc.Find("#content > table > tbody > tr > td.content > table > tbody").First().Text())
	re := regexp.MustCompile(`\d+ Neopoints`)
	np := string(re.Find([]byte(text)))
	resultChan <- np
	log.Print("Got " + np)
}
