package main

import (
//	"fmt"
)

func main() {
	messages := make(chan string)
	go func() { <- messages}()
//	msg := <-messages
//	fmt.Println(msg)
}
