package foldera

import (
	"fmt"
)

var A int

func init() {
	fmt.Println("Inside foldera")
	A = 4
}

func Hello(){
	fmt.Println("Hello foldera")
}