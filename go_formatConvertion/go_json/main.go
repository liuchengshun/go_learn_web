package main

import (
	"fmt"
	"encoding/json"
)

type server struct {
	ServerName string
	ServerIP   string
}

type Servers struct {
	Server []server
}

func main() {
	v := &Servers {}
	value := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},
		{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	err := json.Unmarshal([]byte(value), &v)
	if err != nil {
		fmt.Println("the error of Unmarshal is :", err)
		return
	}
	fmt.Println("value :", value)
}