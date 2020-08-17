package main

import (
	"F-Scrack-Go/serverScan/icmpcheck"
	"F-Scrack-Go/serverScan/portscan"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var hostList []string
var aliveList []string
var aliveAddr []string
var isIP bool

func main() {
	// 避免同时使用 -hf 和 -h
	if hostinfile != "" && hosts != "" {
		color.Red("Can not use -hf and -h at the same time.")
		os.Exit(1)
	}

	// -hf
	c := color.New(color.FgGreen).Add(color.Bold)

	if hostinfile != "" {
		//hostList = StandardIPViaFile(hostinfile, "file")
		hostList = StandardIPViaFile("test.txt", "file")
		aliveList = icmpcheck.ICMPRun(hostList)
		for _, host := range aliveList {
			c.Printf("[ICMP] Target '%s' is alive\n", host)
		}
		portscan.TCPportScan(aliveList, ports, "tcp", timeout)

	}

	// -h
	if hosts != "" {
		// 标准化ip
		hostList = StandardIPViaFile(hosts, "single")
		if model == "" {
			// icmp 存活探测
			aliveList = icmpcheck.ICMPRun(hostList)
			for _, host := range aliveList {
				fmt.Printf("(ICMP) Target '%s' is alive\n", host)
			}
			portscan.TCPportScan(aliveList, ports, "tcp", timeout)

		} else if model == "tcp" {
			a, _ := portscan.TCPportScan(aliveList, "80", "tcp", 2)
			fmt.Println(a)
		}
	}

	// -o result output
	if outFile != "" {

	}
}
