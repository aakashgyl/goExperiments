package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type SetData struct {
	CmData struct {
		InnerXML string `xml:",innerxml"`
	} `xml:"cmData"`
}

func main() {
	var setData SetData
	xml.Unmarshal([]byte(xmldata), &setData)
	val := setData.CmData.InnerXML
	val = strings.Replace(val, "<header />", "", -1)
	val = strings.Replace(val, "<header/>", "", -1)
	fmt.Println(val)
}

const xmldata = `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<raml version="2.1" xmlns="raml21.xsd">
<cmData type="actual" scope="all">
  <header>
   <log action = "get" dateTime="2018-03-26T12:00:06+05:30"></log>
  </header>
  <managedObject class="NOKLTE:LNMME" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNMME-0" operation="update">
   <p name="ipAddrPrim">172.60.1.50</p>
   <p name="transportNwId">0</p>
  </managedObject>
  <managedObject class="NOKLTE:LNMME" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNMME-1" operation="delete">
  </managedObject>
    <managedObject class="NOKLTE:LNMME" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNMME-2" operation="create">
   <p name="administrativeState">unlocked</p>
   <p name="ipAddrPrim">172.60.1.50</p>
   <p name="mmeRatSupport">Wideband-LTE</p>
   <p name="transportNwId">0</p>
  </managedObject>
<managedObject class="NOKLTE:IPNO" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-4/FTM-1/IPNO-1" operation="update">
   <p name="mainTransportNwId">0</p>
  </managedObject>
</cmData>
</raml>` 

const xmlData = `<?xml version="1.0" encoding="UTF-8"?>
<raml xmlns="raml21.xsd" version="2.1">
<cmData type="actual" scope="selection">
<header />
<managedObject class="NOKLTE:LNADJ" distName="MRBTS-1/LNBTS-4/LNADJ-2" operation="create" version="FLF18SP_1803_01_1802_02">
<p name="adjEnbId">2</p>
<p name="cPlaneIpAddr">20.20.20.2</p>
<p name="cPlaneIpAddrCtrl">oamControlled</p>
<p name="rlfBasedRCRsupported">false</p>
<p name="x2LinkStatus">unavailable</p>
</managedObject>

<managedObject class="NOKLTE:LNADJ" distName="MRBTS-1/LNBTS-6/LNADJ-0" operation="create" version="FLF18SP_1803_01_1802_02">
<p name="adjEnbId">0</p>
<p name="cPlaneIpAddr">40.20.40.0</p>
<p name="cPlaneIpAddrCtrl">oamControlled</p>
<p name="rlfBasedRCRsupported">false</p>
<p name="x2LinkStatus">unavailable</p>
</managedObject>
</cmData>
</raml>`


