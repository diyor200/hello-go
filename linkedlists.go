package main

import (
	"fmt"
	"math/rand"
)

type Node struct {
	value int
	next  *Node
}

func linkedlist_func() {
	var head *Node
	head = nil

	// add 5 data
	fmt.Println("+++ add 5 data +++")
	n := 0
	for n < 5 {
		head = add(head, rand.Intn(20))
		fmt.Printf("data %d\n", n)
		n++
	}

	fmt.Println("+++ displaying ++")
	displayList(head)
}

func add(list *Node, data int) *Node {
	if list == nil {
		list := new(Node)
		list.value = data
		list.next = nil

		return list
	} else {
		temp := new(Node)
		temp.value = data
		temp.next = list

		list = temp

		return list
	}
}

func displayList(list *Node) {
	var temp *Node
	temp = list

	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}
