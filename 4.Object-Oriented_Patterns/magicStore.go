package main

import "fmt"

type magicStore struct {
	value interface{}
	name  string
}

func (ms *magicStore) SetValue(v interface{}) {
	ms.value = v
}

func (ms *magicStore) GetValue() interface{} {
	return ms.value
}

func NewMagicStoreWithName(name string) *magicStore {
	return new(magicStore)
}

func NewMagicStore() *magicStore {
	return new(magicStore)
}

func main() {
	istore := NewMagicStoreWithName("Integer Store")
	istore.SetValue(4)

	//check if istore.GetValue() is type int
	if v, ok := istore.GetValue().(int); ok{
		v *= 4
		fmt.Println(v)
	}

	sstore := NewMagicStoreWithName("String Store")
	sstore.SetValue("Hello")

	//check if sstore.GetValue() is type string
	if v, ok := sstore.GetValue().(string); ok{
		v += " World"
		fmt.Println(v)
	}

	mstore := NewMagicStore()
	mstore.SetValue("Hello")
	fmt.Println(mstore.GetValue())
}