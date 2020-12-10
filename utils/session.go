package utils

import (
	"net/http"
	"strings"
)

// NeopetsSession Set Default Cookie
func NeopetsSession(session string) *http.Cookie {
	return &http.Cookie{
		Name:   "neologin",
		Value:  session,
		Domain: ".neopets.com",
	}
}

// CheckSessionInPage check if user is login
func CheckSessionInPage(page string) bool {
	return !strings.Contains(page, "Forgot your Username or Password?")
}
