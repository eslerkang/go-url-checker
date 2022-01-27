package main

import (
	"fmt"
	"net/http"
)

type reqResult struct {
	url string
	status string
}

func main() {
	urls := []string{
		"https://www.naver.com",
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.airbnb.com",
	}

	results := map[string]string{}

	c := make(chan reqResult)

	for _, url:=range urls {
		go hitURL(url, c)
	}
	
	for i:=0; i<len(urls); i++ {
		result := <- c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}


func hitURL(url string, c chan<- reqResult) {
	fmt.Println("Checking: ", url)
	res, err := http.Get(url)
	status := "OK"
	if err != nil || res.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- reqResult{
		url: url,
		status: status,
	}
}