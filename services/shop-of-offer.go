package services

import (
	"log"
	"neopets/common"
	"net/http"
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

	// TODO: check if get successfully and show how much np to get
	log.Print("We got ? NP")
}
