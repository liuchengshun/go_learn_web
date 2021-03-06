package main 

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("./server.xml")
	if err != nil {
		fmt.Println("the error of os.Open:", err)
		return
	}
	defer file.Close()	

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("the error of read data:", err)
		return
	}

	v := &Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("the error of Unmarshal data error:", err)
		return
	}
	fmt.Println("the result is:", *v)
}