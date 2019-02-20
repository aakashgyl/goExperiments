package main

import (
	// Other imports ...
	"fmt"
	xj "github.com/basgys/goxml2json"
	"strings"
)

func main() {
	// Extract data from restful.Request
	xml := strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?><app version="1.0"><request>1111</request></app>`)

	// Convert
	json, err := xj.Convert(xml)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(json)
	// ... Use JSON ...
}
