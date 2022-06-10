package main

import (
	"fmt"
)

type LinkedListNode struct {
	Value interface{}
	Next  *LinkedListNode
}

func printList(head *LinkedListNode) {
	i := head
	// iterate and print until i points off the end of the list
	for i != nil {
		fmt.Println(i.Value)
		i = i.Next
	}
}

func pushFront(head *LinkedListNode, value interface{}) {

	// move current head to new address
	oldHead := LinkedListNode{
		Value: head.Value,
		Next:  head.Next,
	}
	// assign new value to current head
	head.Value = value
	// assign head next old head address
	head.Next = &oldHead
}

func pushBack(head *LinkedListNode, value interface{}) {
	i := head
	// find last node, i.e. when next is nil
	for i.Next != nil {
		i = i.Next
	}
	// create new node
	last := LinkedListNode{
		Value: value,
		Next:  nil,
	}
	// assign next of last node to address of new node
	i.Next = &last
}

func insertAfter(pos *LinkedListNode, value interface{}) {
	// create new node with specified value and address of insert position node
	newNode := LinkedListNode{
		Value: value,
		Next:  pos.Next,
	}
	// assign next of post to address of new node
	pos.Next = &newNode
}

func reverseList(head *LinkedListNode) {
    curr := &LinkedListNode{
		Value: head.Value,
		Next: head.Next,
	}
    var prev *LinkedListNode
    var next *LinkedListNode
    for curr != nil {
        next = curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
	*head = *prev
}

func main() {

	// create linked list
	head := LinkedListNode{}
	second := LinkedListNode{}
	third := LinkedListNode{}

	head.Value = 1
	head.Next = &second

	second.Value = 2
	second.Next = &third

	third.Value = 3
	third.Next = nil

	pushFront(&head, "front")
	pushBack(&head, "back")
	insertAfter(&second, "insert")
	printList(&head)
	reverseList(&head)

	printList(&head)

}
