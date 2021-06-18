package main

import (
	"MasteringGoProgramming/4.Object-Oriented_Patterns/FactoryTutorial/Appliances"
	"fmt"
)

func main() {
	myAppliance, err := Appliances.CreateAppliance()

	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		fmt.Println(err)
	}

}