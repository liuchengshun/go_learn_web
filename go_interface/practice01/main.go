package main

import (
	"fmt"
	"strconv"
)

// fmt.Println可以接收实现了 String() string 方法的对象可以作为参数

type Human struct {
	name string
	age int
	phone string
}

func (h *Human) String() string {
	return "<" + h.name + "-" + strconv.Itoa(h.age) + "-" +"phoneNumber:" + h.phone + ">"
}

func main() {
	Bob := Human{"Bob", 24, "123456"}
	fmt.Println("This human is :", Bob)
}