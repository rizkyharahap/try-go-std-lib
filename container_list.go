package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("Rizki")
	data.PushBack("Mahfuddin")
	data.PushBack("Harahap")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	var next *list.Element = head.Next()
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	for el := data.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value)
	}
}
