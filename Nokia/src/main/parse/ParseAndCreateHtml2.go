package main
//
//import (
//	"fmt"
//	"regexp"
//	"strings"
//	"os"
//	"bufio"
//)
//
//type alarmData struct{
//	TimeStamp string
//	FaultId	string
//	FaultSource	string
//	NumOfFaults string
//	Counts []string
//}
//
//var tableHeadings = [13]string{"TimeStamp", "FaultId", "FaultSource", "NumOfFaults: (Unit)", "Curr 10s", "Prev 10s", "Curr Min", "Prev Min", "Curr Hour", "Prev Hour", "Curr Day", "Prev Day", "Since Reset"}
//var activeAlarmData []alarmData
//var inActiveAlarmData []alarmData
//var alarmFile string = "C:\\Users\\aagoyal\\eclipse-workspace-oxygen\\NDACServer\\src\\main\\resources\\alarmFile.txt"
//
//func main() {
//	f, _ := os.Open(alarmFile)
//    scanner := bufio.NewScanner(f)
//
//	var active int = 0
//	var start int = 0
//	
////	data := "2004-01-01T00:00:04.000Z        (237   ) PowerCycleDetectedAl                              FSP1                                    (18  )            0     0     0     0     1     0     1     0     1     "
//    for scanner.Scan() {
//        line := scanner.Text()
//        
//        if strings.Contains(line, "Active Faults"){
//        	start = 1
//        	active = 1
//        }
//        if strings.Contains(line, "Inactive Faults"){
//        	active = 0
//        }
//        if strings.Count(line, "(") == 2 && start == 1 && active == 1 {
//	        updateAlarmArray(line, &activeAlarmData)	
//        }
//        if strings.Count(line, "(") == 2 && start == 1 && active == 0 {
//	        updateAlarmArray(line, &inActiveAlarmData)	
//        }
//    }
//    fmt.Println("Active -------> ")
//    fmt.Println(len(activeAlarmData))
//    fmt.Println("InActive -------> ")
//    fmt.Println(len(inActiveAlarmData))
//}
//
//func updateAlarmArray(data string, alarmDataArray *[]alarmData){
//// First compile the delimiter expression.
//	re := regexp.MustCompile(`[()]`)
//	result := re.Split(data, -1)
//	
//	alarmVal := new(alarmData)
//	
//	for i,v := range result {
//		v = strings.TrimSpace(v)
//		switch i {
//			case 0:
//				alarmVal.TimeStamp = v
//			case 1:
//				alarmVal.FaultId = "("+v+")"
//			case 2:
//				words := strings.Fields(v)
//				alarmVal.FaultId = alarmVal.FaultId + words[0]
//				alarmVal.FaultSource = words[1]
//			case 3:
//				alarmVal.NumOfFaults = "(" + v +")" 
//			case 4:
//				alarmVal.Counts = strings.Fields(v)
//			default:
//				fmt.Println("Problem")
//		}
//	}
////	fmt.Println("1 -> " + alarmVal.TimeStamp)
////	fmt.Println("2 -> " + alarmVal.FaultId)
////	fmt.Println("3 -> " + alarmVal.FaultSource)
////	fmt.Println("4 -> " + alarmVal.NumOfFaults)
////	fmt.Println(alarmVal.Counts)
////	fmt.Println(len(alarmVal.Counts))
//	
//	*alarmDataArray = append(*alarmDataArray, *alarmVal)
//}