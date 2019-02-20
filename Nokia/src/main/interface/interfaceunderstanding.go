package main

import (
	"fmt"
	"reflect"
)

type One struct{
	data string
}

type Two struct{
	value int
}

func main() {
	data := One{data:"hello"}
	caller(data)
// 	t := reflect.TypeOf(pkg.My{})
// 	fmt.Println(t)

}

func caller (val interface{}){
	fmt.Println((reflect.TypeOf(val)).String())
}