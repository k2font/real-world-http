package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// 一番かんたんなHTTPクライアント

func main() {
	// jsonを返してくれるPostman開発者が作ったWebサーバにリクエスト
	// resp, err := http.Get("https://httpbin.org/get")
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()                // この関数から抜けた後に処理を実行する
	body, err := ioutil.ReadAll(resp.Body) // ボディの内容をバイト列として埋め込む
	if err != nil {
		panic(err)
	}
	// レスポンスボディすべて
	log.Println(string(body))
	// ステータス
	log.Println(string(resp.Status))
	// レスポンスヘッダ
	// 文字列配列のmap型
	log.Println(resp.Header)
}
