package main

import (
	"serviceScan/serverScan/icmpcheck"
	"serviceScan/serverScan/portscan"
	"serviceScan/serverScan/vscan"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var hostList []string
var aliveList []string
var aliveHosts []string
var aliveAddr []string
var TargetBanners []string
var isIP bool

var green = color.New(color.FgGreen)
var red = color.New(color.FgRed).Add(color.Bold)
var blue = color.New(color.FgBlue).Add(color.Bold)

func main() {

	if len(os.Args) == 1 {
		flag.Usage()
	}

	if printVersion {
		fmt.Printf("Port and Service Scanner.\nVersion:%s\nBy:Annevi\n", version)
		os.Exit(0)
	}

	// 避免同时使用 -hf 和 -h
	if hostinfile != "" && hosts != "" {
		color.Red("Can not use -hf and -h at the same time.")
		flag.Usage()
		os.Exit(1)
	}

	// -hf
	if hostinfile != "" {
		hostList = StandardIPViaFile(hostinfile, "file")
		//hostList = StandardIPViaFile("test.txt", "file")
		aliveList = icmpcheck.ICMPRun(hostList)
		for _, host := range aliveList {
			green.Printf("[+] [ICMP] Target '%s' is alive\n", host)
		}
		blue.Println("\nProcess: ")
		aliveHosts, aliveAddr = portscan.TCPportScan(aliveList, ports, "tcp", timeout)
		fmt.Println(" ")

		if service != "" {
			if len(aliveAddr) > 0 {
				TargetBanners = vscan.GetProbes(aliveAddr)
			}
		} else {
			for _, host := range aliveAddr {
				green.Printf("[+] [TCP] %s is open.\n", host)
			}
		}
	}

	// -h
	if hosts != "" {
		hostList = StandardIPViaFile(hosts, "single")
		// icmp 存活探测
		aliveList = icmpcheck.ICMPRun(hostList)
		for _, host := range aliveList {
			green.Printf("[+] [ICMP] Target '%s' is alive\n", host)
		}
		blue.Println("\nProcess: ")
		aliveHosts, aliveAddr = portscan.TCPportScan(aliveList, ports, "tcp", timeout)
		fmt.Println(" ")

		if service != "" {
			if len(aliveAddr) > 0 {
				TargetBanners = vscan.GetProbes(aliveAddr)
			}
		} else {
			for _, host := range aliveAddr {
				green.Printf("[+] [TCP] %s is open.\n", host)
			}
		}
	}

	// -o result output
	if outFile != "" && len(aliveAddr) != 0 {
		if PathCheck(outFile) == 1 {
			f, _ := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
			for _, host := range aliveAddr {
				f.WriteString(host + "\n")
			}
			green.Printf("[+] Output the scanning information in %s\n", outFile)
			defer f.Close()
		} else if PathCheck(outFile) == -1 {
			red.Println("[-] OutFile path error ")
			os.Exit(1)
		} else if PathCheck(outFile) == -2 {
			red.Printf("[-] OutFile File %s already exits\n", outFile)
			os.Exit(1)
		} else if PathCheck(outFile) == -3 {
			red.Println("[-] OutFile File create failed!")
			os.Exit(1)
		} else {
			red.Println("[-] Unknown error")
			os.Exit(1)
		}
	}
}
