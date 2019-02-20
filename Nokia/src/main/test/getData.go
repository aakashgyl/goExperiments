package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func main(){
	resp, err := http.Get("https://api.github.com/repositories/")
	if err != nil{
		fmt.Println("found err ", err)
	}
	fmt.Println(resp)
	
	data, _ := ioutil.ReadAll(resp.Body)
	processData(data)
}

func processData(data []byte) map[string]string{
	fulldata := []Data{}
	err := json.Unmarshal(data, &fulldata)
	fmt.Println(err)
	
	datamap := make(map[string]string)
	
	for _,val := range fulldata{
		fmt.Println(val.FullName)
		fmt.Println(val.Owner)
		datamap[val.FullName] = val.Owner.HtmlUrl
	}
	return datamap
}

type Data struct{
	FullName string `json:"full_name"`
	Owner Owner `json:"owner"`
}

type Owner struct{
	HtmlUrl string `json:"html_url"`
}