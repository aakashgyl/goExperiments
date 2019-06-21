package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	var size int = 1000
	var breakinto int = 10
	var data []int
	fillMySlice(&data, size)
	fmt.Println(data)
	var result chan int
	result = make(chan int)
	go sumItUp(data, breakinto, size, result)
	hello := <- result
	log.Print("Sum is: ", hello)
}

func sumItUp(data []int, breakinto int, size int, result chan int) {
	var total int
	results := make(chan int, breakinto)

	for i:=0;i<breakinto;i++{
		go sumSliceData(data[i*(size/breakinto):(i+1)*(size/breakinto)], results)
	}

	for i:=0;i<breakinto;i++{
		total = total + <-results
	}
	result <- total
}

func sumSliceData(data []int, result chan int) {
	var total int
	for i:=0; i<len(data);i++{
		total = total + data[i]
	}
	result <- total
}


func fillMySlice(data *[]int, size int){
	for i:=0;i<size;i++{
		*data = append(*data, rand.Intn(1000))
	}
}