package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	resp, err := http.Get("http://api.theysaidso.com/qod")

	if err!=nil{
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)

	if err!= nil{
		fmt.Println(err)
		return
	}

	fmt.Println(string(contents))
}
