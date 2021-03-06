package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	c, _ := net.Dial("tcp", ":2300")
	var r io.Reader
	r = c //r now stores (value:c , type descriptor: net.Conn)
	//that's why we can also do this:
	if _, ok := r.(io.Writer); ok {
		/*
		   even though r in theory is only of type io.Reader,
		   the underlying value stored (c) also implements the io.writer interface
		   like r.c .(io.Writer)
		*/
		fmt.Println("We didn't forget there is a writer inside value c")
	}
}
