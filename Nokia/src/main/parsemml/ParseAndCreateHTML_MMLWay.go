package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type alarmData struct {
	TimeStamp   string
	FaultId     string
	FaultSource string
	NumOfFaults string
	Curr10s     string
	Prev10s     string
	CurrMin     string
	PrevMin     string
	CurrHour    string
	PrevHour    string
	CurrDay     string
	PrevDay     string
	SinceReset  string
}

var tableHeadings = [13]string{"TimeStamp", "FaultId", "FaultSource", "NumOfFaults: (Unit)", "Curr 10s", "Prev 10s", "Curr Min", "Prev Min", "Curr Hour", "Prev Hour", "Curr Day", "Prev Day", "Since Reset"}
var activeAlarmData []alarmData
var inActiveAlarmData []alarmData
var alarmFile string = "C:\\Users\\aagoyal\\eclipse-workspace-oxygen\\NDACServer\\src\\main\\resources\\alarmFile.txt"
var templatename string = "C:\\Users\\aagoyal\\eclipse-workspace-oxygen\\NDACServer\\src\\main\\resources\\Report.tmpl"
var htmlreportname string = "C:\\Users\\aagoyal\\eclipse-workspace-oxygen\\NDACServer\\src\\main\\resources\\Report.html"

type AlarmData struct {
	ActiveAlarmData   []alarmData
	InActiveAlarmData []alarmData
}

func main() {
	readFileAndFillDS()
	generateHTMLReport()
}

func readFileAndFillDS() {
	f, _ := os.Open(alarmFile)
	scanner := bufio.NewScanner(f)

	var active int = 0
	var start int = 0

	//	data := "2004-01-01T00:00:04.000Z        (237   ) PowerCycleDetectedAl                              FSP1                                    (18  )            0     0     0     0     1     0     1     0     1     "
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Active Faults") {
			start = 1
			active = 1
		}
		if strings.Contains(line, "Inactive Faults") {
			active = 0
		}
		if strings.Count(line, "(") == 2 && start == 1 && active == 1 {
			updateAlarmArray(line, &activeAlarmData)
		}
		if strings.Count(line, "(") == 2 && start == 1 && active == 0 {
			updateAlarmArray(line, &inActiveAlarmData)
		}
	}
}

func updateAlarmArray(data string, alarmDataArray *[]alarmData) {
	// First compile the delimiter expression.
	re := regexp.MustCompile(`[()]`)
	result := re.Split(data, -1)

	alarmVal := new(alarmData)

	for i, v := range result {
		v = strings.TrimSpace(v)
		switch i {
		case 0:
			alarmVal.TimeStamp = v
		case 1:
			alarmVal.FaultId = "(" + v + ")"
		case 2:
			words := strings.Fields(v)
			alarmVal.FaultId = alarmVal.FaultId + words[0]
			alarmVal.FaultSource = words[1]
		case 3:
			alarmVal.NumOfFaults = "(" + v + ")"
		case 4:
			unpack(strings.Fields(v), &alarmVal.Curr10s, &alarmVal.Prev10s, &alarmVal.CurrMin, &alarmVal.PrevMin, &alarmVal.CurrHour, &alarmVal.PrevHour, &alarmVal.CurrDay, &alarmVal.PrevDay, &alarmVal.SinceReset)
		default:
			fmt.Println("Problem")
		}
	}
	
	*alarmDataArray = append(*alarmDataArray, *alarmVal)
}

func generateHTMLReport() {
	alarmDataMain := new(AlarmData)
	alarmDataMain.ActiveAlarmData = activeAlarmData
	alarmDataMain.InActiveAlarmData = inActiveAlarmData

	if _, err := os.Stat(templatename); err == nil {
		dat, _ := ioutil.ReadFile(templatename)
		t := template.New("Aakash Goyal")
		t, _ = t.Parse(string(dat))
		f, _ := os.Create(htmlreportname)
		t.Execute(f, alarmDataMain)
	}

}

func unpack(s []string, vars ...*string) {
	for i, str := range s {
		*vars[i] = str
	}
}
