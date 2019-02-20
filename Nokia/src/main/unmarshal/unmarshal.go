package main

import (
	"encoding/xml"
	"fmt"
)

var personXML = []byte(`<raml version = "2.1" xmlns = "raml21.xsd">
    <cmData type = "actual" scope = "all">
        <header>
            <log action = "get" dateTime = "2018-03-26T12:00:06+05:30"></log>
        </header>
        <managedObject
            class = "NOKLTE:AMGR"
            version = "FLF18SP_1803_01_1802_02"
            distName = "MRBTS-1/LNBTS-1/FTM-1/AMGR-1"
            operation = "create">
            <p name = "actParallelUserSessions">true</p>
            <p name = "checkCnumPasswdExpiry">false</p>
            <p name = "ldapConnectionType">TLS</p>
            <p name = "primaryLdapPort">389</p>
            <p name = "primaryLdapServer">0.0.0.0</p>
        </managedObject>
    </cmData>
</raml>`)

type Raml struct {
	Data []Data `xml:"cmData>managedObject>p"`
}

type Data struct{
	Name string `xml:"name,attr"`
	Value string `xml:",chardata"`
}

func maind() {
	var luann Raml
	xml.Unmarshal(personXML, &luann)
	fmt.Println(luann.Data)
	for i, val := range luann.Data{
		if val.Name == "ldapConnectionType"{
			fmt.Println(i, val.Value)
		}
	}
}