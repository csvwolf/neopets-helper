package services

import (
	"fmt"
	"log"
	"neopets/common"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const DatabaseSearchUrl = "https://items.jellyneo.net/search/?name=%s&name_type=3"

type PriceWizard struct {
	Name  string `json:"name"`
	Price string `json:"price"`
}

/**
Quick Get the Market Item Price
It will be helpful in jobs, quest and rob cargo.
The result is from: https://items.jellyneo.net/
*/
func QuickGetPrice(name string) string {
	val, err := common.RDB.Read(name)
	if err == nil {
		return val
	}

	url := fmt.Sprintf(DatabaseSearchUrl, strings.Replace(name, " ", "+", -1))
	res, err := common.Got("Get", url, nil, []*http.Cookie{})

	if err != nil {
		log.Print(err)
		log.Print("Try later")
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	price := doc.Find(".price-history-link").First().Text()
	log.Print(price)

	// cache a week
	common.RDB.Write(name, price, time.Hour*24*7)

	return price
}
