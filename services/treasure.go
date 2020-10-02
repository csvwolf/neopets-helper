package services

import (
	"log"
	"neopets/common"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const TarlaTresure = "http://www.neopets.com/freebies/tarlastoolbar.phtml"

func CheckTarlaIsIn() {
	ticker := time.NewTicker(20 * time.Second)
	log.Print("Start to Check Tarla is in")
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			log.Print("Are Tarla at home?")
			ok, t := GetTarlaTresure()
			if ok {
				log.Print("Yes, Tarla is at home", t)
			} else {
				log.Print("Sorry for next time")
			}
		}
	}
}

func GetTarlaTresure() (bool, time.Time) {
	res, err := common.Got("GET", TarlaTresure, nil, []*http.Cookie{})
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	text := strings.TrimSpace(doc.Find(".content div").First().Text())

	if text != SorryWords {
		log.Println(text)
		return true, time.Now()
	} else {
		return false, time.Time{}
	}
}
