package main

import (
	"fmt"

	"github.com/gosnmp/gosnmp"
)

func main() {
	target := "127.0.0.1"
	community := "community"
	oid := ".1.3.6.1.2.1.25.2.3.1.3" // disk

	gosnmp.Default.Target = target
	gosnmp.Default.Community = community
	gosnmp.Default.Version = gosnmp.Version2c

	err := gosnmp.Default.Connect()
	if err != nil {
		fmt.Println("Connect error", err)
		return
	}

	defer gosnmp.Default.Conn.Close()

	err = gosnmp.Default.Walk(oid, func(pdu gosnmp.SnmpPDU) error {
		fmt.Printf("%s = %s\n", pdu.Name, pdu.Value)
		return nil
	})
	if err != nil {
		fmt.Println("Walk error:", err)
	}

}
