package services

import (
	"log"
	"neopets/common"
	"net/http"
	"strings"
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

	// TODO: check if get successfully and show gift
	log.Print("We got ...")
}
