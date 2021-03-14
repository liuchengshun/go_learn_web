package main

import (
	"fmt"
	"net/http"
	"./gee"
)

// 写两个处理请求函数
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path:%q\n", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for i, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q \n", i, v)
	}
}

func main() {
	// new 一个实例请求
	r := gee.New()
	// 增加GET请求的路由
	r.GET("/", indexHandler)
	r.GET("/hello", helloHandler)
	// 启动服务
	r.Run(":9999")
}