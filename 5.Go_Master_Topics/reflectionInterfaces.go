package main

import (
	"fmt"
	"reflect"
)

type Printer interface {
	Print(s string)
}

type pStruct struct {
	s string
}

func (p *pStruct) Print(s string) {
	p.s = s
	fmt.Println(s)
}

func main() {
	p := new(pStruct)
	inspectType2(p)
}

func inspectType2(obj interface{}) {
	v := reflect.ValueOf(obj)
	t := v.Type()
	myInterface := reflect.TypeOf((*Printer)(nil)).Elem() //myInteface will be type *Printer

	//check if t (type *pStruct) implements the Printer interface
	fmt.Println("obj implements Printer?", t.Implements(myInterface))

	if t.Implements(myInterface) {
		printFunc := v.MethodByName("Print") //printFunc is a reflection value object that represents method Print
		args := []reflect.Value{reflect.ValueOf("Printing Hello")} //convert the string to a reflection value object and use it as a argument
		printFunc.Call(args) //use Call method and pass the argument to it
	}
}