package main

import (
	"F-Scrack-Go/serverScan/icmpcheck"
)

func main() {
	icmpcheck.ICMPRun([]string{"127.0.0.1", "192.168.2.1"})
}
