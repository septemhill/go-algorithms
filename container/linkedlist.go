package container

import (
	"fmt"
)

type Node struct {
	value int
	prev  *Node
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) isEmpty() bool {
	return (l.head == nil) && (l.tail == nil)
}

//Insert insert element to specified index
func (l *LinkedList) Insert(v, i int) {

}

//Append add element to linked-list tail
func (l *LinkedList) Append(v int) {
	n := &Node{v, nil, nil}

	if l.isEmpty() {
		l.head = n
		l.tail = n
	} else {
		l.tail.next = n
		n.prev = l.tail
		l.tail = n
	}
}

//Prepend add element to linked-list head
func (l *LinkedList) Prepend(v int) {
	n := &Node{v, nil, nil}

	if l.isEmpty() {
		l.head = n
		l.tail = n
	} else {
		l.head.prev = n
		n.next = l.head
		l.head = n
	}
}

//Remove remove element from specified index
func (l *LinkedList) Remove(i int) int {
	return 1
}

//RemoveHead remove element from linked-list head
func (l *LinkedList) RemoveHead() int {
	n, next := l.head, l.head.next

	n.prev = nil
	n.next = nil

	next.prev = nil
	l.head = next

	return n.value
}

//RemoveTail remove element from linked-list tail
func (l *LinkedList) RemoveTail() int {
	n, prev := l.tail, l.tail.prev

	n.next = nil
	n.prev = nil

	prev.next = nil
	l.tail = prev

	return n.value
}

//Trace traverse whole linked-list element from head
func (l LinkedList) Trace() {
	tmp := l.head

	for tmp != nil {
		fmt.Println(tmp.value)
		tmp = tmp.next
	}
}

//ReverseTrack traverse whole linked-list element from tail
func (l LinkedList) ReverseTrace() {
	tmp := l.tail

	for tmp != nil {
		fmt.Println(tmp.value)
		tmp = tmp.prev
	}
}
