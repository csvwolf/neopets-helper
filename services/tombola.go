package services

import (
	"neopets/common"
	"neopets/utils"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const TombolaURL = "http://www.neopets.com/island/tombola2.phtml"

/**
GetTombola to roll price
*/
func GetTombola(session string) (string, error) {
	res, err := common.Got(http.MethodPost, TombolaURL, nil, []*http.Cookie{utils.NeopetsSession(session)})

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}
	// TODO: Not get the gift now, output whole text to future analyze
	return doc.Text(), nil
}
