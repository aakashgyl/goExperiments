package main

import (
	"fmt"
	"main/initfolder/foldera"
	"main/initfolder/folderab"
	"main/initfolder/config"
)

func init() {
	fmt.Println("Inside main")
}

func main(){
	fmt.Println("MAINNNNNNNNN")
	foldera.Hello()
	folderab.Hello()
	fmt.Println(foldera.A)
	fmt.Println(folderab.Ab)
	fmt.Println(config.MyVariable)
	config.MyVariable = 20
	fmt.Println(config.MyVariable)
}