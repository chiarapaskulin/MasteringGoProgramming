package main

import "fmt"

func main(){
	X := 5
	CopyX(X)
	fmt.Println("Value of X:", X)

	ChangeX(&X)
	fmt.Println("New value of X:", X)
}

func CopyX(X int){
	X = 10
}

func ChangeX(X *int){
	*X = 10
}