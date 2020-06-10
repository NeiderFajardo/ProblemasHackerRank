package main

import "fmt"
import "container/list"

func main(){
	lista := list.New()
	lista.PushBack(2)
	lista.PushBack(3)
	lista.PushBack(4)
	lista.PushBack(5)
	lista.PushFront(1)

	for e := lista.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}