package main

import (
	log "github.com/sirupsen/logrus"
	"fmt"
)

func main(){
	data2 := `{"@target_index":"edge-ne3s-pm-74:fe:48:26:d1:c8-06-06-2018","@target_type":"pmdata","Dn":"NE-MRBTS-1/NE-LNBTS-6/FTM-1/PNASC-1","LTE_Port_Network_Access_Security_Stats_M51153C0":"0","LTE_Port_Network_Access_Security_Stats_M51153C1":"0","ObjectType":"PNASC","Time":"2004-01-14T06:00:00.000+05:30:00","edgename":"74:fe:48:26:d1:c8","nename":"FZM_GMC","neserialno":"EA151410135","request_id":"EA151410135_2004-01-14T06:00:00.000+05:30:00_NE-MRBTS-1_NE-LNBTS-6_FTM-1_PNASC-1"}`
	log.Info("aakash -> ", data2)
	fmt.Println(data2)
}