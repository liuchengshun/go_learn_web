package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List [] Element

type Person struct {
	name string
	age int
}

//定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 123
	list[1] = "string"
	list[2] = Person{"Zhangsan", 18}

	for i, v := range list {
		if _, ok := v.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", i, v)
		} else if _, ok := v.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", i, v)
		} else if element, ok := v.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", i, v)
			fmt.Println("element:", element)
			} 
	}
}