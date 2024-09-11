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
	fmt.Println("Enter a domain to check its email security settings:")

	if scanner.Scan() {
		domain := scanner.Text()
		CheckEmailSecurityRecords(domain)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

func CheckEmailSecurityRecords(domain string) {
	hasMX, mxErr := CheckMXRecords(domain)
	hasSPF, spfRecord := CheckSPFRecords(domain)
	hasDMARC, dmarcRecord := CheckDMARCRecords(domain)

	if mxErr != nil {
		log.Printf("Error checking MX records: %v\n", mxErr)
	}

	fmt.Printf("Domain: %s\n", domain)
	fmt.Printf("Has MX records: %t\n", hasMX)
	fmt.Printf("Has SPF records: %t\nSPF Record: %s\n", hasSPF, spfRecord)
	fmt.Printf("Has DMARC records: %t\nDMARC Record: %s\n", hasDMARC, dmarcRecord)
}

func CheckMXRecords(domain string) (bool, error) {
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		return false, err
	}

	return len(mxRecords) > 0, nil
}

func CheckSPFRecords(domain string) (bool, string) {
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error checking SPF records: %v\n", err)
		return false, ""
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			return true, record
		}
	}

	return false, ""
}

func CheckDMARCRecords(domain string) (bool, string) {
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error checking DMARC records: %v\n", err)
		return false, ""
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			return true, record
		}
	}

	return false, ""
}
