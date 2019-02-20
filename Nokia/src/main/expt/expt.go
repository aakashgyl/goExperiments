package main

import (
	"fmt"
)



func main() {
	fmt.Println("Hello, playground")
//	a := "data"
//	b := []string{a}
//	fmt.Println(b)
	ch := make(chan bool, 2)
	ch <- true
	ch <- true
	close(ch)
	
	for i:=0;i<3;i++{
		a, v :=  <- ch
		fmt.Println(a,v)
	} 
}

//package main
//
//import (
//	"fmt"
//	"io/ioutil"
//	"os"
//)
//
//var filename string = "C:\\Users\\aagoyal\\eclipse-workspace-oxygen\\NDACServer\\src\\main\\resources\\datafile"
//
//func main() {
//	fmt.Println("Lets try it out")
//
//	var runes []rune
//
//	if _, err := os.Stat(filename + ".txt"); err == nil {
//		fmt.Println("Lets try it out Again...")
//		dat, _ := ioutil.ReadFile(filename + ".txt")
//
//		runes = []rune(string(dat))
//	}
//
//	f, err := os.Create(filename + "data.txt")
//	check(err)
//	fmt.Println("Coming", string(runes[0:123]))
//	
//	for i := 0; i < len(runes); i += 1024 * 1024 {
//		if i+1024*1024 > len(runes) {
//			fmt.Println("Ahha ", i, len(runes))
//			f.Write([]byte(string(runes[i:len(runes)])))
//			break
//		}
//		fmt.Println("Coming2")
//		n2, err := f.Write([]byte(string(runes[i:i+1024*1024])))
//		fmt.Println("Coming3")
//		check(err)
//		fmt.Printf("wrote %d bytes\n", n2)
//	}
//	fmt.Println("Coming4")
//}
//
//func check(e error) {
//    if e != nil {
//        panic(e)
//    }
//}