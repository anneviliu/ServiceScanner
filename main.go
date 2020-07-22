package main

import (
	"F-Scrack-Go/serverScan/icmpcheck"
	"fmt"
)

var hostList []string
var DEBUG bool = true

func main() {
	if hostinfile != "" && !DEBUG {
		hostList = StandardHostsViaFile(hostinfile)
	} else if DEBUG {
		hostList = StandardHostsViaFile("test.txt")
	}
	fmt.Println(icmpcheck.ICMPRun(hostList))
}
