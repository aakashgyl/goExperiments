package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	var luann Feedback
	xml.Unmarshal(feedback, &luann)
	fmt.Println(luann)
}

type Feedback struct {
	Name           string          `xml:"name"`
	Status         string          `xml:"status,attr"`
	Successful     string          `xml:"successful,attr"`
	Unsuccessful   string          `xml:"unsuccessful,attr"`
	MessageResult  string          `xml:"message_result"`
	MessageObjects []MessageObject `xml:"managedObjects>managedObject"`
}

type MessageObject struct {
	Status    string `xml:"status,attr"`
	DistName  string `xml:"distName,attr"`
	Class     string `xml:"class,attr"`
	Operation string `xml:"operation,attr"`
	ErrorText string `xml:"errorText"`
}

var feedback = []byte(`<?xml version="1.0" encoding="UTF-8"?>
	<NE3SFeedback
		status = "ok"
		successful = "1"
		unsuccessful = "0"
		version = "1.0"
		xmlns = "http://www.nsn.com/schemas/ne3s/feedback/1.0/">
		<message_result>Validation successful, configuration change validated:Activation successful:Provision of the plan was successful.</message_result>
		<managedObjects>
			<managedObject
				status = "ok"
				distName = "MRBTS-1/LNBTS-1/LNCEL-0"
				class = "LNCEL"
				operation = "update">
			</managedObject>
			<managedObject
				status = "not_ok"
				distName = "MRBTS-1/LNBTS-1/LNCEL-0"
				class = "LNCEL"
				operation = "update">
				<error>Information for the object could not be stored to the system.</error>
			</managedObject>
		</managedObjects>
	</NE3SFeedback>`)
