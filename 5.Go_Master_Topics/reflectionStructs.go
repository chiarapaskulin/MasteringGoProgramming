package main

import (
	"fmt"
	"reflect"
)

func main() {
	type myStruct struct {
		Field1 int     `alias:"f1" desc:"field number 1"`
		Field2 string  `alias:"f2" desc:"field number 2"`
		Field3 float64 `alias:"f3" desc:"field number 3"`
	}
	mys := myStruct{2, "Hello", 2.4}
	inspectType(&mys)
}

func inspectType(i interface{}) {
	mysRValue := reflect.ValueOf(i)

	//Kind() return the type of the value
	if mysRValue.Kind() == reflect.Ptr { //check if mysRValue is a pointer
		mysRValue = mysRValue.Elem() //if it is, keep it's element
	}

	if mysRValue.Kind() != reflect.Struct {
		fmt.Println("No struct type detected, cancelling", mysRValue.Kind())
		return
	} else {
		fmt.Println("Struct type detected, continuing")
	}

	mysRType := mysRValue.Type() //return the type object
	for i := 0; i < mysRType.NumField(); i++ {
		fieldRType := mysRType.Field(i)   //returns a struct field | datatype: StructField
		fieldRValue := mysRValue.Field(i) //returns a value

		fmt.Printf("Field Name: '%s' , field type: '%s' , field value: '%v' \n", fieldRType.Name, fieldRType.Type, fieldRValue.Interface())

		fmt.Println("struct tags, alias:", fieldRType.Tag.Get("alias"), "desc:", fieldRType.Tag.Get("desc"), fieldRValue.CanSet())
	}
}