package main

import (
	"fmt"
)

type LinkedList struct {
	Value interface{}
	Next  *LinkedList
}

func printList(i *LinkedList) {
	for i != nil {
		fmt.Println(i.Value)
		i = i.Next
	}
}

func pushFront(head *LinkedList, value interface{}) {

	// move current head to new address
	oldHead := LinkedList{
		Value: head.Value,
		Next:  head.Next,
	}
	// assign new value to current head
	head.Value = value
	// assign head next old head address
	head.Next = &oldHead
}

func main() {

	// create linked list
	head := LinkedList{}
	second := LinkedList{}
	third := LinkedList{}

	head.Value = 1
	head.Next = &second

	second.Value = 2
	second.Next = &third

	third.Value = 3
	third.Next = nil

	pushFront(&head, 11)

	printList(&head)

}
