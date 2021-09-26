# real-world-http

[![check-server](https://github.com/k2font/real-world-http/actions/workflows/check-server.yml/badge.svg?branch=main)](https://github.com/k2font/real-world-http/actions/workflows/check-server.yml)

## これは何?
- Real World HTTP 第2版を読んで書いたサンプルコード集
- サンプルWebサーバを立ち上げてからクライアント側を動かします

## 動かし方
- サーバ立ち上げ
```bash
$ go run simple-server/main.go
```

- 各種クライアント起動
```bash
$ go run chapter*/**/main.go
```