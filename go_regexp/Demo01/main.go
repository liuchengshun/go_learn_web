package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "regexp"
)

func main() {
	// 发出get请求
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("the error of http get:", err)
		return
	}
	defer resp.Body.Close()
	// 读取数据流
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("the error of read response:", err)
		return
	}
	// 将body转换为string
	src := string(body)

}