package main

import (
	"fmt"
	"net"
)

func main() {
	cln, err := net.Dial("tcp", "localhost:8087")
	if err != nil{
		fmt.Println("Error while dialing")
		return
	}
	fmt.Fprintln(cln, "hello")
	var output string
	fmt.Fscanf(cln, "%s\n", &output)
	fmt.Println(output)

	fmt.Fprintln(cln, "time")
	fmt.Fscanf(cln, "%s\n", &output)
	fmt.Println(output)

	fmt.Fprintln(cln, "close")
}

func hello(cln net.Conn) string{
	var output string
	fmt.Fprintln(cln, "hello")
	fmt.Fscanf(cln, "%s\n", &output)
	return output
}

type TPClient struct{
	conn net.Conn
}

func NewTPClient(c net.Conn) *TPClient{
	return &TPClient{conn:c}
}

func (c *TPClient)Hello() string{
	var output string
	fmt.Fprintln(c.conn, "hello")
	fmt.Fscanf(c.conn, "%s\n", &output)
	return output
}