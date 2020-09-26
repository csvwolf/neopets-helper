package services

import (
	"log"
	"neopets/common"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const GrundoGift = "http://www.neopets.com/faerieland/tdmbgpop.phtml"

type BlueGrundoGift struct {
	Gift string `json:"gift"`
}

/**
Once a day to get magic blue grundo's gift
*/
func GetMagicBlueGrundoGift(session string, resultChan chan string, errorChan chan error) {
	res, err := common.Got("POST", GrundoGift, strings.NewReader("talkto=1"), []*http.Cookie{
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
	text := strings.TrimSpace(doc.Find("#content > table > tbody > tr > td.content > div[align=center] > b").First().Text())

	log.Print("Got " + text)

	resultChan <- text
}
