package services

import (
	"log"
	"neopets/common"
	"neopets/utils"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const GrundoGift = "http://www.neopets.com/faerieland/tdmbgpop.phtml"

type BlueGrundoGift struct {
	Gift  string `json:"gift"`
	Error string `json:"error"`
}

/**
Once a day to get magic blue grundo's gift
*/
func GetMagicBlueGrundoGift(session string) (grundo *BlueGrundoGift, err error) {
	res, err := common.Got("POST", GrundoGift, strings.NewReader("talkto=1"), []*http.Cookie{
		utils.NeopetsSession(session),
	})
	if err != nil {
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return
	}
	if !utils.CheckSessionInPage(doc.Find("body").Text()) {
		grundo = &BlueGrundoGift{Error: "Session expired"}
		return
	}

	text := strings.TrimSpace(doc.Find("#content > table > tbody > tr > td.content > div[align=center] > b").First().Text())
	log.Print("Got " + text)
	grundo = &BlueGrundoGift{Gift: text}
	return
}
