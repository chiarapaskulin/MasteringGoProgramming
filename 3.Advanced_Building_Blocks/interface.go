package main

import "fmt"

//Interface:
type Node interface {
	SetValue(v int)
	GetValue() int
}

//type SLLNode
type SLLNode struct {
	next  *SLLNode
	value int
}
func NewSLLNode()*SLLNode{
	return new(SLLNode)
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}
func (sNode *SLLNode) GetValue() int {
	return sNode.value
}


//type PowerNode
type PowerNode struct {
	next  *PowerNode
	value int
}
func NewPowerNode()*PowerNode{
	return new(PowerNode)
}
func (sNode *PowerNode) SetValue(v int) {
	sNode.value = v * 10
}

func (sNode *PowerNode) GetValue() int {
	return sNode.value
}


func main() {
	var node Node
	node = NewSLLNode()
	node.SetValue(4)
	fmt.Println("Node is of value ",node.GetValue())


	node = NewPowerNode()
	node.SetValue(5)
	fmt.Println("Node is of value ",node.GetValue())

	//if PowerNode is a concrete type that implements the node interface{}
	if n,ok := node.(*PowerNode); ok {
		fmt.Println(n.value)
	}
}