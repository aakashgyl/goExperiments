package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"io"
	"net/textproto"
//	"main/util"
)

const registrationURL = "http://10.71.11.188:8080/NE3S/1.0/NE3SRegistrationService"
const tmplLocation = "C:\\Users\\aagoyal\\eclipse-workspace-oxygen\\NDACServer\\src\\main\\resources\\tmpl\\"
const startRegTmpl = "startRegistration.xml"
const completeRegTmpl = "completeRegistration.xml"
const socFile = "C:\\Users\\aagoyal\\eclipse-workspace-oxygen\\NDACServer\\src\\main\\resources\\ne3sws_alarm_soc.xml"

var data Data

type ManagerData struct {
	UniqueId string
	Type     string
	Vendor   string
	Rel      string

	Nonce           string
	NotificationUrl string
	ReregInterval   string

	SocCid string
}

type AgentData struct {
	UniqueId string
}

type Data struct {
	Agent   AgentData
	Manager ManagerData
}

func getStartRegistrationReq() *http.Request {
	agent := new(AgentData)
	agent.UniqueId = "123456"

	manager := new(ManagerData)
	manager.Nonce = "MTgyMjcxMjAlMA=="
	manager.NotificationUrl = "http://10.142.148.239:8080/handle/"
	manager.UniqueId = "123456789"
	manager.Rel = "1.0"
	manager.ReregInterval = "10"
	manager.SocCid = "startRegSocCid"
	manager.Vendor = "NSN"
	manager.Type = "VS"

	data = Data{*agent, *manager}
	soapXml := getSoapBody(data, startRegTmpl)

	req, _ := loadSocFile( soapXml)

	return req
}

func getCompleteRegistrationReq(startRegResp string) *http.Request {
	data.Agent.UniqueId = "98765"
	soapXml := getSoapBody(data, completeRegTmpl)

	return createSoapRequest(soapXml, registrationURL)
}

func getSoapBody(data Data, tmpl string) string {

	if _, err := os.Stat(tmplLocation + tmpl); err == nil {
		dat, _ := ioutil.ReadFile(tmplLocation + tmpl)
		t := template.New("RegistrationTemplate")
		t, _ = t.Parse(string(dat))

		var tpl bytes.Buffer

		t.Execute(&tpl, data)
		return tpl.String()
	}
	
	return ""
}

func createSoapRequest(soapXml string, url string) *http.Request {
	payload := strings.NewReader(soapXml)

	req, _ := http.NewRequest("POST", url, payload)

	req.SetBasicAuth("wsuser", "wspassword")
	req.Header.Add("content-type", "text/xml")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("SOAPAction", "")
//	req.Header.Add("SOAPAction", "http://www.nokiasiemens.com/ne3s/1.0/startRegistration")

	return req
}

func loadSocFile(soapxml string) (*http.Request, error){
	//form multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	
	//create soap request
	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", "text/xml; charset=UTF-8")
	h.Set("Content-Transfer-Encoding", "8bit")
	h.Set("Content-ID", "<mainsoaprequest>")
	
	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, strings.NewReader(soapxml))
	

	//attach soc xml
	file, err := os.Open(socFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	h = make(textproto.MIMEHeader)
	h.Set("Content-Type", "text/xml; charset=UTF-8")
	h.Set("Content-ID", "<startRegSocCid>")
	
	part, err = writer.CreatePart(h)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)


	//close writer
	err = writer.Close()
	if err != nil {
		return nil, err
	}


	//final request
	req, err := http.NewRequest("POST", registrationURL, body)
	req.Header.Add("content-type", strings.Replace(writer.FormDataContentType(), "multipart/form-data", "multipart/related", -1) + "; type=\"text/xml\"; start=\"mainsoaprequest\"")
	
	req.SetBasicAuth("wsuser", "wspassword")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("SOAPAction", "")
	
	return req, err
}

func main() {
	//Start Registration
	req := getStartRegistrationReq()

		//	fmt.Println("REQ -> ", req)
		//	fmt.Println("HEAD -> ", req.Header)
		//	fmt.Println("BODY -> ", util.ReadCloserToString(req.Body))
	
	res, err := http.DefaultClient.Do(req)

	fmt.Println("Error on Start Registration Request-> ", err)

	body, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()
	
	fmt.Println("Response Received for Start Registration Request-> ", string(body))
	fmt.Println("Header -> ", res.Header)
	fmt.Println("Response -> ", res)
	
	
	//Complete Registration
	req = getCompleteRegistrationReq(string(body))
	res, err = http.DefaultClient.Do(req)
	
	fmt.Println("Error on Complete Registration Request -> ", err)

	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)

	fmt.Println("Response Received for Complete Registration Request-> ", string(body))
}