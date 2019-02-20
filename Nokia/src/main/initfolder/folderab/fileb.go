package folderab

import (
	"fmt"
)

var Ab int

func init() {
	fmt.Println("Inside folderb")
	Ab = 5
}

func Hello(){
	fmt.Println("Hello folderb")
}