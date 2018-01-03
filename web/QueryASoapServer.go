package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://localhost:8080/SOAPExample/services/PersonServiceImpl"
	soapxml := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.jaxws.journaldev.com"><soapenv:Header/><soapenv:Body><ser:getAllPersons/></soapenv:Body></soapenv:Envelope>`
	payload := strings.NewReader(soapxml)

	request, _ := http.NewRequest("POST", url, payload)

	request.Header.Add("content-type", "text/xml")
	request.Header.Add("cache-control", "no-cache")
	request.Header.Add( "SOAPAction", "");

	fmt.Println(request)

	response, _ := http.DefaultClient.Do(request)
	defer response.Body.Close()
	
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(response)
	fmt.Println(string(body))

}
