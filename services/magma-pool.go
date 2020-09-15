package services

import (
	"log"
	"neopets/common"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const MagmaPool = "http://www.neopets.com/magma/pool.phtml"
const SorryWords = "I'm sorry, only those well-versed in the ways of Moltara are permitted to enter the Pool.  Learn more and try again later."

/**
Get Magma Pool Open Time
One minute a day, every user is unique.
*/
func FindOpenTime() {
	ticker := time.NewTicker(time.Minute)
	quit := make(chan time.Time)
	log.Print("Start to Find Open Time")
	defer ticker.Stop()
	for {
		select {
		case t := <-quit:
			log.Print("Found your Magma Time:")
			log.Print(t)
			return
		case <-ticker.C:
			log.Print("Go To Magma...")
			ok, t := GoToMagma()
			if ok {
				quit <- t
			}
		}
	}
}

/**
One time Request to Go To Magma Pool
*/
func GoToMagma() (bool, time.Time) {
	res, err := common.Got("GET", MagmaPool, nil, []*http.Cookie{})
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	text := strings.TrimSpace(doc.Find("#poolOuter").First().Text())

	if text != SorryWords {
		log.Println(text)
		return true, time.Now()
	} else {
		return false, time.Time{}
	}
}
