package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord")

	for scanner.Scan() {
		checkDoamin(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from Input %v \n", err)
	}
}

func checkDoamin(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, demarcRecord string

	mxRecord, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}
	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	demarcRecords, err := net.LookupTXT("_demarc." + domain)
	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	for _, record := range demarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			demarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, demarcRecord)

}
