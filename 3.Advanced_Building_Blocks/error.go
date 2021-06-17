package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var scMapping2 = map[string]int{
	"James" : 5,
	"Kevin" : 10,
	"Rahul" : 9,
}

//error message
var errCrewNotFound = errors.New("Crew member not found")

//type findError that implements Error interface (has Error() method)
type findError struct{
	Name, Server, Msg string
}

func(e findError) Error()string{
	return e.Msg
}

func findSC2(name, server string)(int,error){
	//simulate searching
	//time.Sleep(time.Duration(rand.Intn(50)) * time.Minute)

	if v,ok := scMapping2[name] ; !ok {
		//panic call
		//panic("Crew member not found")

		return -1, findError{name,server,"Crew member not found"}

		//other ways to return the error:
		//return -1, errors.New("Crew member not found")
		//return -1, fmt.Errorf("Crew member %s not found on server '%s'", name, server)
		//return -1, errCrewNotFound

	} else {
		return v,nil
	}

}

func main() {
	//panic recoverer
	rand.Seed(time.Now().UnixNano())
	defer func(){
		if err := recover(); err!=nil{
			fmt.Println("A Panic recovered ",err)
		}
	}()

	if i,err := findSC2("Ruku","server 1"); err!=nil{
		fmt.Println("Error occured while searching ",err)

		if err == errCrewNotFound {
			//do some logic here to handle the error
			fmt.Println("Confirmed error is errCrewNotFound")
		}

		//check if err is type findError and access it's attributes
		if v,ok := err.(findError); ok{
			fmt.Println("Server name ", v.Server)
			fmt.Println("Crew member name ", v.Name)
		}

	} else {
		fmt.Println("Crew member has security clearance",i)
	}
}