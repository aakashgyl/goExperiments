package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://10.71.11.188:8080/view/ANewPage"
	
	response, _ := http.Get(url)
	defer response.Body.Close()
	
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(response)
	fmt.Println(string(body))
}
