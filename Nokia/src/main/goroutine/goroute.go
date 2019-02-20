package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int,1)
	
	go dosomething(c)
	
	for{
		c <- 3
		time.Sleep(time.Second)
	}
	
}

func dosomething(c chan int){
	for i:=0;i<5;i++{
		val := <-c
		fmt.Println(val)
	}
	close(c)
}