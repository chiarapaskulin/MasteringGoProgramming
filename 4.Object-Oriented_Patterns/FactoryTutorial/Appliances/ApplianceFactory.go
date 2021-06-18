package Appliances

import (
	"errors"
	"fmt"
)

type Appliance interface {
	Start()
	GetPurpose() string
}

const (
	STOVE = iota
	FRIDGE
	MICROWAVE
)

func CreateAppliance() (Appliance, error) {
	fmt.Println("Enter preferred appliance type")
	fmt.Println("0: Stove ")
	fmt.Println("1: Fridge")
	fmt.Println("2: Microwave")

	var myType int
	fmt.Scan(&myType)

	switch myType {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	case MICROWAVE:
		return new(Microwave), nil

	default:
		return nil, errors.New("Invalid Appliance Type")
	}
}
