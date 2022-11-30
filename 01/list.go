package main

import (
	"container/list"
	"fmt"
)

//Exemplo do container/list package

func main() {
	fmt.Println("----------------- list -----------------")
	var iniList list.List
	iniList.PushBack(10)
	iniList.PushBack(20)
	iniList.PushBack(30)

	for element := iniList.Front(); element != nil; element = element.Next() {
		println(element.Value.(int))
	}

	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
