package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("http://10.71.11.188:8080/view/ANewPage")

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(response)
	fmt.Println(string(body))

}
