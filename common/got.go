package common

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Got(method string, url string, body io.Reader, cookies []*http.Cookie) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	// 自定义 Header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36")
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		return nil, errors.New("status code not 200")
	}
	return resp, err
}
