package main

import "fmt"

type CMember struct {
	name      string
	age       int
	address   string
	rank      string
	clearance int
}

func main() {
	cm := CMember{
		name: "Kevin",
		clearance: 5,
		age: 32,
		rank: "Sergeant",
		address: "Station Mars",
	}

	fmt.Println("CM Clearance:", cm.clearance)
	cmp := &cm
	cmp.clearance = 4
	fmt.Println("CM Clearance:",cm.clearance)
	fmt.Printf("CMP Clearance: %d \n\n", cmp.clearance)

	var crew []CMember
	crew = append(crew, cm, CMember{"Jo", 32, "Station Mars", "Sergeant", 5})

	for i, v := range crew {
		fmt.Println(i, v)
	}

	fmt.Println("")

	//Maps with Structs
	/*
		var m map[string]CMember
		m = make(map[string]CMember)
		m["Keving"] = cm  //Keving->key & cm->value
	*/
	//or

	m := map[string]CMember{
		"Kevin": CMember{name: "Kevin", address: "Stations Mars"},
		"Jo":    CMember{name: "Jo", address: "Station Jupiter"},
	}

	//add
	m["Cisco"] = CMember{name: "Cisco", address: "Station Mars", clearance: 5}

	//retrieve
	elem := m["Cisco"]
	fmt.Printf("%v \n\n", elem)

	//check if the value exists
	v, ok := m["Jo"]
	fmt.Printf("%v %t \n\n", v, ok)

	//delete
	delete(m, "Jo")

	//print map
	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}
}

func (cm CMember) PrintSecurityClearance() {
	fmt.Println(cm.clearance)
}
