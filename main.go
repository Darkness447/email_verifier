package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// bufio is for input and

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain,hasMX,hasSPF,sprRecord,hasDMARC")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecor, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		panic(err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		panic(err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecor = record
			break
		}
	}

}
