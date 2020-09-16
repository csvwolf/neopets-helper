package services

import (
	"log"
	"neopets/common"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const GrundoGift = "http://www.neopets.com/faerieland/tdmbgpop.phtml"

/**
Once a day to get magic blue grundo's gift
*/
func GetMagicBlueGrundoGift() {
	res, err := common.Got("POST", GrundoGift, strings.NewReader("talkto=1"), []*http.Cookie{})
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	text := strings.TrimSpace(doc.Find("#content > table > tbody > tr > td.content > div[align=center] > b").First().Text())

	log.Print("Got " + text)
}
