package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type MOData struct {
	MoClass string `xml:"class,attr"`
	Version string `xml:"version,attr"`
	MoName  string `xml:"distName,attr"`
}

type MOArray struct {
	MO []MOData `xml:"managedObject"`
}

var object MOArray

func main(){
	var resp = Data
//	var moData MOData

	_ = xml.Unmarshal([]byte(resp), &object)
	
	for _, data := range object.MO {
		if strings.HasSuffix(data.MoClass, "PMCADM") {
			fmt.Println(data)
		}
	}
	
	tryme()
}

func tryme(){
	for _, data := range object.MO {
		if strings.HasSuffix(data.MoClass, "LNCEL") {
			fmt.Println(data)
		}
	}
}

var Data = `<cmdata><managedObject class="NOKLTE:ACBPR" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/ACBPR-0" operation="create">
   <p name="eCallAcBarredStatic">Not barred</p>
  </managedObject>
  <managedObject class="NOKLTE:AMGR" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/AMGR-1" operation="create">
   <p name="actParallelUserSessions">true</p>
   <p name="checkCnumPasswdExpiry">false</p>
   <p name="ldapConnectionType">TLS</p>
   <p name="primaryLdapPort">389</p>
   <p name="primaryLdapServer">0.0.0.0</p>
  </managedObject>
  <managedObject class="NOKLTE:ANR" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/ANR-0" operation="create">
   <list name = "anrIdleTimeThresLte">
    <item>
   <p name="idleTimeThresLteNR">72</p>
   <p name="idleTimeThresNbEnbExch">18</p>
   <p name="idleTimeThresNbeNB">48</p>
   <p name="idleTimeThresX2">80</p>
   <p name="nbEnbExchWaitTmr">6</p>
    </item>
   </list>
   <p name="actAutoLteNeighRemoval">false</p>
   <p name="actAutoRttNeighRemoval">false</p>
   <p name="actAutoUtranNeighRemoval">false</p>
   <p name="anrIfTRSC">1000</p>
   <p name="anrRobLevel">middle</p>
   <p name="anrRobLevelUtran">middle</p>
   <p name="anrRttTRSCFS">5000ms</p>
   <p name="anrUtraTRSCFS">5000ms</p>
   <p name="autoUtraNeighRemovalThres">0</p>
   <p name="consecHoFailThres">3</p>
   <p name="consecUtranHoFailThres">3</p>
   <p name="idleTimeThresRttNR">24</p>
   <p name="idleTimeThresUtranNR">96</p>
   <p name="maxNumX2LinksIn">64</p>
   <p name="maxNumX2LinksOut">60</p>
   <p name="minNotActivatedRttRSCFS">5</p>
   <p name="minNotActivatedUtraRSCFS">1440</p>
   <p name="minNumThresRttNR">16</p>
   <p name="s1PrdRevalWaitTmr">60min</p>
   <p name="samNumThresRttNR">50</p>
   <p name="tTmpBlacklistRttANR">1440</p>
   <p name="utranPrdRevalWaitTmr">960min</p>
   <p name="x2PrdRevalWaitTmr">480min</p>
  </managedObject>
  <managedObject class="NOKLTE:ANRPRL" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/ANRPRL-0" operation="create">
   <p name="actAlsoForUeBasedANR">true</p>
   <p name="anrThresRSRPNbCell">-100</p>
   <p name="anrThresRSRPNbCellMobEv">-100</p>
   <p name="anrThresRSRQNbCell">-24</p>
   <p name="nrLimitInterFreq">16</p>
   <p name="nrLimitIntraFreq">32</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.eqm:ANTL" version="EQM18SP_FZM_1803_001" distName="MRBTS-1/EQM-1/APEQM-1/RMOD-1/ANTL-1" operation="create">
   <p name="antennaRoundTripDelay">45</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.eqm:ANTL" version="EQM18SP_FZM_1803_001" distName="MRBTS-1/EQM-1/APEQM-1/RMOD-1/ANTL-2" operation="create">
   <p name="antennaRoundTripDelay">45</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.eqmr:ANTL_R" version="EQMR18SP_FZM_1803_001" distName="MRBTS-1/EQM_R-1/APEQM_R-1/RMOD_R-1/ANTL_R-1" operation="create">
   <p name="configDN">MRBTS-1/EQM-1/APEQM-1/RMOD-1/ANTL-1</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.eqmr:ANTL_R" version="EQMR18SP_FZM_1803_001" distName="MRBTS-1/EQM_R-1/APEQM_R-1/RMOD_R-1/ANTL_R-2" operation="create">
   <p name="configDN">MRBTS-1/EQM-1/APEQM-1/RMOD-1/ANTL-2</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.eqm:APEQM" version="EQM18SP_FZM_1803_001" distName="MRBTS-1/EQM-1/APEQM-1" operation="create">
  </managedObject>
  <managedObject class="com.nokia.srbts.eqmr:APEQM_R" version="EQMR18SP_FZM_1803_001" distName="MRBTS-1/EQM_R-1/APEQM_R-1" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:BFDGRP" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/BFDGRP-8" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:BFDGRP" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/BFDGRP-7" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:BFDGRP" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/BFDGRP-6" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:BFDGRP" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/BFDGRP-5" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:BFDGRP" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/BFDGRP-4" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:BFDGRP" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/BFDGRP-1" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:BFDGRP" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/BFDGRP-2" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:BFDGRP" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/BFDGRP-3" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:BTOOTH" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/BTOOTH-1" operation="create">
   <list name = "pairingRelationships">
   </list>
   <p name="actBluetooth">disabled</p>
   <p name="actBluetoothTransAlarm">true</p>
   <p name="allowInqScanResp">false</p>
   <p name="allowLegacyPairing">true</p>
   <p name="allowSspPairing">false</p>
   <p name="btoothName"></p>
   <p name="legacyPairingOPTMode">-1</p>
   <p name="legacyPairingSecret"></p>
   <p name="pairingRelationTimer">3</p>
   <p name="sspSecret"></p>
  </managedObject>
  <managedObject class="com.nokia.srbts.eqm:CABINET" version="EQM18SP_FZM_1803_001" distName="MRBTS-1/EQM-1/APEQM-1/CABINET-1" operation="create">
  </managedObject>
  <managedObject class="com.nokia.srbts.eqmr:CABINET_R" version="EQMR18SP_FZM_1803_001" distName="MRBTS-1/EQM_R-1/APEQM_R-1/CABINET_R-1" operation="create">
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:CELLMAPPING" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/CELLMAPPING-1" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:CERTH" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/CERTH-1" operation="create">
   <list name = "BTSOperatorCertificate">
   </list>
   <list name = "CRLDistributionPoint">
   </list>
   <list name = "CRLInfo">
   </list>
   <p name="CRLDPSource">Certificate_extension</p>
   <p name="CRLUsageEnabled">Enabled</p>
   <p name="btsCertificateUpdateTime">90</p>
   <p name="btsDomainOfSubjectAltNames"></p>
   <p name="btsNumberOfSubjectAltNames">0</p>
   <p name="caCertificateUpdateTime">90</p>
   <p name="caSubjectName">0</p>
   <p name="cmpDirectory">pkix</p>
   <p name="cmpServerIpAddress">0.0.0.0</p>
   <p name="cmpServerPort">1024</p>
   <p name="crlUpdatePeriod">24</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:CHANNEL" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/CELLMAPPING-1/LCELL-1/CHANNELGROUP-1/CHANNEL-1" operation="create">
   <p name="antlDN">MRBTS-1/EQM-1/APEQM-1/RMOD-1/ANTL-1</p>
   <p name="direction">TX</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:CHANNEL" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/CELLMAPPING-1/LCELL-1/CHANNELGROUP-1/CHANNEL-2" operation="create">
   <p name="antlDN">MRBTS-1/EQM-1/APEQM-1/RMOD-1/ANTL-2</p>
   <p name="direction">TX</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:CHANNEL" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/CELLMAPPING-1/LCELL-1/CHANNELGROUP-1/CHANNEL-3" operation="create">
   <p name="antlDN">MRBTS-1/EQM-1/APEQM-1/RMOD-1/ANTL-1</p>
   <p name="direction">RX</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:CHANNEL" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/CELLMAPPING-1/LCELL-1/CHANNELGROUP-1/CHANNEL-4" operation="create">
   <p name="antlDN">MRBTS-1/EQM-1/APEQM-1/RMOD-1/ANTL-2</p>
   <p name="direction">RX</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:CHANNELGROUP" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/CELLMAPPING-1/LCELL-1/CHANNELGROUP-1" operation="create">
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:CLOCK" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/SYNC-1/CLOCK-1" operation="create">
   <list name = "syncInputList">
    <item>
   <p name="syncInputPrio">1</p>
   <p name="syncInputType">internal GNSS receiver</p>
    </item>
    <item>
   <p name="syncInputPrio">2</p>
   <p name="syncInputType">transport reference source</p>
    </item>
   </list>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:CLOCK_R" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/SYNC-1/CLOCK-1/CLOCK_R-1" operation="create">
   <p name="activeSyncMode">FreqSync</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:CMD" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/CMD-1" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:DRX" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNCEL-0/DRX-0" operation="create">
   <list name = "drxProfile1">
    <item>
   <p name="drxProfileIndex">1</p>
   <p name="drxProfilePriority">45</p>
    </item>
   </list>
   <list name = "drxProfile2">
    <item>
   <p name="drxInactivityT">4</p>
   <p name="drxLongCycle">40ms</p>
   <p name="drxOnDuratT">6</p>
   <p name="drxProfileIndex">2</p>
   <p name="drxProfilePriority">40</p>
   <p name="drxRetransT">4</p>
    </item>
   </list>
   <list name = "drxProfile3">
    <item>
   <p name="drxInactivityT">10</p>
   <p name="drxLongCycle">80ms</p>
   <p name="drxOnDuratT">6</p>
   <p name="drxProfileIndex">3</p>
   <p name="drxProfilePriority">30</p>
   <p name="drxRetransT">4</p>
    </item>
   </list>
   <list name = "drxProfile4">
    <item>
   <p name="drxInactivityT">500</p>
   <p name="drxLongCycle">320ms</p>
   <p name="drxOnDuratT">10</p>
   <p name="drxProfileIndex">4</p>
   <p name="drxProfilePriority">20</p>
   <p name="drxRetransT">16</p>
    </item>
   </list>
   <list name = "drxProfile5">
    <item>
   <p name="drxInactivityT">2560</p>
   <p name="drxLongCycle">2560ms</p>
   <p name="drxOnDuratT">10</p>
   <p name="drxProfileIndex">5</p>
   <p name="drxProfilePriority">10</p>
   <p name="drxRetransT">16</p>
    </item>
   </list>
   <p name="applyOutOfSyncState">extendedDrxOnly</p>
   <p name="drxApplyDeviceType">true</p>
   <p name="stInactFactor">1</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.eqm:EQM" version="EQM18SP_FZM_1803_001" distName="MRBTS-1/EQM-1" operation="create">
  </managedObject>
  <managedObject class="com.nokia.srbts.eqmr:EQM_R" version="EQMR18SP_FZM_1803_001" distName="MRBTS-1/EQM_R-1" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:ETHLK" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/ETHLK-1-1" operation="create">
   <list name = "l2VlanIdList">
   </list>
   <p name="acceptableFrameTypes">ADMIT_ALL</p>
   <p name="administrativeState">unlocked</p>
   <p name="ethBhType">Copper</p>
   <p name="l2BurstSize">0</p>
   <p name="l2IngressRate">RT_LINE_RATE</p>
   <p name="l2ShaperRate">RT_LINE_RATE</p>
   <p name="linkOAMDyingGaspEnabled">false</p>
   <p name="linkOAMEnabled">false</p>
   <p name="linkOAMProfile"></p>
   <p name="macAddr">60:a8:fe:b5:12:16</p>
   <p name="portDefaultPriority">0</p>
   <p name="portDefaultVlanId">1</p>
   <p name="speedAndDuplex">AUTODETECT</p>
  </managedObject>
  <managedObject class="NOKLTE:ETHLK" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/ETHLK-1-2" operation="create">
   <list name = "l2VlanIdList">
   </list>
   <p name="acceptableFrameTypes">ADMIT_ALL</p>
   <p name="administrativeState">unlocked</p>
   <p name="ethBhType">Copper</p>
   <p name="l2BurstSize">0</p>
   <p name="l2IngressRate">RT_LINE_RATE</p>
   <p name="l2ShaperRate">RT_LINE_RATE</p>
   <p name="linkOAMDyingGaspEnabled">false</p>
   <p name="linkOAMEnabled">false</p>
   <p name="linkOAMProfile"></p>
   <p name="macAddr">74:da:ea:56:58:ac</p>
   <p name="portDefaultPriority">0</p>
   <p name="portDefaultVlanId">1</p>
   <p name="speedAndDuplex">AUTODETECT</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:FEATCADM" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/FEATCADM-1" operation="create">
   <p name="actCli">false</p>
   <p name="actLocationLock">false</p>
   <p name="actRemoteRfDiag">false</p>
   <p name="actTemperatureReport">false</p>
   <p name="actTransportConfigFallback">false</p>
   <p name="systemAcctPermEnable">true</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:FMCADM" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/FMCADM-1" operation="create">
   <p name="actCategoryAlarms">false</p>
  </managedObject>
  <managedObject class="NOKLTE:FTM" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1" operation="create">
   <p name="authKeyData">BtDevAddress=A0:E6:F8:05:66:D9;BoardSerialNumber=4542313634363130373635;MacAddress=60a8feb51216;LeafNodeSerialNumber=4542313634393131363138;OscillatorData=;TestEngineeringTimestamp=3230313730333036313331373039;TestEngineeringDate=3030303030303030303030303030;BtDevName=D0815E1E</p>
   <p name="daisyChainPort">none</p>
   <p name="lmtPort">EIF2 (port B)</p>
   <p name="locationName">AH</p>
   <p name="primBackhaulPort">EIF1 (port A)</p>
   <p name="sfnPort">none</p>
   <p name="softwareReleaseVersion">FTM_L18SP_SR18SP_W19_G19_2018.03.21_1105</p>
   <p name="systemTitle">FZM_GMC</p>
   <p name="userLabel"></p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:GNSSI" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/SYNC-1/CLOCK-1/GNSSI-1" operation="create">
   <p name="actGnssOutputLnaPowerSupply">false</p>
   <p name="gnssControlMode">GPS-GLONASS</p>
   <p name="locationMode">automatic</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.eqmr:GNSSI_R" version="EQMR18SP_FZM_1803_001" distName="MRBTS-1/EQM_R-1/APEQM_R-1/CABINET_R-1/SMOD_R-1/GNSSI_R-1" operation="create">
   <p name="configDN">MRBTS-1/MNL-1/MNLENT-1/SYNC-1/CLOCK-1/GNSSI-1</p>
   <p name="gnssAntennaAltitude">0</p>
   <p name="gnssAntennaLatitude">0</p>
   <p name="gnssAntennaLongitude">0</p>
   <p name="gnssConstellationMode">GPS-GLONASS</p>
   <p name="gnssUnitName">FYGA</p>
  </managedObject>
  <managedObject class="NOKLTE:GTPU" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/GTPU-1" operation="create">
   <list name = "sgwIpAddressList">
    <item>
   <p name="sgwIpAddress">172.60.1.50</p>
   <p name="transportNwId">0</p>
    </item>
   </list>
   <p name="gtpuN3Reqs">5</p>
   <p name="gtpuPathSupint">60</p>
   <p name="gtpuT3Resp">2</p>
  </managedObject>
  <managedObject class="NOKLTE:IDNS" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/IDNS-1" operation="create">
   <p name="serverIpAddress">0.0.0.0</p>
   <p name="serverIpAddress2">0.0.0.0</p>
  </managedObject>
  <managedObject class="NOKLTE:IEIF" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/IEIF-1" operation="create">
   <p name="localIpAddr">10.53.8.26</p>
   <p name="localIpv6Addr">::</p>
   <p name="localIpv6PrefixLength">0</p>
   <p name="mtu">1500</p>
   <p name="netmask">255.255.255.0</p>
   <p name="overrideMTUSize">0</p>
   <p name="qosEnabled">true</p>
   <p name="sbs">4000</p>
   <p name="sbsTotal">4000</p>
   <p name="sir">1000000</p>
   <p name="sirTotal">1000000</p>
   <p name="trafficPathShapingEnable">TPS_OFF</p>
   <p name="upperLayerShaping">false</p>
   <p name="wfqSchedQueueWeight">1000</p>
  </managedObject>
  <managedObject class="NOKLTE:INTP" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/INTP-1" operation="create">
   <list name = "ntpServers">
   <p>10.171.8.4</p>
   </list>
   <p name="maxNtpTimeError">1500</p>
   <p name="ntpDscp">46</p>
   <p name="ntpSourceIpAddress">10.53.8.26</p>
   <p name="ntpSyncPortMode">Standard</p>
   <p name="ntpTimeLocalIpMapping">M-Plane</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.hwinv:INVUNIT_R" version="HWINV18SP_FZM_1712_001" distName="MRBTS-1/EQM_R-1/APEQM_R-1/CABINET_R-1/SMOD_R-1/INVUNIT_R-1" operation="create">
   <p name="inventoryUnitType">FWHH</p>
   <p name="vendorName">Nokia</p>
   <p name="vendorUnitFamilyType">FZM</p>
  </managedObject>
  <managedObject class="NOKLTE:IPNO" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1" operation="create">
   <list name = "addTransportNwIpAddrList">
   </list>
   <list name = "twampFlag">
   </list>
   <p name="actACNonOpBackhaul">false</p>
   <p name="actDualUPlaneIpAddress">false</p>
   <p name="actDynamicFirewall">false</p>
   <p name="actDynamicIP">false</p>
   <p name="actFastIpRerouting">false</p>
   <p name="actIPv6Mplane">false</p>
   <p name="actIpTnlMeasure">false</p>
   <p name="actIpv6">false</p>
   <p name="actLoSPropagation">false</p>
   <p name="actMultipleDynamicIP">false</p>
   <p name="actSeparationRanSharing">false</p>
   <p name="bfdHoldUpTime">0</p>
   <p name="btsId">1</p>
   <p name="btsSubnetMacAddr">00:00:00:00:00:00</p>
   <p name="cPlaneIpAddress">172.60.1.116</p>
   <p name="cPlaneIpAddressSec">0.0.0.0</p>
   <p name="cPlaneIpv6Address">::</p>
   <p name="cPlaneIpv6AddressSec">::</p>
   <p name="enableSoam">false</p>
   <p name="ftmBtsSubnetAddress">0.0.0.0</p>
   <p name="ftmBtsSubnetPrefixLength">0</p>
   <p name="icmpResponseEnabled">true</p>
   <p name="mPlaneIpAddress">10.53.8.26</p>
   <p name="mainTransportNwId">0</p>
   <p name="oamIpAddr">138.203.40.1</p>
   <p name="oamTlvReply">false</p>
   <p name="omsTls">forced</p>
   <p name="planeMacAddr">60:a8:fe:b5:12:17</p>
   <p name="retransTimer">1000</p>
   <p name="sPlaneIpAddress">10.53.8.26</p>
   <p name="secOmsIpAddr">0.0.0.0</p>
   <p name="servingOms">none</p>
   <p name="servingOmsAdminSetting">primaryForced</p>
   <p name="srcIpForCmpCrl">10.53.8.26</p>
   <p name="tlsRenegotiationInterval">86400</p>
   <p name="twampMessageRate">RATE_10</p>
   <p name="twampReflectorPort">5000</p>
   <p name="uPlane2IpAddress">0.0.0.0</p>
   <p name="uPlane2Ipv6Address">::</p>
   <p name="uPlaneIpAddress">172.60.1.116</p>
   <p name="uPlaneIpv6Address">::</p>
   <p name="wfqSchedOamWeight">10</p>
  </managedObject>
  <managedObject class="NOKLTE:IPRT" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/IPRT-1" operation="create">
   <list name = "staticRoutes">
    <item>
   <p name="bfdId">0</p>
   <p name="destIpAddr">172.60.1.0</p>
   <p name="gateway">172.60.1.113</p>
   <p name="netmask">255.255.255.0</p>
   <p name="preSrcIpv4Addr">0.0.0.0</p>
   <p name="preference">1</p>
    </item>
    <item>
   <p name="bfdId">0</p>
   <p name="destIpAddr">0.0.0.0</p>
   <p name="gateway">10.53.8.254</p>
   <p name="netmask">0.0.0.0</p>
   <p name="preSrcIpv4Addr">0.0.0.0</p>
   <p name="preference">1</p>
    </item>
   </list>
  </managedObject>
  <managedObject class="NOKLTE:IPRTV6" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/IPRTV6-1" operation="create">
   <list name = "staticIpv6Routes">
   </list>
  </managedObject>
  <managedObject class="NOKLTE:IPSECC" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPSECC-1" operation="create">
   <list name = "dynamicIpCfg">
    <item>
   <p name="cPlaneIpDynCfg"></p>
   <p name="cPlaneIpSecDynCfg"></p>
   <p name="cPlaneIpv6DynCfg"></p>
   <p name="cPlaneIpv6SecDynCfg"></p>
   <p name="cmpCrlIpDynCfg"></p>
   <p name="mPlaneIpDynCfg"></p>
   <p name="sPlaneIpDynCfg">DHCP</p>
   <p name="uPlaneIpDynCfg"></p>
   <p name="uPlaneIpv6DynCfg"></p>
    </item>
   </list>
   <list name = "ipSecTunnels">
   </list>
   <list name = "securityPolicies">
   </list>
   <list name = "x2IpsecProfile">
   </list>
   <list name = "x2TrafficSelectorList">
   </list>
   <p name="actIpSecBkupTun">false</p>
   <p name="authenticationScheme">CERTIFICATE</p>
   <p name="ipSecEmBypassCtrlIpAddr">0.0.0.0</p>
   <p name="ipSecEmBypassEnabled">false</p>
   <p name="ipSecEmBypassPingTimer">30</p>
   <p name="ipSecEmBypassState">Off</p>
   <p name="ipSecEnabled">false</p>
  </managedObject>
  <managedObject class="NOKLTE:IVIF" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/IEIF-1/IVIF-1" operation="create">
   <p name="localIpAddr">172.60.1.116</p>
   <p name="localIpv6Addr">::</p>
   <p name="localIpv6PrefixLength">0</p>
   <p name="netmask">255.255.255.240</p>
   <p name="qosEnabled">true</p>
   <p name="sbs">4000</p>
   <p name="sir">1000000</p>
   <p name="vlanId">3001</p>
   <p name="wfqSchedQueueWeight">1000</p>
  </managedObject>
  <managedObject class="NOKLTE:L2SWI" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/L2SWI-1" operation="create">
   <list name = "dscpMap">
    <item>
   <p name="dscp">0</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">1</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">2</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">3</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">4</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">5</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">6</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">7</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">8</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">9</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">10</p>
   <p name="priorityQueue">4</p>
    </item>
    <item>
   <p name="dscp">11</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">12</p>
   <p name="priorityQueue">4</p>
    </item>
    <item>
   <p name="dscp">13</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">14</p>
   <p name="priorityQueue">4</p>
    </item>
    <item>
   <p name="dscp">15</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">16</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">17</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">18</p>
   <p name="priorityQueue">4</p>
    </item>
    <item>
   <p name="dscp">19</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">20</p>
   <p name="priorityQueue">4</p>
    </item>
    <item>
   <p name="dscp">21</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">22</p>
   <p name="priorityQueue">4</p>
    </item>
    <item>
   <p name="dscp">23</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">24</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">25</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">26</p>
   <p name="priorityQueue">3</p>
    </item>
    <item>
   <p name="dscp">27</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">28</p>
   <p name="priorityQueue">3</p>
    </item>
    <item>
   <p name="dscp">29</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">30</p>
   <p name="priorityQueue">3</p>
    </item>
    <item>
   <p name="dscp">31</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">32</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">33</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="priorityQueue">2</p>
    </item>
    <item>
   <p name="dscp">35</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">36</p>
   <p name="priorityQueue">2</p>
    </item>
    <item>
   <p name="dscp">37</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">38</p>
   <p name="priorityQueue">2</p>
    </item>
    <item>
   <p name="dscp">39</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">40</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">41</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">42</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">43</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">44</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">45</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">46</p>
   <p name="priorityQueue">1</p>
    </item>
    <item>
   <p name="dscp">47</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">48</p>
   <p name="priorityQueue">1</p>
    </item>
    <item>
   <p name="dscp">49</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">50</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">51</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">52</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">53</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">54</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">55</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">56</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">57</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">58</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">59</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">60</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">61</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">62</p>
   <p name="priorityQueue">6</p>
    </item>
    <item>
   <p name="dscp">63</p>
   <p name="priorityQueue">6</p>
    </item>
   </list>
   <p name="enableLayer2Switching">false</p>
   <p name="l2PriorityQueueWeight2">8</p>
   <p name="l2PriorityQueueWeight3">4</p>
   <p name="l2PriorityQueueWeight4">1</p>
   <p name="l2PriorityQueueWeight5">1</p>
   <p name="l2PriorityQueueWeight6">1</p>
   <p name="portDefaultPriority">0</p>
   <p name="portDefaultVlanId">1</p>
   <p name="priorityQueueNonIP">1</p>
   <p name="priorityQueuePcp0">6</p>
   <p name="priorityQueuePcp1">4</p>
   <p name="priorityQueuePcp2">4</p>
   <p name="priorityQueuePcp3">4</p>
   <p name="priorityQueuePcp4">3</p>
   <p name="priorityQueuePcp5">3</p>
   <p name="priorityQueuePcp6">2</p>
   <p name="priorityQueuePcp7">1</p>
   <p name="priorityQueueUntagged">1</p>
   <p name="qosClassification">DSCP</p>
   <p name="vlanAwareSwitch">false</p>
  </managedObject>
  <managedObject class="NOKLTE:LBPUCCHRDPR" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LBPUCCHRDPR-0" operation="create">
   <p name="countdownPucchCompr">60min</p>
   <p name="countdownPucchExp">1min</p>
   <p name="rrcConnectedLowerThresh">10</p>
   <p name="rrcConnectedUpperThresh">80</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:LCELL" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/CELLMAPPING-1/LCELL-1" operation="create">
  </managedObject>
  <managedObject class="NOKLTE:LNADJ" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNADJ-1" operation="create">
   <p name="adjEnbId">1</p>
   <p name="cPlaneIpAddr">10.10.10.10</p>
   <p name="cPlaneIpAddrCtrl">oamControlled</p>
   <p name="rlfBasedRCRsupported">false</p>
   <p name="x2LinkStatus">unavailable</p>
  </managedObject>
  <managedObject class="NOKLTE:LNADJG" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNADJG-0" operation="create">
   <p name="arfcnValueGeran">10</p>
   <p name="bandIndicatorGeran">dcs1800</p>
   <p name="basestationColourCode">1</p>
   <p name="dtm">false</p>
   <p name="gTargetCi">11111</p>
   <p name="gTargetLac">222</p>
   <p name="mcc">111</p>
   <p name="mnc">12</p>
   <p name="mncLength">2</p>
   <p name="networkColourCode">1</p>
   <p name="networkControlOrder">NC0</p>
   <p name="rimStatus">disabled</p>
   <p name="systemInfoType">none</p>
  </managedObject>
  <managedObject class="NOKLTE:LNBTS" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1" operation="create">
   <list name = "amRlcPBTab1">
    <item>
   <p name="dlPollByte">25kB</p>
   <p name="ueCategory">1</p>
   <p name="ulPollByte">50kB</p>
    </item>
   </list>
   <list name = "amRlcPBTab2">
    <item>
   <p name="dlPollByte">125kB</p>
   <p name="ueCategory">2</p>
   <p name="ulPollByte">250kB</p>
    </item>
   </list>
   <list name = "amRlcPBTab3">
    <item>
   <p name="dlPollByte">250kB</p>
   <p name="ueCategory">3</p>
   <p name="ulPollByte">500kB</p>
    </item>
   </list>
   <list name = "amRlcPBTab4">
    <item>
   <p name="dlPollByte">250kB</p>
   <p name="ueCategory">4</p>
   <p name="ulPollByte">750kB</p>
    </item>
   </list>
   <list name = "amRlcPBTab5">
    <item>
   <p name="dlPollByte">375kB</p>
   <p name="ueCategory">5</p>
   <p name="ulPollByte">1000kB</p>
    </item>
   </list>
   <list name = "cipherPrefL">
    <item>
   <p name="eea0">1</p>
   <p name="eea1">9</p>
   <p name="eea2">9</p>
   <p name="eea3">8</p>
    </item>
   </list>
   <list name = "integrityPrefL">
    <item>
   <p name="eia0">1</p>
   <p name="eia1">9</p>
   <p name="eia2">9</p>
   <p name="eia3">8</p>
    </item>
   </list>
   <list name = "pdcpProf1">
    <item>
   <p name="pdcpProfileId">1</p>
   <p name="snSize">12bit</p>
   <p name="statusRepReq">3</p>
   <p name="tDiscard">infinity</p>
    </item>
   </list>
   <list name = "pdcpProf1001">
    <item>
   <p name="pdcpProfileId">1001</p>
   <p name="snSizeDl">18bit</p>
   <p name="snSizeUl">18bit</p>
   <p name="statusRepReq">0</p>
   <p name="tDiscard">300ms</p>
    </item>
   </list>
   <list name = "pdcpProf101">
    <item>
   <p name="pdcpProfileId">101</p>
   <p name="rohcMaxCid">4</p>
   <p name="snSize">7bit</p>
   <p name="tDiscard">100ms</p>
    </item>
   </list>
   <list name = "pdcpProf102">
    <item>
   <p name="pdcpProfileId">102</p>
   <p name="snSize">7bit</p>
   <p name="tDiscard">150ms</p>
    </item>
   </list>
   <list name = "pdcpProf103">
    <item>
   <p name="pdcpProfileId">103</p>
   <p name="snSize">12bit</p>
   <p name="tDiscard">100ms</p>
    </item>
   </list>
   <list name = "pdcpProf104">
    <item>
   <p name="pdcpProfileId">104</p>
   <p name="snSize">12bit</p>
   <p name="tDiscard">500ms</p>
    </item>
   </list>
   <list name = "pdcpProf2">
    <item>
   <p name="pdcpProfileId">2</p>
   <p name="snSize">12bit</p>
   <p name="statusRepReq">3</p>
   <p name="tDiscard">750ms</p>
    </item>
   </list>
   <list name = "pdcpProf3">
    <item>
   <p name="pdcpProfileId">3</p>
   <p name="snSize">12bit</p>
   <p name="statusRepReq">3</p>
   <p name="tDiscard">750ms</p>
    </item>
   </list>
   <list name = "pdcpProf4">
    <item>
   <p name="pdcpProfileId">4</p>
   <p name="snSize">12bit</p>
   <p name="statusRepReq">3</p>
   <p name="tDiscard">750ms</p>
    </item>
   </list>
   <list name = "pdcpProf5">
    <item>
   <p name="pdcpProfileId">5</p>
   <p name="snSize">12bit</p>
   <p name="statusRepReq">3</p>
   <p name="tDiscard">750ms</p>
    </item>
   </list>
   <list name = "qciTab1">
    <item>
   <p name="delayTarget">80ms</p>
   <p name="drxProfileIndex">2</p>
   <p name="dscp">46</p>
   <p name="enforceTtiBundling">true</p>
   <p name="lcgid">1</p>
   <p name="maxGbrDl">512</p>
   <p name="maxGbrUl">512</p>
   <p name="pdcpProfIdx">101</p>
   <p name="prio">20</p>
   <p name="qci">1</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">GBR</p>
   <p name="rlcMode">RLC_UM</p>
   <p name="rlcProfIdx">101</p>
   <p name="schedulBSD">100ms</p>
   <p name="schedulPrio">5</p>
    </item>
   </list>
   <list name = "qciTab2">
    <item>
   <p name="delayTarget">80ms</p>
   <p name="drxProfileIndex">2</p>
   <p name="dscp">26</p>
   <p name="enforceTtiBundling">false</p>
   <p name="l2OHFactorDL">16</p>
   <p name="l2OHFactorUL">35</p>
   <p name="lcgid">2</p>
   <p name="maxGbrDl">840</p>
   <p name="maxGbrUl">840</p>
   <p name="pdcpProfIdx">102</p>
   <p name="prio">40</p>
   <p name="qci">2</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">GBR</p>
   <p name="rlcMode">RLC_UM</p>
   <p name="rlcProfIdx">102</p>
   <p name="schedulBSD">100ms</p>
   <p name="schedulPrio">7</p>
    </item>
   </list>
   <list name = "qciTab3">
    <item>
   <p name="delayTarget">80ms</p>
   <p name="drxProfileIndex">2</p>
   <p name="dscp">46</p>
   <p name="enforceTtiBundling">false</p>
   <p name="l2OHFactorDL">16</p>
   <p name="l2OHFactorUL">35</p>
   <p name="lcgid">2</p>
   <p name="maxGbrDl">128</p>
   <p name="maxGbrUl">64</p>
   <p name="pdcpProfIdx">103</p>
   <p name="prio">30</p>
   <p name="qci">3</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">GBR</p>
   <p name="rlcMode">RLC_UM</p>
   <p name="rlcProfIdx">103</p>
   <p name="schedulBSD">100ms</p>
   <p name="schedulPrio">6</p>
    </item>
   </list>
   <list name = "qciTab4">
    <item>
   <p name="delayTarget">300ms</p>
   <p name="drxProfileIndex">2</p>
   <p name="dscp">28</p>
   <p name="enforceTtiBundling">false</p>
   <p name="l2OHFactorDL">16</p>
   <p name="l2OHFactorUL">35</p>
   <p name="lcgid">2</p>
   <p name="maxGbrDl">320</p>
   <p name="maxGbrUl">32</p>
   <p name="pdcpProfIdx">2</p>
   <p name="prio">50</p>
   <p name="qci">4</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">GBR</p>
   <p name="rlcMode">RLC_AM</p>
   <p name="rlcProfIdx">2</p>
   <p name="schedulBSD">300ms</p>
   <p name="schedulPrio">8</p>
    </item>
   </list>
   <list name = "qciTab5">
    <item>
   <p name="drxProfileIndex">3</p>
   <p name="dscp">34</p>
   <p name="enforceTtiBundling">false</p>
   <p name="lcgid">0</p>
   <p name="pdcpProfIdx">1</p>
   <p name="prio">10</p>
   <p name="qci">5</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">nonGBR</p>
   <p name="rlcMode">RLC_AM</p>
   <p name="rlcProfIdx">1</p>
   <p name="schedulBSD">100ms</p>
   <p name="schedulPrio">9</p>
   <p name="schedulType">SIGNALLING</p>
   <p name="schedulWeight">40</p>
    </item>
   </list>
   <list name = "qciTab6">
    <item>
   <p name="drxProfileIndex">3</p>
   <p name="dscp">18</p>
   <p name="enforceTtiBundling">false</p>
   <p name="lcgid">3</p>
   <p name="lteNrDualConnectSupport">disabled</p>
   <p name="pdcpProfIdx">2</p>
   <p name="prio">60</p>
   <p name="qci">6</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">nonGBR</p>
   <p name="rlcMode">RLC_AM</p>
   <p name="rlcProfIdx">2</p>
   <p name="rlcProfIdx3cc">6</p>
   <p name="schedulBSD">300ms</p>
   <p name="schedulPrio">9</p>
   <p name="schedulWeight">20</p>
    </item>
   </list>
   <list name = "qciTab65">
    <item>
   <p name="delayTarget">60ms</p>
   <p name="drxProfileIndex">2</p>
   <p name="dscp">46</p>
   <p name="enforceTtiBundling">true</p>
   <p name="lcgid">1</p>
   <p name="maxGbrDl">31</p>
   <p name="maxGbrUl">31</p>
   <p name="pdcpProfIdx">101</p>
   <p name="prio">7</p>
   <p name="qci">65</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">GBR</p>
   <p name="rlcMode">RLC_UM</p>
   <p name="rlcProfIdx">101</p>
   <p name="schedulBSD">100ms</p>
   <p name="schedulPrio">5</p>
    </item>
   </list>
   <list name = "qciTab66">
    <item>
   <p name="delayTarget">80ms</p>
   <p name="drxProfileIndex">2</p>
   <p name="dscp">46</p>
   <p name="enforceTtiBundling">true</p>
   <p name="lcgid">1</p>
   <p name="maxGbrDl">31</p>
   <p name="maxGbrUl">31</p>
   <p name="pdcpProfIdx">101</p>
   <p name="prio">20</p>
   <p name="qci">66</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">GBR</p>
   <p name="rlcMode">RLC_UM</p>
   <p name="rlcProfIdx">101</p>
   <p name="schedulBSD">100ms</p>
   <p name="schedulPrio">5</p>
    </item>
   </list>
   <list name = "qciTab69">
    <item>
   <p name="drxProfileIndex">3</p>
   <p name="dscp">34</p>
   <p name="enforceTtiBundling">false</p>
   <p name="lcgid">0</p>
   <p name="pdcpProfIdx">1</p>
   <p name="prio">5</p>
   <p name="qci">69</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">nonGBR</p>
   <p name="rlcMode">RLC_AM</p>
   <p name="rlcProfIdx">1</p>
   <p name="schedulBSD">100ms</p>
   <p name="schedulPrio">4</p>
    </item>
   </list>
   <list name = "qciTab7">
    <item>
   <p name="drxProfileIndex">3</p>
   <p name="dscp">20</p>
   <p name="enforceTtiBundling">false</p>
   <p name="lcgid">3</p>
   <p name="lteNrDualConnectSupport">disabled</p>
   <p name="pdcpProfIdx">2</p>
   <p name="prio">70</p>
   <p name="qci">7</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">nonGBR</p>
   <p name="rlcMode">RLC_AM</p>
   <p name="rlcProfIdx">2</p>
   <p name="rlcProfIdx3cc">6</p>
   <p name="schedulBSD">100ms</p>
   <p name="schedulPrio">10</p>
   <p name="schedulWeight">10</p>
    </item>
   </list>
   <list name = "qciTab70">
    <item>
   <p name="drxProfileIndex">3</p>
   <p name="dscp">18</p>
   <p name="enforceTtiBundling">false</p>
   <p name="lcgid">3</p>
   <p name="pdcpProfIdx">2</p>
   <p name="prio">55</p>
   <p name="qci">70</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">nonGBR</p>
   <p name="rlcMode">RLC_AM</p>
   <p name="rlcProfIdx">2</p>
   <p name="rlcProfIdx3cc">6</p>
   <p name="schedulBSD">300ms</p>
   <p name="schedulPrio">9</p>
   <p name="schedulWeight">30</p>
    </item>
   </list>
   <list name = "qciTab8">
    <item>
   <p name="drxProfileIndex">3</p>
   <p name="dscp">10</p>
   <p name="enforceTtiBundling">false</p>
   <p name="lcgid">3</p>
   <p name="lteNrDualConnectSupport">disabled</p>
   <p name="pdcpProfIdx">2</p>
   <p name="prio">80</p>
   <p name="qci">8</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">nonGBR</p>
   <p name="rlcMode">RLC_AM</p>
   <p name="rlcProfIdx">2</p>
   <p name="rlcProfIdx3cc">6</p>
   <p name="schedulBSD">300ms</p>
   <p name="schedulPrio">11</p>
   <p name="schedulWeight">5</p>
    </item>
   </list>
   <list name = "qciTab9">
    <item>
   <p name="drxProfileIndex">3</p>
   <p name="dscp">0</p>
   <p name="enforceTtiBundling">false</p>
   <p name="lcgid">3</p>
   <p name="lteNrDualConnectSupport">disabled</p>
   <p name="pdcpProfIdx">2</p>
   <p name="prio">90</p>
   <p name="qci">9</p>
   <p name="qciSupp">ENABLE</p>
   <p name="resType">nonGBR</p>
   <p name="rlcMode">RLC_AM</p>
   <p name="rlcProfIdx">2</p>
   <p name="rlcProfIdx3cc">6</p>
   <p name="schedulBSD">300ms</p>
   <p name="schedulPrio">12</p>
   <p name="schedulWeight">1</p>
    </item>
   </list>
   <list name = "rlcProf1">
    <item>
   <p name="maxRetxThresh">t16</p>
   <p name="pollPdu">infinity</p>
   <p name="rlcProfileId">1</p>
   <p name="tPollRetr">120ms</p>
   <p name="tProhib">50ms</p>
   <p name="tReord">50ms</p>
   <p name="tReordIsca">50ms</p>
    </item>
   </list>
   <list name = "rlcProf101">
    <item>
   <p name="rlcProfileId">101</p>
   <p name="snFieldLengthDL">5bit</p>
   <p name="snFieldLengthUL">5bit</p>
   <p name="tReord">50ms</p>
    </item>
   </list>
   <list name = "rlcProf102">
    <item>
   <p name="rlcProfileId">102</p>
   <p name="snFieldLengthDL">5bit</p>
   <p name="snFieldLengthUL">5bit</p>
   <p name="tReord">50ms</p>
    </item>
   </list>
   <list name = "rlcProf103">
    <item>
   <p name="rlcProfileId">103</p>
   <p name="snFieldLengthDL">10bit</p>
   <p name="snFieldLengthUL">10bit</p>
   <p name="tReord">50ms</p>
    </item>
   </list>
   <list name = "rlcProf104">
    <item>
   <p name="rlcProfileId">104</p>
   <p name="snFieldLengthDL">10bit</p>
   <p name="snFieldLengthUL">5bit</p>
   <p name="tReord">50ms</p>
    </item>
   </list>
   <list name = "rlcProf2">
    <item>
   <p name="maxRetxThresh">t16</p>
   <p name="pollPdu">64</p>
   <p name="rlcProfileId">2</p>
   <p name="tPollRetr">120ms</p>
   <p name="tProhib">50ms</p>
   <p name="tReord">50ms</p>
   <p name="tReordIsca">50ms</p>
    </item>
   </list>
   <list name = "rlcProf3">
    <item>
   <p name="maxRetxThresh">t16</p>
   <p name="pollPdu">64</p>
   <p name="rlcProfileId">3</p>
   <p name="tPollRetr">120ms</p>
   <p name="tProhib">50ms</p>
   <p name="tReord">50ms</p>
   <p name="tReordIsca">50ms</p>
    </item>
   </list>
   <list name = "rlcProf4">
    <item>
   <p name="maxRetxThresh">t16</p>
   <p name="pollPdu">64</p>
   <p name="rlcProfileId">4</p>
   <p name="tPollRetr">120ms</p>
   <p name="tProhib">50ms</p>
   <p name="tReord">50ms</p>
   <p name="tReordIsca">50ms</p>
    </item>
   </list>
   <list name = "rlcProf5">
    <item>
   <p name="maxRetxThresh">t16</p>
   <p name="pollPdu">64</p>
   <p name="rlcProfileId">5</p>
   <p name="tPollRetr">120ms</p>
   <p name="tProhib">50ms</p>
   <p name="tReord">50ms</p>
   <p name="tReordIsca">50ms</p>
    </item>
   </list>
   <list name = "rlcProf6">
    <item>
   <p name="maxRetxThresh">t16</p>
   <p name="pollPdu">64</p>
   <p name="rlcProfileId">6</p>
   <p name="tPollRetr">95ms</p>
   <p name="tProhib">25ms</p>
   <p name="tReord">25ms</p>
   <p name="tReordIsca">50ms</p>
    </item>
   </list>
   <p name="act1xCsfb">false</p>
   <p name="act8EpsBearers">false</p>
   <p name="actA3ScellSelect">false</p>
   <p name="actAcBarringCFR">false</p>
   <p name="actAcBarringRrcConn">None</p>
   <p name="actAcBarringRrcReq">false</p>
   <p name="actAcSrvcc">false</p>
   <p name="actAdvScellMeas">false</p>
   <p name="actAnrRtt">false</p>
   <p name="actArpBasedQosProfiling">false</p>
   <p name="actAutoAcBarring">false</p>
   <p name="actAutoPlmnRemoval">false</p>
   <p name="actCMAS">false</p>
   <p name="actCSFBRedir">Disabled</p>
   <p name="actCellTrace">false</p>
   <p name="actCellTraceWithIMSI">false</p>
   <p name="actCiphering">true</p>
   <p name="actCoMp">disabled</p>
   <p name="actCompChecks">true</p>
   <p name="actContextPreemption">false</p>
   <p name="actCplaneOvlHandling">false</p>
   <p name="actCsfbECRestrRem">false</p>
   <p name="actCsfbPsHoToUtra">false</p>
   <p name="actCsgS1Mobility">false</p>
   <p name="actDLCAggr">false</p>
   <p name="actDualRx1xCsfb">false</p>
   <p name="actDynMBMSRes">false</p>
   <p name="actERabModify">true</p>
   <p name="actETWS">false</p>
   <p name="actEmerCallRedir">Disabled</p>
   <p name="actEnhAcAndGbrServices">true</p>
   <p name="actExtMeasCtrl">false</p>
   <p name="actFlexQCIARPPMProfiles">false</p>
   <p name="actFlexScellSelect">false</p>
   <p name="actGsmRedirWithSI">false</p>
   <p name="actHOtoHrpd">false</p>
   <p name="actHOtoWcdma">false</p>
   <p name="actHeNBMobility">false</p>
   <p name="actHeNBS1Mobility">false</p>
   <p name="actHoFromUtran">false</p>
   <p name="actHybridS1Mobility">false</p>
   <p name="actIPv6MBMS">false</p>
   <p name="actIdleLB">false</p>
   <p name="actIdleLBCaAware">false</p>
   <p name="actIdleLBCaAwareAdv">false</p>
   <p name="actIfHo">enabled</p>
   <p name="actImmHRPD">false</p>
   <p name="actImmXrtt">false</p>
   <p name="actInterEnbDLCAggr">false</p>
   <p name="actInterEnbULCAggr">false</p>
   <p name="actInterFreqMDTCellTrace">false</p>
   <p name="actInterFreqRstdMeasSupp">false</p>
   <p name="actIntraFreqPciSharing">false</p>
   <p name="actL1PM">false</p>
   <p name="actLBPowerSaving">false</p>
   <p name="actLPPaEcid">false</p>
   <p name="actLPPaOtdoa">false</p>
   <p name="actLTES1Ho">true</p>
   <p name="actLocRep">false</p>
   <p name="actLteNrDualConnectivity">false</p>
   <p name="actMBMS">false</p>
   <p name="actMBMSPS">false</p>
   <p name="actMBMSServiceContinuity">false</p>
   <p name="actMDTCellTrace">false</p>
   <p name="actMDTSubscriberTrace">false</p>
   <p name="actMDTloggedCellTrace">false</p>
   <p name="actMeasBasedIMLB">false</p>
   <p name="actMmecPlmnIdMmeSelection">false</p>
   <p name="actMroInterRatUtran">false</p>
   <p name="actMultBearers">true</p>
   <p name="actMultGbrBearers">false</p>
   <p name="actMultipleCarrier">false</p>
   <p name="actNonGbrServiceDiff">true</p>
   <p name="actNwReqUeCapa">false</p>
   <p name="actOTNRecovery">false</p>
   <p name="actOperatorQCI">false</p>
   <p name="actOperatorQCIGBR">false</p>
   <p name="actOptMmeSelection">false</p>
   <p name="actPCMDReport">false</p>
   <p name="actPartialAcHo">false</p>
   <p name="actPdcpRohc">true</p>
   <p name="actPlmnBasedPreemption">false</p>
   <p name="actQCIPLMNIDProfiles">false</p>
   <p name="actRIMforGSM">false</p>
   <p name="actRIMforUTRAN">false</p>
   <p name="actRLFReportEval">false</p>
   <p name="actRLFbasedRCR">false</p>
   <p name="actRRCConnReestablRLF">false</p>
   <p name="actRSRPRSRQHist">false</p>
   <p name="actRadioPosFlexibility">false</p>
   <p name="actRedirect">enabled</p>
   <p name="actRrcConnNoActivity">false</p>
   <p name="actRsrqInterFreqMobility">false</p>
   <p name="actRsrqInterRatMobility">false</p>
   <p name="actRtPerfMonitoring">false</p>
   <p name="actS1Flex">enabled</p>
   <p name="actS1OlHandling">true</p>
   <p name="actSatBackhaul">false</p>
   <p name="actSdl">false</p>
   <p name="actSelMobPrf">false</p>
   <p name="actServBasedMobThr">false</p>
   <p name="actSubscriberTrace">false</p>
   <p name="actTaHistCounters">false</p>
   <p name="actTcpBoost">false</p>
   <p name="actTempHoBlacklisting">false</p>
   <p name="actULCAggr">false</p>
   <p name="actUTCBroadcast">false</p>
   <p name="actUeBasedAnrInterFreqLte">false</p>
   <p name="actUeBasedAnrIntraFreqLte">false</p>
   <p name="actUeBasedAnrUtran">false</p>
   <p name="actUeLevelMro">false</p>
   <p name="actUeLocInfo">false</p>
   <p name="actUeThroughputMeas">false</p>
   <p name="actUlInDevCoexGNSS">false</p>
   <p name="actUplaneOvlHandling">false</p>
   <p name="actUserLayerTcpMssClamping">false</p>
   <p name="actVendSpecCellTraceEnh">false</p>
   <p name="actVoLTEQualSRVCCtoGsm">false</p>
   <p name="actZUC">false</p>
   <p name="actZson">false</p>
   <p name="acteNACCtoGSM">false</p>
   <p name="anrOmExtEnable">true</p>
   <p name="btsType">Flexi Zone Indoor Pico</p>
   <p name="caIntraCellHoMode">disabled</p>
   <p name="caMinDlAmbr">512</p>
   <p name="caSchedFairFact">0</p>
   <p name="congWeightAlg">qciPrio</p>
   <p name="defProfIdxAM">2</p>
   <p name="defProfIdxUM">104</p>
   <p name="enableAutoLock">Enabled</p>
   <p name="enableBwCombCheck">true</p>
   <p name="enableGrflShdn">Enabled</p>
   <p name="enbName">FZM-GMC</p>
   <p name="etwsPrimNotifBcDur">60</p>
   <p name="forcedPlanFileActivation">false</p>
   <p name="keyRefrMarg">50000</p>
   <p name="maxNumAnrMoiAllowed">4000</p>
   <p name="maxNumClusterUe">10000</p>
   <p name="maxNumOfLnadjLimit">128</p>
   <p name="maxNumPreEmptions">20</p>
   <p name="maxParallelIncX2SetupOrECUOvl0">10</p>
   <p name="maxParallelIncX2SetupOrECUOvl1">5</p>
   <p name="maxParallelOutX2SetupOrECU">5</p>
   <p name="maxRetxThreshSrbDL">t16</p>
   <p name="maxRetxThreshSrbUL">t16</p>
   <p name="maxX2CfUpRetry">5</p>
   <p name="mcc">244</p>
   <p name="mmeAssignmentMode">selectZeroCapacityToo</p>
   <p name="mnc">9</p>
   <p name="mncLength">2</p>
   <p name="moProfileSelect">spid</p>
   <p name="multScellReleaseTimer">0</p>
   <p name="nCqiDtx">100</p>
   <p name="nCqiRec">2</p>
   <p name="neighbCellChkPeriodicity">0</p>
   <p name="neighbCellChkStartTime">4</p>
   <p name="nullFallback">true</p>
   <p name="operationalState">configured</p>
   <p name="otnRecoveryPeriod">500</p>
   <p name="pbrNonGbr">8kBps</p>
   <p name="planFileActivationMode">totalServiceLossAllowed</p>
   <p name="prioTopoHO">all equal</p>
   <p name="prohibitLBHOTimer">10s</p>
   <p name="pwrFallbackCa">true</p>
   <p name="pwsWithEmAreaId">false</p>
   <p name="rachAccessForHoFromUtran">contentionBased</p>
   <p name="recoveryResetDelay">0</p>
   <p name="rlfBasedRCRdefault">false</p>
   <p name="rrcGuardTimer">20</p>
   <p name="s1InducedCellDeactDelayTime">1</p>
   <p name="sCellActivationCyclePeriod">0.001s</p>
   <p name="sCellActivationMethod">nonGBRBufferBasedStepWise</p>
   <p name="sCellDeactivationTimereNB">rf128</p>
   <p name="sCellpCellHARQFdbkUsage">conservative</p>
   <p name="scellActivationLevel">1TTI</p>
   <p name="scellMeasCycle">sf320</p>
   <p name="shutdownStepAmount">6</p>
   <p name="shutdownWindow">60</p>
   <p name="srvccDelayTimer">0</p>
   <p name="supportedCellTechnology">FDD</p>
   <p name="tChangeMobilityReq">5</p>
   <p name="tHalfRrcCon">3000</p>
   <p name="tHoOverallD">20</p>
   <p name="tPollRetrSrbDL">100ms</p>
   <p name="tPollRetrSrbUL">100ms</p>
   <p name="tRsrInitWait">300</p>
   <p name="tRsrResFirst">10</p>
   <p name="tS1RelOvDeltG">1000</p>
   <p name="tS1RelOvDeltU">1000</p>
   <p name="tS1RelPrepG">1000</p>
   <p name="tS1RelPrepL">500</p>
   <p name="tS1RelPrepU">2000</p>
   <p name="tX2RelocPrep">500</p>
   <p name="tagMaxAM">10</p>
   <p name="tagMaxUM">10</p>
   <p name="timDelACContPreempt">100</p>
   <p name="tx2RelODelta">150</p>
   <p name="txPathFailureMode">keepCellInService</p>
   <p name="userLayerTcpMss">0</p>
   <p name="voiceSuppMatchInd">errorIndication</p>
   <p name="x2CfUpRdmDelayTmr">3</p>
  </managedObject>
  <managedObject class="NOKLTE:LNBTS_FDD" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNBTS_FDD-0" operation="create">
   <p name="act1xSrvcc">false</p>
   <p name="actConvVoice">true</p>
   <p name="actDedFreqListsLB">none</p>
   <p name="actDedVoLteInterFreqHo">false</p>
   <p name="actDistributedSite">false</p>
   <p name="actDlIntShaping">false</p>
   <p name="actDualBand">false</p>
   <p name="actFlexBbUsage">false</p>
   <p name="actHighPrioServices">false</p>
   <p name="actHighRrc">false</p>
   <p name="actIMSEmerSessR9">true</p>
   <p name="actInDevCoexLaa">false</p>
   <p name="actInterFreqServiceBasedHo">false</p>
   <p name="actMFBI">false</p>
   <p name="actMultefire">false</p>
   <p name="actPeriodicCarrierBlinking">false</p>
   <p name="actPubSafetyBearers">false</p>
   <p name="actSfn">false</p>
   <p name="actSrvccToGsm">false</p>
   <p name="actSrvccToWcdma">false</p>
   <p name="actTempRadioMaster">false</p>
   <p name="actUnlicensedAcc">off</p>
   <p name="actUnlicensedDfs">false</p>
   <p name="actUtranLoadBal">false</p>
   <p name="numTxWithHighNonGbr">0</p>
   <p name="preventPsHOtoWcdma">None</p>
  </managedObject>
  <managedObject class="NOKLTE:LNCEL" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNCEL-0" operation="create">
   <list name = "dFListPucch">
    <item>
   <p name="dFpucchF1">0</p>
   <p name="dFpucchF1b">1</p>
   <p name="dFpucchF2">0</p>
   <p name="dFpucchF2a">0</p>
   <p name="dFpucchF2b">0</p>
    </item>
   </list>
   <list name = "drxProfile101">
    <item>
   <p name="drxInactivityT">10</p>
   <p name="drxLongCycle">160ms</p>
   <p name="drxOnDuratT">6</p>
   <p name="drxProfileIndex">101</p>
   <p name="drxProfilePriority">30</p>
   <p name="drxRetransT">4</p>
    </item>
   </list>
   <list name = "drxProfile102">
    <item>
   <p name="drxInactivityT">20</p>
   <p name="drxLongCycle">1280ms</p>
   <p name="drxOnDuratT">10</p>
   <p name="drxProfileIndex">102</p>
   <p name="drxProfilePriority">30</p>
   <p name="drxRetransT">16</p>
    </item>
   </list>
   <list name = "drxProfile103">
    <item>
   <p name="drxInactivityT">10</p>
   <p name="drxLongCycle">160ms</p>
   <p name="drxOnDuratT">6</p>
   <p name="drxProfileIndex">103</p>
   <p name="drxProfilePriority">30</p>
   <p name="drxRetransT">16</p>
    </item>
   </list>
   <list name = "qci1eVTTConfig">
    <item>
   <p name="qci1DlTargetBler">100</p>
   <p name="qci1HarqMaxTrDl">5</p>
   <p name="qci1HarqMaxTrUl">5</p>
   <p name="qci1ReconStopTimer">10</p>
   <p name="qci1ThroughputFactorDl">16</p>
   <p name="qci1ThroughputFactorUl">16</p>
   <p name="qci1UlTargetBler">10</p>
    </item>
   </list>
   <list name = "redBwRbUlConfig">
    <item>
   <p name="redBwMaxRbUl">50</p>
   <p name="redBwMinRbUl">1</p>
    </item>
   </list>
   <list name = "ulpcPucchConfig">
    <item>
   <p name="ulpcLowlevCch">-115</p>
   <p name="ulpcLowqualCch">-1</p>
   <p name="ulpcUplevCch">-113</p>
   <p name="ulpcUpqualCch">2</p>
    </item>
   </list>
   <list name = "ulpcPuschConfig">
    <item>
   <p name="ulpcLowlevSch">-100</p>
   <p name="ulpcLowqualSch">0</p>
   <p name="ulpcUplevSch">-88</p>
   <p name="ulpcUpqualSch">20</p>
    </item>
   </list>
   <p name="a1TimeToTriggerDeactInterMeas">480ms</p>
   <p name="a2RedirectQci1">enabled</p>
   <p name="a2TimeToTriggerActGERANMeas">320ms</p>
   <p name="a2TimeToTriggerActInterFreqMeas">480ms</p>
   <p name="a2TimeToTriggerActWcdmaMeas">320ms</p>
   <p name="a2TimeToTriggerRedirect">256ms</p>
   <p name="a3Offset">6</p>
   <p name="a3ReportInterval">640ms</p>
   <p name="a3TimeToTrigger">320ms</p>
   <p name="a3TriggerQuantityHO">RSRP</p>
   <p name="a5ReportInterval">640ms</p>
   <p name="a5TimeToTrigger">320ms</p>
   <p name="actAmle">false</p>
   <p name="actCsiRsSubFNonTM9Sch">false</p>
   <p name="actDataSessionProf">false</p>
   <p name="actDlsOldtc">false</p>
   <p name="actDlsVoicePacketAgg">true</p>
   <p name="actDownSampling">false</p>
   <p name="actDrx">true</p>
   <p name="actDrxDuringMeasGap">false</p>
   <p name="actDrxDuringTTIB">false</p>
   <p name="actEDrxIdle">false</p>
   <p name="actEicic">false</p>
   <p name="actFlexDupGap">false</p>
   <p name="actGsmSrvccMeasOpt">false</p>
   <p name="actInactiveTimeForwarding">false</p>
   <p name="actInterFreqLB">false</p>
   <p name="actIntraFreqLoadBal">false</p>
   <p name="actLBSpidUeSel">false</p>
   <p name="actLdPdcch">true</p>
   <p name="actMfbiDupFre">False</p>
   <p name="actMicroDtx">false</p>
   <p name="actModulationSchemeDl">64QAM</p>
   <p name="actModulationSchemeUL">16QAMHighMCS</p>
   <p name="actNbrForNonGbrBearers">false</p>
   <p name="actNoIntraBandIFMeasurements">false</p>
   <p name="actOlLaPdcch">true</p>
   <p name="actOtdoa">false</p>
   <p name="actProactSchedBySrb">true</p>
   <p name="actPrsTxDiv">false</p>
   <p name="actQci1RfDrx">false</p>
   <p name="actQci1eVTT">true</p>
   <p name="actSixIfMeasurements">false</p>
   <p name="actSmartDrx">false</p>
   <p name="actSrb1Robustness">false</p>
   <p name="actTtiBundling">true</p>
   <p name="actTtibAcqi">false</p>
   <p name="actTtibRsc">true</p>
   <p name="actUePowerBasedMobThr">false</p>
   <p name="actUlGrpHop">false</p>
   <p name="actUlLnkAdp">eUlLa</p>
   <p name="actUlpcMethod">PuschOLPucchOL</p>
   <p name="actUlpcRachPwrCtrl">false</p>
   <p name="actVoipCovBoost">false</p>
   <p name="addGbrTrafficRrHo">5</p>
   <p name="addGbrTrafficTcHo">10</p>
   <p name="addSpectrEmi">1</p>
   <p name="administrativeState">unlocked</p>
   <p name="allowPbIndexZero">false</p>
   <p name="amleMaxNumHo">10</p>
   <p name="amlePeriodLoadExchange">2000ms</p>
   <p name="cellIndOffServ">0dB</p>
   <p name="cellResourceSharingMode">none</p>
   <p name="cellTechnology">FDD</p>
   <p name="cellType">very_small</p>
   <p name="cqiAperEnable">true</p>
   <p name="cqiAperMode">FBT2</p>
   <p name="cqiPerSbCycK">1</p>
   <p name="cqiPerSimulAck">true</p>
   <p name="csgType">openAccess</p>
   <p name="dSrTransMax">64n</p>
   <p name="defPagCyc">64rf</p>
   <p name="deltaLbCioMargin">0</p>
   <p name="deltaPreMsg3">1</p>
   <p name="deltaTfEnabled">false</p>
   <p name="dlCellPwrRed">0</p>
   <p name="dlInterferenceEnable">false</p>
   <p name="dlInterferenceLevel">0</p>
   <p name="dlInterferenceModulation">QPSK</p>
   <p name="dlOlqcEnable">true</p>
   <p name="dlPathlossChg">3 db</p>
   <p name="dlPcfichBoost">0</p>
   <p name="dlPhichBoost">0</p>
   <p name="dlTargetBler">100</p>
   <p name="dlamcCqiDef">2</p>
   <p name="dlamcEnable">true</p>
   <p name="dlsDciCch">DCI format 1C</p>
   <p name="dlsOldtcTarget">98</p>
   <p name="dlsUsePartPrb">PRBs with PSS or SSS and PBCH used</p>
   <p name="eUlLaAtbPeriod">30</p>
   <p name="eUlLaBlerAveWin">30</p>
   <p name="eUlLaDeltaMcs">3</p>
   <p name="eUlLaLowMcsThr">1</p>
   <p name="eUlLaLowPrbThr">1</p>
   <p name="eUlLaPrbIncDecFactor">16</p>
   <p name="enableAmcPdcch">true</p>
   <p name="enableBetterCellHo">true</p>
   <p name="enableCovHo">true</p>
   <p name="enableLowAgg">true</p>
   <p name="enablePcPdcch">true</p>
   <p name="energySavingState">notEnergySaving</p>
   <p name="expectedCellSize">2.1km</p>
   <p name="filterCoeff">fc4</p>
   <p name="filterCoefficientCSFBCpichEcn0">fc4</p>
   <p name="filterCoefficientCSFBCpichRscp">fc4</p>
   <p name="filterCoefficientCpichEcn0">fc4</p>
   <p name="filterCoefficientCpichRscp">fc4</p>
   <p name="filterCoefficientRSRP">fc4</p>
   <p name="filterCoefficientRSRQ">fc4</p>
   <p name="filterCoefficientRSSI">fc2</p>
   <p name="gbrCongHandling">l2andl3</p>
   <p name="grpAssigPUSCH">0</p>
   <p name="harqMaxMsg3">5</p>
   <p name="harqMaxTrDl">5</p>
   <p name="harqMaxTrUlTtiBundling">8</p>
   <p name="harqMaxTxUl">5</p>
   <p name="hopModePusch">interSubFrame</p>
   <p name="hysA3Offset">0</p>
   <p name="hysThreshold2GERAN">0</p>
   <p name="hysThreshold2InterFreq">0</p>
   <p name="hysThreshold2Wcdma">0</p>
   <p name="hysThreshold2a">4</p>
   <p name="hysThreshold3">0</p>
   <p name="hysThreshold4">0</p>
   <p name="idleLBPercCaUe">0</p>
   <p name="idleLBPercentageOfUes">0</p>
   <p name="ilReacTimerUl">0</p>
   <p name="inactivityTimer">10</p>
   <p name="inactivityTimerQci5Sign">0</p>
   <p name="iniMcsDl">4</p>
   <p name="iniMcsUl">4</p>
   <p name="iniPrbsUl">10</p>
   <p name="lbLoadFilCoeff">4</p>
   <p name="lcrId">1</p>
   <p name="longPeriodScellChEst">2000</p>
   <p name="lowerMarginCio">-10</p>
   <p name="maxBitrateDl">170000</p>
   <p name="maxBitrateUl">75000</p>
   <p name="maxCrPgDl">12</p>
   <p name="maxCrRa4Dl">12</p>
   <p name="maxCrRaDl">12</p>
   <p name="maxCrRaDlHo">12</p>
   <p name="maxCrSibDl">26</p>
   <p name="maxGbrTrafficLimit">75</p>
   <p name="maxNumCaUeScell">400</p>
   <p name="maxNumScells">1</p>
   <p name="maxNumVoLteProactUlGrantsPerTti">0</p>
   <p name="maxPdcchAgg">16</p>
   <p name="maxPdcchAggHighLoad">16</p>
   <p name="maxPeriodicCqiIncrease">15</p>
   <p name="maxPhyCces">8</p>
   <p name="mbrSelector">ueCapability</p>
   <p name="measQuantityCSFBUtra">cpichRSCP</p>
   <p name="measQuantityUtra">cpichRSCP</p>
   <p name="minBitrateDl">30</p>
   <p name="minBitrateUl">30</p>
   <p name="mobStateParamNCelChgHgh">10</p>
   <p name="mobStateParamNCelChgMed">5</p>
   <p name="mobStateParamTEval">30s</p>
   <p name="mobStateParamTHystNorm">30s</p>
   <p name="msg3ConsecutiveDtx">0</p>
   <p name="msg3DtxOffset">0</p>
   <p name="multLoadMeasRrm">2</p>
   <p name="multLoadMeasRrmEicic">2</p>
   <p name="nbIoTMode">disabled</p>
   <p name="offsetFreqIntra">0dB</p>
   <p name="operationalState">disabled</p>
   <p name="p0NomPucch">-116</p>
   <p name="p0NomPusch">-80</p>
   <p name="pMax">240</p>
   <p name="pagingNb">quarterT</p>
   <p name="pdcchAggDefUe">4</p>
   <p name="pdcchAggMsg4">8</p>
   <p name="pdcchAggPaging">4</p>
   <p name="pdcchAggPreamb">4</p>
   <p name="pdcchAggRaresp">4</p>
   <p name="pdcchAggRarespHo">4</p>
   <p name="pdcchAggSib">4</p>
   <p name="pdcchAlpha">16</p>
   <p name="pdcchCqiShift">-50</p>
   <p name="pdcchHarqTargetBler">10</p>
   <p name="pdcchOrderConfig">enabled</p>
   <p name="pdcchUlDlBal">10</p>
   <p name="periodicCqiFeedbackType">wideband</p>
   <p name="phyCellId">242</p>
   <p name="preventCellActivation">false</p>
   <p name="prsPowerBoost">0</p>
   <p name="puschAckOffI">10</p>
   <p name="puschCqiOffI">8</p>
   <p name="puschRiOffI">10</p>
   <p name="raContResoT">64ms</p>
   <p name="raLargeMcsUl">5</p>
   <p name="raMsgPoffGrB">10 dB</p>
   <p name="raNondedPreamb">40</p>
   <p name="raPreGrASize">32</p>
   <p name="raRespWinSize">10</p>
   <p name="raSmallMcsUl">5</p>
   <p name="raSmallVolUl">144 bits</p>
   <p name="rcAmbrMgnDl">103</p>
   <p name="rcAmbrMgnUl">102</p>
   <p name="rcEnableDl">false</p>
   <p name="rcEnableUl">false</p>
   <p name="redBwEnDl">false</p>
   <p name="redBwMaxRbDl">50</p>
   <p name="redBwRpaEnUl">false</p>
   <p name="rttCsfbType">forbidden</p>
   <p name="sbHoMode">all</p>
   <p name="scellBadChQualThr">6</p>
   <p name="scellGoodChQualThr">8</p>
   <p name="scellNotDetectableThr">4</p>
   <p name="shortPeriodScellChEst">200</p>
   <p name="t302">8000</p>
   <p name="t304InterRAT">500ms</p>
   <p name="t304InterRATGsm">500ms</p>
   <p name="t304IntraLte">1000ms</p>
   <p name="t304eNaccGsm">4000ms</p>
   <p name="t320">180min</p>
   <p name="tExtendedWait">300</p>
   <p name="tLoadMeasX2">10000ms</p>
   <p name="tLoadMeasX2Eicic">2000</p>
   <p name="tPeriodicBsr">10ms</p>
   <p name="tPeriodicPhr">20sf</p>
   <p name="tPingPong">3000</p>
   <p name="tProhibitPhr">0sf</p>
   <p name="tReTxBsrTime">320 ms</p>
   <p name="tStoreUeCntxt">2000</p>
   <p name="taMaxOffset">52</p>
   <p name="taTimer">10240</p>
   <p name="taTimerMargin">2000</p>
   <p name="tac">153</p>
   <p name="targetSelMethod">SIB_based</p>
   <p name="threshold1">90</p>
   <p name="threshold2GERAN">23</p>
   <p name="threshold2InterFreq">40</p>
   <p name="threshold2Wcdma">25</p>
   <p name="threshold2a">44</p>
   <p name="threshold2aQci1">44</p>
   <p name="threshold3">24</p>
   <p name="threshold3a">26</p>
   <p name="threshold4">18</p>
   <p name="timeToTriggerSfHigh">0.5</p>
   <p name="timeToTriggerSfMedium">0.5</p>
   <p name="ttiBundlingBlerTarget">12</p>
   <p name="ttiBundlingBlerThreshold">15</p>
   <p name="ttiBundlingSinrThreshold">10</p>
   <p name="ttibAltUlPrbThreshold">0</p>
   <p name="ttibMinUlPrb">3</p>
   <p name="ttibOperMode">blerBasedNotWithMeasGap</p>
   <p name="ttibUlsMinTbs">144</p>
   <p name="ulRsCs">0</p>
   <p name="ulTargetBler">10</p>
   <p name="ulamcAllTbEn">false</p>
   <p name="ulamcSwitchPer">30</p>
   <p name="ulatbEventPer">1</p>
   <p name="ulpcAlpha">alpha 0.8</p>
   <p name="ulsFdPrbAssignAlg">MixedFD</p>
   <p name="ulsMaxPacketAgg">2</p>
   <p name="ulsMinRbPerUe">3</p>
   <p name="ulsMinTbs">104</p>
   <p name="ulsSchedMethod">channel unaware</p>
   <p name="upperMarginCio">10</p>
  </managedObject>
  <managedObject class="NOKLTE:LNCEL_FDD" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNCEL-0/LNCEL_FDD-0" operation="create">
   <list name = "mimoClConfig">
    <item>
   <p name="mimoClCqiThD">70</p>
   <p name="mimoClCqiThU">80</p>
   <p name="mimoClRiThD">28</p>
   <p name="mimoClRiThU">32</p>
    </item>
   </list>
   <list name = "mimoOlConfig">
    <item>
   <p name="mimoOlCqiThD">70</p>
   <p name="mimoOlCqiThU">80</p>
   <p name="mimoOlRiThD">28</p>
   <p name="mimoOlRiThU">32</p>
    </item>
   </list>
   <p name="act1TxIn2Tx">false</p>
   <p name="actAutoPucchAlloc">false</p>
   <p name="actCatM">false</p>
   <p name="actCombSuperCell">false</p>
   <p name="actDlSlimCarrier">false</p>
   <p name="actFastMimoSwitch">false</p>
   <p name="actLiquidCell">false</p>
   <p name="actMM		imo">false</p>
   <p name="actPdcchLoadGen">false</p>
   <p name="actPuschMask">false</p>
   <p name="actRedFreqOffset">false</p>
   <p name="actRepeaterMode">none</p>
   <p name="actSdlc">false</p>
   <p name="actShutdownTxPath">false</p>
   <p name="actSuperCell">false</p>
   <p name="actUciOnlyGrants">false</p>
   <p name="actUlMultiCluster">false</p>
   <p name="actUlPwrRestrScn">none</p>
   <p name="addNumDrbRadioReasHo">54</p>
   <p name="addNumDrbTimeCriticalHo">72</p>
   <p name="addNumQci1DrbRadioReasHo">18</p>
   <p name="addNumQci1DrbTimeCriticalHo">24</p>
   <p name="blankedPucch">0</p>
   <p name="dlChBw">10 MHz</p>
   <p name="dlMimoMode">Dynamic Open Loop MIMO</p>
   <p name="dlRsBoost">0dB</p>
   <p name="dlpcMimoComp">0</p>
   <p name="earfcnDL">2800</p>
   <p name="earfcnUL">20800</p>
   <p name="maxNrSymPdcch">3</p>
   <p name="maxNumActDrb">1296</p>
   <p name="maxNumActUE">100</p>
   <p name="maxNumCaConfUe">50</p>
   <p name="maxNumCaConfUe3c">20</p>
   <p name="maxNumCaConfUeDc">50</p>
   <p name="maxNumQci1Drb">352</p>
   <p name="maxNumUeDl">14</p>
   <p name="maxNumUeUl">14</p>
   <p name="prachCS">4</p>
   <p name="prachConfIndex">3</p>
   <p name="prachFreqOff">5</p>
   <p name="prachHsFlag">false</p>
   <p name="prsConfigurationIndex">151</p>
   <p name="prsNumDlFrames">1</p>
   <p name="pucchNAnCs">0</p>
   <p name="rfSensitivity">0</p>
   <p name="rootSeqIndex">0</p>
   <p name="selectOuterPuschRegion">None</p>
   <p name="srsActivation">false</p>
   <p name="srsPwrOffset">7</p>
   <p name="syncSigTxMode">TxDiv</p>
   <p name="ulChBw">10 MHz</p>
   <p name="ulCombinationMode">MRC</p>
   <p name="ulpcRarespTpc">2</p>
  </managedObject>
  <managedObject class="NOKLTE:LNHOIF" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNCEL-0/LNHOIF-1" operation="create">
   <p name="a3OffsetRsrpInterFreq">6</p>
   <p name="a3ReportIntervalRsrpInterFreq">240ms</p>
   <p name="a3TimeToTriggerRsrpInterFreq">320ms</p>
   <p name="a5ReportIntervalInterFreq">240ms</p>
   <p name="a5TimeToTriggerInterFreq">256ms</p>
   <p name="eutraCarrierInfo">3375</p>
   <p name="hysA3OffsetRsrpInterFreq">0</p>
   <p name="hysThreshold3InterFreq">0</p>
   <p name="interPresAntP">false</p>
   <p name="measQuantInterFreq">rsrp</p>
   <p name="measurementBandwidth">mbw25</p>
   <p name="offsetFreqInter">0dB</p>
   <p name="threshold3InterFreq">24</p>
   <p name="threshold3aInterFreq">26</p>
  </managedObject>
  <managedObject class="NOKLTE:LNMME" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNMME-0" operation="create">
   <list name = "accMmePlmnsList">
    <item>
   <p name="mcc">244</p>
   <p name="mnc">9</p>
   <p name="mncLength">2</p>
    </item>
    <item>
   <p name="mcc">244</p>
   <p name="mnc">40</p>
   <p name="mncLength">2</p>
    </item>
   </list>
   <list name = "rimReachablePLMNList">
    <item>
   <p name="mcc">244</p>
   <p name="mnc">9</p>
   <p name="mncLength">2</p>
    </item>
   </list>
   <p name="administrativeState">unlocked</p>
   <p name="ipAddrPrim">172.60.1.50</p>
   <p name="mmeName">ucMME</p>
   <p name="mmeRatSupport">Wideband-LTE</p>
   <p name="relMmeCap">1</p>
   <p name="s1LinkStatus">unavailable</p>
   <p name="transportNwId">0</p>
  </managedObject>
  <managedObject class="NOKLTE:LTAC" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/TAC-1/LTAC-1" operation="create">
   <p name="IncludeMbmsTraffic">true</p>
   <p name="tacLimitGbrEmergency">1000000</p>
   <p name="tacLimitGbrHandover">1000000</p>
   <p name="tacLimitGbrNormal">1000000</p>
   <p name="tacOverbookingLimit">0</p>
   <p name="transportNwId">0</p>
  </managedObject>
  <managedObject class="NOKLTE:LUAC" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/AMGR-1/LUAC-1" operation="create">
   <p name="actRestrictLoginToCnum">false</p>
   <p name="appAdmActivated">false</p>
   <p name="appAdmLocalPasswdLastChangedTime">0</p>
   <p name="btsAccountLockoutDuration">15</p>
   <p name="btsFailedLoginCountingPeriod">5</p>
   <p name="btsMaxFailedLoginAttempts">5</p>
   <p name="btsSessionLoginDelay">1</p>
   <p name="localPasswdExpiryPeriod">0</p>
   <p name="localPasswdLastChangedTime">0</p>
   <p name="localPasswdWarningPeriod">0</p>
   <p name="readOnlyActivated">false</p>
   <p name="readOnlyLocalPasswdLastChangedTime">0</p>
   <p name="secAdmActivated">false</p>
   <p name="secAdmLocalPasswdLastChangedTime">0</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:MNL" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1" operation="create">
   <p name="productVariantPlanned">fzmBTS</p>
   <p name="validatePlanAgainstDetectedHW">false</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:MNLENT" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1" operation="create">
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:MNL_R" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNL_R-1" operation="create">
   <p name="activeSWActivationTime">2018-03-23T12:03:01.424+05:30</p>
   <p name="activeSWReleaseVersion">FLF18SP_ENB_9999_180322_000344</p>
   <p name="configDataRevisionNumber">419</p>
  </managedObject>
  <managedObject class="NOKLTE:MPUCCH_FDD" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNCEL-0/LNCEL_FDD-0/MPUCCH_FDD-0" operation="create">
   <p name="addAUeRrHo">18</p>
   <p name="addAUeTcHo">24</p>
   <p name="addEmergencySessions">48</p>
   <p name="cellSrPeriod">20ms</p>
   <p name="cqiPerNp">40ms</p>
   <p name="deltaPucchShift">2</p>
   <p name="maxNumRrc">432</p>
   <p name="maxNumRrcEmergency">480</p>
   <p name="n1PucchAn">36</p>
   <p name="nCqiRb">4</p>
   <p name="nPucchF3Prbs">0</p>
   <p name="phichDur">Normal</p>
   <p name="phichRes">N = 1/6</p>
   <p name="riEnable">true</p>
   <p name="riPerM">1</p>
   <p name="riPerOffset">-1</p>
  </managedObject>
  <managedObject class="com.nokia.mrbts:MRBTS" version="MRBTS18_1711_002" distName="MRBTS-1" operation="create">
   <p name="altitude">0</p>
   <p name="btsName">FZM_GMC</p>
   <p name="latitude">0</p>
   <p name="longitude">0</p>
  </managedObject>
  <managedObject class="NOKLTE:NTPS" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/NTPS-1" operation="create">
   <list name = "ntpSyncServers">
   </list>
   <p name="ntpSyncLocalIpMapping">M-Plane</p>
   <p name="ntpSyncPollingRate">16</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:PMCADM" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/PMCADM-1" operation="create">
   <p name="reportingIntervalPm">60min</p>
   <p name="sdlMaxUploadFileNumber">1</p>
   <p name="sdlPrimaryDestIp">0.0.0.0</p>
  </managedObject>
  <managedObject class="NOKLTE:PMCCP" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/PMRNL-1/PMCCP-1" operation="create">
   <p name="cfgFastRRCConnSetReqThresh">1000</p>
   <p name="cfgPRBUtilThreshold">15</p>
   <p name="maxVBreakDistThresh">2</p>
  </managedObject>
  <managedObject class="NOKLTE:PMRNL" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/PMRNL-1" operation="create">
   <p name="mtActUsersLatencyPMQAPProfile">15min</p>
   <p name="mtBBPooling">15min</p>
   <p name="mtBroadcast">15min</p>
   <p name="mtCellAvailability">15min</p>
   <p name="mtCellLoad">15min</p>
   <p name="mtCellRes">15min</p>
   <p name="mtCellSpare">15min</p>
   <p name="mtCellThruput">15min</p>
   <p name="mtDataVolThrPMQAPProfile">15min</p>
   <p name="mtEPSBearer">15min</p>
   <p name="mtERABPMQAPProfile">15min</p>
   <p name="mtEutraFrequency">15min</p>
   <p name="mtFZAPGNSS">15min</p>
   <p name="mtGNSS">15min</p>
   <p name="mtHoRlf">15min</p>
   <p name="mtInterHomeeNBHo">15min</p>
   <p name="mtInterSiteX2">15min</p>
   <p name="mtInterSysHo">15min</p>
   <p name="mtInterSysHoEhrpdBc">15min</p>
   <p name="mtInterSysHoGsmNb">15min</p>
   <p name="mtInterSysHoUtranNb">15min</p>
   <p name="mtIntereNBHo">15min</p>
   <p name="mtIntereNBS1Ho">15min</p>
   <p name="mtIntereNBX2Ho">15min</p>
   <p name="mtIntraeNBHo">15min</p>
   <p name="mtIntraeNBHoExt">15min</p>
   <p name="mtLTEHo">15min</p>
   <p name="mtLTEM">15min</p>
   <p name="mtLTEMbmsLNBTS">15min</p>
   <p name="mtLTEMbmsMbsfnArea">15min</p>
   <p name="mtLTEMperLNBTS">15min</p>
   <p name="mtLaaResource">15min</p>
   <p name="mtLteULaaResources">disabled</p>
   <p name="mtM3SctpStatistics">15min</p>
   <p name="mtMAC">15min</p>
   <p name="mtMBMS">15min</p>
   <p name="mtMobilityEvents">15min</p>
   <p name="mtMroUtranFrequency">15min</p>
   <p name="mtNBIoT">15min</p>
   <p name="mtPIMCancellation">15min</p>
   <p name="mtPowQualDL">15min</p>
   <p name="mtPowQualUL">15min</p>
   <p name="mtQoS">15min</p>
   <p name="mtRIP">disabled</p>
   <p name="mtRLFPMQAPProfile">15min</p>
   <p name="mtRRC">15min</p>
   <p name="mtRSRPRSRQHist">15min</p>
   <p name="mtRadBearer">15min</p>
   <p name="mtRanSharing">15min</p>
   <p name="mtRibsSync">disabled</p>
   <p name="mtS1AP">15min</p>
   <p name="mtS1SctpStatistics">15min</p>
   <p name="mtSFN">15min</p>
   <p name="mtSIMonitoring">15min</p>
   <p name="mtSINR">15min</p>
   <p name="mtSfp">disabled</p>
   <p name="mtTranspLoad">15min</p>
   <p name="mtUEQuantity">15min</p>
   <p name="mtUEandServiceDiff">15min</p>
   <p name="mtUEstate">15min</p>
   <p name="mtVoLTEBLERHist">15min</p>
   <p name="mtVoLTEVoiceBPerioHist">15min</p>
   <p name="mtX2AP">15min</p>
   <p name="mtX2SctpStatistics">15min</p>
   <p name="mtX2gNB">15min</p>
   <p name="mteNBSpare">15min</p>
   <p name="mteNBload">15min</p>
   <p name="mtintraLTEHoNb">15min</p>
  </managedObject>
  <managedObject class="NOKLTE:PMTNL" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/PMTNL-1" operation="create">
   <p name="lte_ETHIF_Stats_Interval">15min</p>
   <p name="lte_Ethernet_Link_Interval">15min</p>
   <p name="lte_IPSec_Interval">15min</p>
   <p name="lte_IPSec_SA_Stats_interval">15min</p>
   <p name="lte_IP_Filtering_Interval">15min</p>
   <p name="lte_IP_Stats_Interval">15min</p>
   <p name="lte_IPv4_Stats_Interval">15min</p>
   <p name="lte_IPv6_Stats_Interval">15min</p>
   <p name="lte_NTP_Stats_interval">15min</p>
   <p name="lte_PHB_Stats_Interval">15min</p>
   <p name="lte_TAC_Statistics_Interval">15min</p>
   <p name="lte_TOP_FreqSync_Stats_Interval">15min</p>
   <p name="lte_TOP_PhaseSync_Stats_Interval">15min</p>
   <p name="lte_TWAMP_Stats_Interval">15min</p>
   <p name="lte_VLAN_IP_Stats_Interval">15min</p>
   <p name="lte_VLAN_IPv4_Stats_Interval">15min</p>
   <p name="lte_VLAN_IPv6_Stats_Interval">15min</p>
   <p name="lte_VLAN_PHB_Stats_Interval">15min</p>
   <p name="lte_VLAN_Stats_Interval">15min</p>
   <p name="lte_WiFi_Stats_Interval">15min</p>
   <p name="lte_loose_phase_and_time_sync_Stats_interval">15min</p>
   <p name="lte_port_network_access_security_Stats_interval">15min</p>
  </managedObject>
  <managedObject class="NOKLTE:PNASC" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/PNASC-1" operation="create">
   <p name="actPortNetAccCli">false</p>
   <p name="portNetAccCliEapGuardTimer">30</p>
   <p name="portNetAccDestGroup">PAE group address 01-80-C2-00-00-03</p>
  </managedObject>
  <managedObject class="NOKLTE:QOS" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/QOS-1" operation="create">
   <list name = "dscpMap">
    <item>
   <p name="dscp">56</p>
   <p name="pHB">expeditedForwarding</p>
   <p name="vlanPrio">7</p>
    </item>
    <item>
   <p name="dscp">46</p>
   <p name="pHB">expeditedForwarding</p>
   <p name="vlanPrio">6</p>
    </item>
    <item>
   <p name="dscp">38</p>
   <p name="pHB">assuredForwardingClass43</p>
   <p name="vlanPrio">5</p>
    </item>
    <item>
   <p name="dscp">36</p>
   <p name="pHB">assuredForwardingClass42</p>
   <p name="vlanPrio">5</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="pHB">assuredForwardingClass41</p>
   <p name="vlanPrio">5</p>
    </item>
    <item>
   <p name="dscp">30</p>
   <p name="pHB">assuredForwardingClass33</p>
   <p name="vlanPrio">4</p>
    </item>
    <item>
   <p name="dscp">28</p>
   <p name="pHB">assuredForwardingClass32</p>
   <p name="vlanPrio">4</p>
    </item>
    <item>
   <p name="dscp">26</p>
   <p name="pHB">assuredForwardingClass31</p>
   <p name="vlanPrio">4</p>
    </item>
    <item>
   <p name="dscp">22</p>
   <p name="pHB">assuredForwardingClass23</p>
   <p name="vlanPrio">3</p>
    </item>
    <item>
   <p name="dscp">20</p>
   <p name="pHB">assuredForwardingClass22</p>
   <p name="vlanPrio">3</p>
    </item>
    <item>
   <p name="dscp">18</p>
   <p name="pHB">assuredForwardingClass21</p>
   <p name="vlanPrio">3</p>
    </item>
    <item>
   <p name="dscp">14</p>
   <p name="pHB">assuredForwardingClass13</p>
   <p name="vlanPrio">1</p>
    </item>
    <item>
   <p name="dscp">12</p>
   <p name="pHB">assuredForwardingClass12</p>
   <p name="vlanPrio">1</p>
    </item>
    <item>
   <p name="dscp">10</p>
   <p name="pHB">assuredForwardingClass11</p>
   <p name="vlanPrio">1</p>
    </item>
    <item>
   <p name="dscp">0</p>
   <p name="pHB">bestEffort</p>
   <p name="vlanPrio">0</p>
    </item>
   </list>
   <list name = "perHopBehaviourWeightList">
    <item>
   <p name="assuredForwardingClass1">10</p>
   <p name="assuredForwardingClass2">100</p>
   <p name="assuredForwardingClass3">1000</p>
   <p name="assuredForwardingClass4">10000</p>
   <p name="bestEffort">1</p>
    </item>
   </list>
   <list name = "trafficTypesMap">
    <item>
   <p name="dscp">46</p>
   <p name="trafficType">CPLANE</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">MPLANE</p>
    </item>
    <item>
   <p name="dscp">46</p>
   <p name="trafficType">SPLANE</p>
    </item>
    <item>
   <p name="dscp">10</p>
   <p name="trafficType">ICMP</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">IKE</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD1</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD2</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD3</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD4</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD5</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD6</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD7</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD8</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD9</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD10</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD11</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD12</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD13</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD14</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD15</p>
    </item>
    <item>
   <p name="dscp">34</p>
   <p name="trafficType">BFD16</p>
    </item>
    <item>
   <p name="dscp">56</p>
   <p name="trafficType">HighPrioICMPv6</p>
    </item>
   </list>
   <p name="enablePhbCounters">false</p>
  </managedObject>
  <managedObject class="NOKLTE:REDRT" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNCEL-0/REDRT-0" operation="create">
   <p name="addGsmSIToRedirMsg">both</p>
   <p name="csFallBPrio">1</p>
   <p name="emerCallPrio">1</p>
   <p name="redirFreqUtra">10582</p>
   <p name="redirRat">utraFDD</p>
   <p name="redirectPrio">1</p>
  </managedObject>
  <managedObject class="NOKLTE:RIM" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/RIM-0" operation="create">
   <p name="goodRadioCondForCsfbThres">-90</p>
   <p name="maxSizeUtraSIGoodRadioCond">7168</p>
   <p name="maxSizeUtraSIWeakRadioCond">4096</p>
   <p name="nRimRirG">3</p>
   <p name="nRimRirU">3</p>
   <p name="tRimKaG">150</p>
   <p name="tRimKaU">150</p>
   <p name="tRimPollG">150</p>
   <p name="tRimPollU">150</p>
   <p name="tRimRirG">5</p>
   <p name="tRimRirU">5</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.eqm:RMOD" version="EQM18SP_FZM_1803_001" distName="MRBTS-1/EQM-1/APEQM-1/RMOD-1" operation="create">
   <p name="prodCodePlanned">FWHH_FRHH</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.eqmr:RMOD_R" version="EQMR18SP_FZM_1803_001" distName="MRBTS-1/EQM_R-1/APEQM_R-1/RMOD_R-1" operation="create">
   <p name="configDN">MRBTS-1/EQM-1/APEQM-1/RMOD-1</p>
   <p name="productCode">FWHH_FRHH</p>
   <p name="productName">FWHH_FRHH</p>
   <p name="serialNumber">EB164911618</p>
  </managedObject>
  <managedObject class="NOKLTE:RTPOL" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/IPNO-1/RTPOL-1" operation="create">
   <list name = "routingPolicies">
   </list>
  </managedObject>
  <managedObject class="NOKLTE:SCTP" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/SCTP-1" operation="create">
   <p name="assocMaxRetrans">7</p>
   <p name="maxTimeSctpSetup">100</p>
   <p name="pathMaxRetrans">5</p>
   <p name="rtoMax">30000</p>
   <p name="rtoMin">1000</p>
   <p name="sctpHeartbeatInterval">20</p>
   <p name="sctpMaxPayloadSize">1395</p>
   <p name="sctpSackDelay">200</p>
   <p name="sctpSackFreq">2</p>
   <p name="sctpWaitTimeInitSeqRetryM3">240</p>
   <p name="sctpWaitTimeInitSeqRetryS1">240</p>
   <p name="sctpWaitTimeInitSeqRetryX2">240</p>
  </managedObject>
  <managedObject class="NOKLTE:SDRX" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNCEL-0/SDRX-0" operation="create">
   <list name = "drxSmartProfile2">
    <item>
   <p name="drxInactivityT">4</p>
   <p name="drxLongCycle">40ms</p>
   <p name="drxOnDuratT">6</p>
   <p name="drxProfileIndex">2</p>
   <p name="drxProfilePriority">40</p>
   <p name="drxRetransT">4</p>
   <p name="drxShortCycle">40ms</p>
   <p name="drxShortCycleT">1</p>
   <p name="smartStInactFactor">124</p>
    </item>
   </list>
   <list name = "drxSmartProfile3">
    <item>
   <p name="drxInactivityT">10</p>
   <p name="drxLongCycle">80ms</p>
   <p name="drxOnDuratT">6</p>
   <p name="drxProfileIndex">3</p>
   <p name="drxProfilePriority">30</p>
   <p name="drxRetransT">4</p>
   <p name="drxShortCycle">40ms</p>
   <p name="drxShortCycleT">5</p>
   <p name="smartStInactFactor">6</p>
    </item>
   </list>
   <list name = "drxSmartProfile4">
    <item>
   <p name="drxInactivityT">50</p>
   <p name="drxLongCycle">320ms</p>
   <p name="drxOnDuratT">10</p>
   <p name="drxProfileIndex">4</p>
   <p name="drxProfilePriority">20</p>
   <p name="drxRetransT">16</p>
   <p name="drxShortCycle">80ms</p>
   <p name="drxShortCycleT">4</p>
   <p name="smartStInactFactor">4</p>
    </item>
   </list>
   <list name = "drxSmartProfile5">
    <item>
   <p name="drxInactivityT">50</p>
   <p name="drxLongCycle">2560ms</p>
   <p name="drxOnDuratT">10</p>
   <p name="drxProfileIndex">5</p>
   <p name="drxProfilePriority">10</p>
   <p name="drxRetransT">16</p>
   <p name="drxShortCycle">80ms</p>
   <p name="drxShortCycleT">16</p>
   <p name="smartStInactFactor">32</p>
    </item>
   </list>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:SECADM" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/SECADM-1" operation="create">
   <p name="actServiceAccountSsh">false</p>
   <p name="actServicePortState">false</p>
   <p name="ethernetPortSecurityEnabled">true</p>
   <p name="idleSessionTimeWebUI">30</p>
   <p name="keyContinuation">false</p>
   <p name="passwdHistoryLength">0</p>
   <p name="serviceAccountSshStatus">disabled</p>
   <p name="servicePortStatus">disabled</p>
   <p name="sshClientAliveTimer">45</p>
   <p name="sshSessionLoginDelayTimer">10</p>
  </managedObject>
  <managedObject class="NOKLTE:SECPRM" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/SECPRM-1" operation="create">
   <p name="idleSessionTimeTwp">45</p>
  </managedObject>
  <managedObject class="NOKLTE:SIB" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/LNCEL-0/SIB-0" operation="create">
   <list name = "celResTiF">
    <item>
   <p name="celResTiFHM">0.5</p>
   <p name="celResTiFMM">0.5</p>
    </item>
   </list>
   <list name = "sib2Scheduling">
    <item>
   <p name="siMessagePeriodicity">160ms</p>
   <p name="siMessageRepetition">1</p>
   <p name="siMessageSibType">SIB2</p>
    </item>
   </list>
   <list name = "sib3Scheduling">
    <item>
   <p name="siMessagePeriodicity">320ms</p>
   <p name="siMessageRepetition">1</p>
   <p name="siMessageSibType">SIB3</p>
    </item>
   </list>
   <list name = "spStResPars">
    <item>
   <p name="nCellChgHigh">10</p>
   <p name="nCellChgMed">5</p>
   <p name="qHystSfHigh">-4 dB</p>
   <p name="qHystSfMed">-4 dB</p>
   <p name="tEvaluation">30s</p>
   <p name="tHystNormal">30s</p>
    </item>
   </list>
   <p name="acBarSkipForMMTELVideo">false</p>
   <p name="acBarSkipForMMTELVoice">false</p>
   <p name="acBarSkipForSMS">false</p>
   <p name="acbProfileId">0</p>
   <p name="autoAcBarringStartTimer">300</p>
   <p name="autoAcbPlmnRmvlStopTimer">3600</p>
   <p name="autoPlmnRmvlStartTimer">300</p>
   <p name="cellBarred">notBarred</p>
   <p name="cellReSelPrio">7</p>
   <p name="intrFrqCelRes">allowed</p>
   <p name="intraPresAntP">false</p>
   <p name="modPeriodCoeff">8</p>
   <p name="n310">n10</p>
   <p name="n311">n1</p>
   <p name="pMaxOwnCell">23</p>
   <p name="prachPwrRamp">2dB</p>
   <p name="preambTxMax">10</p>
   <p name="primPlmnCellres">Not Reserved</p>
   <p name="qHyst">2dB</p>
   <p name="qrxlevmin">-130</p>
   <p name="qrxlevminintraF">-130</p>
   <p name="sIntrasearch">62</p>
   <p name="sNonIntrsearch">16</p>
   <p name="siWindowLen">10ms</p>
   <p name="sib2xTransmit">false</p>
   <p name="t300">400ms</p>
   <p name="t301">400ms</p>
   <p name="t310">2000ms</p>
   <p name="t311">3000ms</p>
   <p name="tReselEutr">1</p>
   <p name="threshSrvLow">4</p>
   <p name="ulpcIniPrePwr">-96 dBm</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.eqm:SMOD" version="EQM18SP_FZM_1803_001" distName="MRBTS-1/EQM-1/APEQM-1/CABINET-1/SMOD-1" operation="create">
   <p name="ledControl">30</p>
   <p name="moduleLocation">AH</p>
   <p name="operatingMode">Normal</p>
   <p name="prodCodePlanned">472946A</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.eqmr:SMOD_R" version="EQMR18SP_FZM_1803_001" distName="MRBTS-1/EQM_R-1/APEQM_R-1/CABINET_R-1/SMOD_R-1" operation="create">
   <p name="configDN">MRBTS-1/EQM-1/APEQM-1/CABINET-1/SMOD-1</p>
   <p name="productCode">472946A</p>
   <p name="productName">FWHH</p>
   <p name="serialNumber">EB164911618</p>
  </managedObject>
  <managedObject class="NOKLTE:STPG" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/SYNC-1/STPG-1" operation="create">
   <list name = "synchroSourceList">
    <item>
   <p name="clockProtocol">clkToP</p>
   <p name="priority">1</p>
   <p name="ssmEnabled">true</p>
   <p name="ssmTimeout">5</p>
   <p name="unitNumber">1</p>
    </item>
   </list>
   <p name="compClkSync">false</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:SYNC" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/SYNC-1" operation="create">
   <p name="btsSyncMode">FreqSync</p>
   <p name="manualFrameTimingAdjustment">0</p>
  </managedObject>
  <managedObject class="NOKLTE:SYNC" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/SYNC-1" operation="create">
   <p name="actHybridSynch">false</p>
  </managedObject>
  <managedObject class="NOKLTE:TAC" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/TAC-1" operation="create">
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:TIME" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/SYNC-1/TIME-1" operation="create">
   <p name="timeZone">(GMT+5.5) Asia/Calcutta</p>
  </managedObject>
  <managedObject class="NOKLTE:TOPB" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/TOPB-1" operation="create">
   <p name="actTopBoundaryClock">false</p>
  </managedObject>
  <managedObject class="NOKLTE:TOPF" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/TOPB-1/TOPF-1" operation="create">
   <list name = "acceptedClockQuality">
   <p>13</p>
   <p>14</p>
   <p>6</p>
   <p>7</p>
   </list>
   <list name = "topMasters">
    <item>
   <p name="lockState">OutOfLocked</p>
   <p name="masterClockState">MASTER_CLOCK_UNKNOWN</p>
   <p name="masterIpAddr">10.58.130.251</p>
   <p name="priority_1">0</p>
   <p name="priority_2">0</p>
   <p name="receivedClockQuality">0</p>
   <p name="receivedPriority_1">0</p>
   <p name="receivedPriority_2">0</p>
   <p name="topMasterActive">inactive</p>
    </item>
   </list>
   <p name="actTopFreqSynch">false</p>
   <p name="announceRequestMode">ANNOUNCE_ALL</p>
   <p name="ieeeTelecomProfile">IEEE1588</p>
   <p name="logMeanSyncValue">-4</p>
   <p name="topDomainNumber">0</p>
   <p name="waitingTime">237</p>
  </managedObject>
  <managedObject class="NOKLTE:TOPP" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/TOPB-1/TOPP-1" operation="create">
   <list name = "acceptedClockQuality">
   <p>6</p>
   <p>7</p>
   <p>135</p>
   </list>
   <list name = "topMasters">
    <item>
   <p name="masterIpAddr">10.58.130.251</p>
    </item>
   </list>
   <p name="actTopPhaseSynch">true</p>
   <p name="lockState">OutOfLocked</p>
   <p name="logMeanSyncValue">-7</p>
   <p name="phaseErrorComp">0</p>
   <p name="topComMode">IP Unicast</p>
   <p name="topDomainNumber">0</p>
   <p name="topEthMulticastAddress">01-80-C2-00-00-0E</p>
   <p name="tuningProfile">Normal</p>
  </managedObject>
  <managedObject class="com.nokia.srbts.mnl:TRBLCADM" version="MNL18SP_FZM_1803_001" distName="MRBTS-1/MNL-1/MNLENT-1/TRBLCADM-1" operation="create">
   <p name="actCoreDumpFileCollection">true</p>
   <p name="actRemoteSyslogTransmission">false</p>
  </managedObject>
  <managedObject class="NOKLTE:UNIT" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/UNIT-1" operation="create">
   <p name="unitTypeActual">472946A</p>
   <p name="unitTypeExpected">472946A</p>
  </managedObject>
  <managedObject class="NOKLTE:WIFICNF" version="FLF18SP_1803_01_1802_02" distName="MRBTS-1/LNBTS-1/FTM-1/WIFICNF-1" operation="create">
   <list name = "l2VlanIdList">
   </list>
   <p name="acceptableFrameTypes">ADMIT_ALL</p>
   <p name="actWiFi">false</p>
   <p name="portDefaultPriority">0</p>
  </managedObject></cmdata>`