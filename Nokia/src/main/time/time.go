package main
//
//import (
//	"fmt"
////	"strconv"
//	"time"
//)
//
//const (
//	// See http://golang.org/pkg/time/#Parse
//	timeFormat = "2006-01-02 15:04 UTC"
//)
//
//func main() {
//	fmt.Println("Before")
//	time.Sleep((time.Duration(9) * time.Second))
//	fmt.Println("After")
//	
////	reregDuration, _ := strconv.Atoi("10")
////	fmt.Println((reregDuration - 1) * 24)
//
////	timePrevReg := 
////	diff := time.Since(timePrevReg)
////	finaltime := (reregDuration-1)*24 - int(diff.Hours())
////	fmt.Println(finaltime)
//
//	//	RegisteredIPs = append(RegisteredIPs, ip)
//	//	log.Debug(configs.LogPrefix+"[enbops] RegisteredIPs updated to - ", RegisteredIPs)
//
//	//	go sleepAndReregister(regData, (reregDuration-1)*24-int(diff.Hours())) //register 24 hrs before
//
//	//	durationHrs == reregDuration*24
//	//	sleepTime := time.Duration(durationHrs) * time.Hour
//
//	//	v := "2018-05-18 12:00 UTC"
//	//	then, err := time.Parse(timeFormat, v)
//	//	fmt.Println(then)
//	//	fmt.Println(time.Now().UTC())
//	//
//	//	if err != nil {
//	//		fmt.Println(err)
//	//		return
//	//	}
//	//
//	//	duration := time.Since(then)
//	//	fmt.Println(duration.Hours())
//
//	////	timePrevReg, err := time.Parse("2006-01-02 15:04:05.0000000 +0000 UTC", "2018-05-18T03:17:46.399842904Z")
//	//	currTime := time.Now().UTC()
//	////	diff := currTime.Sub(timePrevReg)
//	//	prevTime := "20180518115854"
//	//	currTimeStr := currTime.Format("20060102150405")
//	//
//	//	prevTimeInt, _ := strconv.Atoi(prevTime)
//	//	currTimeInt, _ := strconv.Atoi(currTimeStr)
//	//
//	//	fmt.Println(currTimeInt - prevTimeInt, currTimeStr)
//	//
//	//	fmt.Println(currTime)
//	//	fmt.Println(timePrevReg, err)
//	//	fmt.Println(int(diff.Hours())) //register 24 hrs before
//}
