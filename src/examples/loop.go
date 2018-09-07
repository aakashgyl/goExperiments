package main

import (
	"fmt"
)

func forLoop(i int) bool{
	for ; i<7; i++{
		fmt.Println(i)
	}
	
	return true
}

func forLoopInfi(i int) bool{
	for{
		fmt.Println(i)
		i++
		if i >8{
			break
		}else{
			continue
		}
	}
	return true
}