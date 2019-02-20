package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func mainQuery() {

	url := "http://localhost:8080/SOAPExample/services/PersonServiceImpl"

	payload := strings.NewReader(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="http://service.jaxws.journaldev.com"><soapenv:Header/><soapenv:Body><ser:getAllPersons/></soapenv:Body></soapenv:Envelope>`)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "text/xml")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add( "SOAPAction", "");

	fmt.Println(req)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
