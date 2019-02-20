package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type s struct {
	Int       int
	String    string
	ByteSlice []byte
	StSt      t
}

type t struct {
	NewSt string
	OldSt string
}

func main() {
//	b := t{"Aakash", "Goyal"}
	a := &s{Int: 42, String: "Hello World!", ByteSlice: []byte{0, 1, 2, 3, 4}}

	out, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	var c s

	fmt.Println(string(out), a)
	
	json.Unmarshal(out, &c)
	fmt.Printf("%+v\n", c)
//	fmt.Printf("%+v\n", dounmars(out, c))	
	
}

func dounmars(data []byte, structtype interface{}) (interface{}){
	fmt.Println(reflect.TypeOf(structtype))
	json.Unmarshal(data, &structtype)
	fmt.Println("Here -> ")
	fmt.Println(reflect.TypeOf(structtype))
	fmt.Println(structtype)
	var c s
	fmt.Println(reflect.TypeOf(c))
	json.Unmarshal(data, &c)
	fmt.Println(c)
	return structtype
}
