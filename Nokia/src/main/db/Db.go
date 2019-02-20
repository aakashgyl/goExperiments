package main 

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome")
	dbobj := GetDbObj("aakash", "goyal")
	dbobj.Create()
	dbobj.Update()
}

type CRUD interface{
	Create()
	Read()
	Update()
	Delete()
}

type db struct{}

func (dbobj db) Create(){
	fmt.Println("create")	
}

func (dbobj db) Read(){
	fmt.Println("read")	
}

func (dbobj db) Update(){
	fmt.Println("update")	
}

func (dbobj db) Delete(){
	fmt.Println("delete")	
}

func GetDbObj(user, pass string) CRUD{
	dbobj := db{}
	return dbobj
}