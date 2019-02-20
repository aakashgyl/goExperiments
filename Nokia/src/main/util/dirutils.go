package util

import (
	"fmt"
	"os"
)

func GetCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return "not found"
	}
	return dir
}
