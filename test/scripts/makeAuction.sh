#!/bin/bash

url="http://localhost:8080/auction"

for i in $(seq 1 100); do
	# 引数を生成（A1からA100まで）
	# curlでPOSTリクエストを送信
	# レスポンスを見やすくするための改行
	arg="A${i}"
	curl -X POST "$url" -d "$arg"
	echo ""
done
# curl -X POST http://localhost:8080/auction -d "A1"