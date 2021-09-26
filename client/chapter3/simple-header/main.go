package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTPヘッダの取得
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", resp.Status)
	log.Println("Headers: ", resp.Header)
}
