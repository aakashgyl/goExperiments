package main

import (
	"fmt"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

func main() {
	var tst timestamp.Timestamp;
	fmt.Println("Hello, playground")
	tst.Seconds = 100
	now := time.Now()
    secs := now.Unix()
    fmt.Println(tst.String())
    fmt.Println(secs)
}