package main

import (
	"fmt"
)

type LinkedListNode struct {
	Value int
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

func pushFront(head *LinkedListNode, value int) {

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

func pushBack(head *LinkedListNode, value int) {
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

func insertAfter(pos *LinkedListNode, value int) {
	// create new node with specified value and address of insert position node
	newNode := LinkedListNode{
		Value: value,
		Next:  pos.Next,
	}
	// assign next of post to address of new node
	pos.Next = &newNode
}

func reverseList(head *LinkedListNode) {
	// create cursor, next and previous to allow for swapping
	var cur = new(LinkedListNode)
	*cur = *head
	var next = new(LinkedListNode)
	var prev = new(LinkedListNode)

	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	*head = *prev
}

func mergeLists(a, b *LinkedListNode) *LinkedListNode {
	var dummy = new(LinkedListNode)
	var cursor = dummy

	for a != nil && b != nil {
		if a.Value < b.Value {
			cursor.Next = a
			a = a.Next
		} else {
			cursor.Next = b
			b = b.Next
		}
		cursor = cursor.Next
	}
	if a != nil {
		cursor.Next = a
	} else {
		cursor.Next = b
	}
	return dummy.Next
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

	var head2 = new(LinkedListNode)
	pushBack(head2, 2)
	pushBack(head2, 3)
	insertAfter(head2, 1)
	// printList(&head)
	// reverseList(&head)

	mergeLists(&head, head2)

	printList(&head)
}
