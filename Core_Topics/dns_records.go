package main

import (
	"fmt"
	"net"
)

func main() {
	domain := "example.com"

	// 1. A Records (IPv4)
	fmt.Println("🔹 A Records (IPv4):")
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, ip := range ips {
			if ip.To4() != nil {
				fmt.Println(" -", ip.String())
			}
		}
	}

	// 2. AAAA Records (IPv6)
	fmt.Println("\n🔹 AAAA Records (IPv6):")
	for _, ip := range ips {
		if ip.To16() != nil && ip.To4() == nil {
			fmt.Println(" -", ip.String())
		}
	}

	// 3. MX Records
	fmt.Println("\n🔹 MX Records:")
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, mx := range mxRecords {
			fmt.Printf(" - %s (Priority: %d)\n", mx.Host, mx.Pref)
		}
	}

	// 4. CNAME Record
	fmt.Println("\n🔹 CNAME Record:")
	cname, err := net.LookupCNAME(domain)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(" -", cname)
	}

	// 5. TXT Records
	fmt.Println("\n🔹 TXT Records:")
	txts, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, txt := range txts {
			fmt.Println(" -", txt)
		}
	}
}
