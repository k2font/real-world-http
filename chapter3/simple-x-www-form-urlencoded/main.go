package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	// x-www-form-urlencoded を使ったPOSTリクエスト
	values := url.Values{
		"test": {"values"},
	}

	resp, err := http.PostForm("http;//localhost:18888", values)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
