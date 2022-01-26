package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errReqFail = errors.New("Request Failed")

func main() {
	urls := []string{
		"https://www.naver.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.airbnb.com",
	}

	results := map[string]string{}

	for _, url:=range urls {
		result := "OK"
		err := hitURL(url)
		if err!=nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}


func hitURL(url string) error {
	fmt.Println("Checking ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errReqFail
	}
	return nil
}