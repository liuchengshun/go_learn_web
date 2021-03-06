package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello,world")
	})

	// 开启一个文件服务器，并指向文件
	fs := http.FileServer(http.Dir("./test.html"))
	// 将文件服务器指向URL
	http.Handle("/static/", fs)

	// 开启端口监听并服务
	http.ListenAndServe(":8080", nil)
}