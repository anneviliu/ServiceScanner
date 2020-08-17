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
var aliveHosts []string
var aliveAddr []string
var isIP bool

var green = color.New(color.FgGreen).Add(color.Bold)
var red = color.New(color.FgRed).Add(color.Bold)

func main() {
	// 避免同时使用 -hf 和 -h
	if hostinfile != "" && hosts != "" {
		color.Red("Can not use -hf and -h at the same time.")
		os.Exit(1)
	}

	// -hf

	if hostinfile != "" {
		//hostList = StandardIPViaFile(hostinfile, "file")
		hostList = StandardIPViaFile("test.txt", "file")
		aliveList = icmpcheck.ICMPRun(hostList)
		for _, host := range aliveList {
			green.Printf("[+] [ICMP] Target '%s' is alive\n", host)
		}

		//r := make([]string, len(hostList))
		//copy(r, hostList)
		//for k, v := range r {
		//	for _, alive := range aliveList {
		//		if v == alive {
		//			i := k
		//			r = append(r[:i], r[i+1:]...)
		//			break
		//		}
		//		//fmt.Println(hostList)
		//	}
		//}
		//fmt.Println(r)
		aliveHosts, aliveAddr = portscan.TCPportScan(aliveList, ports, "tcp", timeout)
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
			_, _ = portscan.TCPportScan(aliveList, "80", "tcp", 2)
		}
	}

	// -o result output
	if outFile != "" && len(aliveList) != 0 {
		if PathCheck(outFile) == 1 {
			f, _ := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
			for _, host := range aliveAddr {
				f.WriteString(host + "\n")
			}
			green.Printf("[+] Output the scanning information in %s\n", outFile)
			defer f.Close()
		} else if PathCheck(outFile) == -1 {
			red.Println("[-] File path error \n")
			os.Exit(1)
		} else if PathCheck(outFile) == -2 {
			red.Printf("[-] File %s already exits\n", outFile)
			os.Exit(1)
		} else if PathCheck(outFile) == -3 {
			red.Println("[-] File create failed!\n")
			os.Exit(1)
		} else {
			red.Println("[-] Unknown error\n")
			os.Exit(1)
		}
	}
}
