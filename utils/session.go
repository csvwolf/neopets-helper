package utils

import "net/http"

func NeopetsSession(session string) *http.Cookie {
	return &http.Cookie{
		Name:   "neologin",
		Value:  session,
		Domain: ".neopets.com",
	}
}
