package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strings"
)

func mainPageGetRequest(client *resty.Client, url requestUrl) {
	// Default case
	if url.shortUrl == "" {
		url.shortUrl = "68747470733a2f2f70726163"
	}

	r, err := client.R().
		Get(url.shortUrl)

	if err != nil {
		// Redirect error processing
		if strings.Contains(err.Error(), "auto redirect is disabled") {
			fmt.Println("Status:" + r.Status())
			fmt.Println("Location:", r.Header().Get("Location"))
			return
		}
		panic(err)
	}

}

func mainPagePostRequest(client *resty.Client, url requestUrl) {
	//Default case
	if url.longUrl == "" {
		url.longUrl = "https://practicum.yandex.ru/"
	}

	body := []byte(url.longUrl)
	r, err := client.R().
		SetBody(body).
		SetHeader("Content-Type", "text/plain").
		Post("")
	if err != nil {
		panic(err)
	}
	fmt.Println("Status:" + r.Status())
	fmt.Println("Content-Type: " + r.Header().Get("Content-Type"))
	fmt.Println("Body: " + string(r.Body()))

}

func shortenApiPostRequest(client *resty.Client, url requestUrl) {
	//Default case
	if url.longUrl == "" {
		url.longUrl = "https://practicum.yandex.ru/"
	}

	body, err := json.Marshal(url.longUrl)
	if err != nil {
		panic(err)
	}
	r, err := client.R().
		SetBody(body).
		SetHeader("Content-Type", "application/json").
		Post("api/shorten")
	if err != nil {
		panic(err)
	}
	fmt.Println("Status:" + r.Status())
	fmt.Println("Content-Type: " + r.Header().Get("Content-Type"))
	fmt.Println("Body: " + string(r.Body()))
}
func pingDBRequest(client *resty.Client) {
	r, err := client.R().
		Get("ping")
	if err != nil {
		panic(err)
	}
	fmt.Println("Status:" + r.Status())
}
