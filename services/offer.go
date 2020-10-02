package services

import (
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
func GetShopOfOffer(session string) (string, error) {
	res, err := common.Got("GET", ShopOfOffer, nil, []*http.Cookie{
		{
			Name:   "neologin",
			Value:  session,
			Domain: ".neopets.com",
		},
	})
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}
	text := strings.TrimSpace(doc.Find("#content > table > tbody > tr > td.content > table > tbody").First().Text())
	re := regexp.MustCompile(`\d+ Neopoints`)
	np := string(re.Find([]byte(text)))

	return np, nil
}
