package src

import (
	"log"
	"net"
	"strings"
)

var (
	hasMax bool
	hasSPF bool
	hasDMARC bool
	spfRecord string
	dmarcRecord string
)

type CheckDomainMapType = map[string]interface{}

func CheckDomain(domain string) CheckDomainMapType{	
	precedureMxRecords(domain)
	procedureTxtRecords(domain)
	procedureDmarcRecords(domain)

	return makeVerifyMap()
}

func makeVerifyMap() CheckDomainMapType{
	m := make(CheckDomainMapType)
	m["hasMax"] = hasMax
	m["hasSPF"] = hasSPF
	m["hasDMARC"] = hasDMARC
	m["spfRecord"] = spfRecord
	m["dmarcRecord"] = dmarcRecord

	return m
}

func precedureMxRecords(domain string){
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Println("Error : ", err)
	}

	if len(mxRecords) > 0 {
		hasMax = true
	}
}

func procedureTxtRecords(domain string) {
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Println("Error : ", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break	
		}
	}
}

func procedureDmarcRecords(domain string) {

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Println("Error : ", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
}