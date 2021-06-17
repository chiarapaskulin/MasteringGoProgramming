package main

import "fmt"

//type SLLNode
type SLLNode2 struct {
	next  *SLLNode2
	value int
}

func (sNode *SLLNode2) SetValue(v int) {
	sNode.value = v
}

func (sNode *SLLNode2) GetValue() int {
	return sNode.value
}

func NewSLLNode2()*SLLNode2{
	return new(SLLNode2)
}

//type linked list
type SingleLinkedList struct {
	head *SLLNode2
	tail *SLLNode2
}

func newSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}

func (list *SingleLinkedList) Add(v int) {
	newNode := &SLLNode2{value: v}
	if list.head == nil {
		list.head = newNode
	} else if list.tail == list.head {
		list.head.next = newNode
	} else if list.tail != nil {
		list.tail.next = newNode
	}
	list.tail = newNode
}

func (list *SingleLinkedList) String() string {
	s := ""
	for n := list.head; n != nil ; n = n.next {
		s += fmt.Sprintf(" {%d} ", n.GetValue())
	}
	return s
}

func main() {
	list := newSingleLinkedList()
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Add(6)
	fmt.Println("Hello, playground", list)
}