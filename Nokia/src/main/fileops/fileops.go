package main

import (
	"fmt"
	"os"
)

func main() {
	//	f, _ := os.OpenFile("jsonpm.json", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	//	f.Close()
	//	data := "Hare Krsna"
	//
	//	f, err := os.Create("xmlData_" + time.Now().Format("2006-01-02T15-04-05") + ".xml")
	//	if err != nil{
	//		fmt.Println(err)
	//	}
	//
	//	_, err=f.WriteString(data)
	//
	//	if err != nil{
	//		fmt.Println(err)
	//	}
	//
	//	f.Close()

	if _, err := os.Stat("/d/userdata/aagoyal/Desktop/Capture.PNG"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file does not exist")
			os.Exit(1)
		} else {
			fmt.Println("file some other error")
			os.Exit(2)
		}
	}
	fmt.Println("file exists")
	os.Exit(0)
}