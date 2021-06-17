package main

import (
	"fmt"
	"io"
	"net"
)

func main(){
	l, err := net.Listen("tcp", ":2000")
	if err != nil{
		fmt.Println(err)
	}

	defer l.Close()

	for{
		conn, err = l.Accept()
		if err != nil{
			fmt.Println(err)
		}

		go func(c net.Conn){
			io.Copy(c,c)
			c.Close()
		}(conn)
	}
}