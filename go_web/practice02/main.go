package main

import (
	"fmt"
	"strings"
	"net/http"
)

func sayHello (w http.ResponseWriter, r *http.Request) {
	// 解析参数
	r.ParseForm()
	fmt.Println(r.Form)
	// 路径
	fmt.Println("Path:", r.URL.Path)
	// 协议
	fmt.Println("scheme:", r.URL.Scheme)
	// 
	fmt.Println(r.Form["url_long"])
	for i, v := range r.Form {
		fmt.Println("key:", i)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello,go web")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("the error of http:", err)
		return
	}
}