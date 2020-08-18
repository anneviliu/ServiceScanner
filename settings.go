package main

import "flag"

var version = "v0.0.1"
var printVersion bool
var hosts = ""
var ports = ""
var hostinfile = ""
var outFile = ""
var timeout int
var service = ""
var DEBUG bool = true

func init() {
	flag.BoolVar(&printVersion, "v", false, "ServerScan Build Version")
	flag.StringVar(&hosts, "h", "", "Host to be scanned, supports four formats:\n192.168.1.1\n192.168.1.1-10\n192.168.1.*\n192.168.1.0/24.")
	flag.StringVar(&hostinfile, "hf", "", "Hosts list to be scanned, one line one host/ip, support formats are as same as -h")
	flag.StringVar(&ports, "p", "80-99,7000-9000,9001-9999,4430,1433,1521,3306,5000,5432,6379,21,22,100-500,873,4440,6082,3389,5560,5900-5909,1080,1900,10809,50030,50050,50070", "Customize port list, separate with ',' example: 21,22,80-99,8000-8080 ...")
	//flag.StringVar(&ports, "p", "10-1000", "Customize port list, separate with ',' example: 21,22,80-99,8000-8080 ...")
	flag.StringVar(&service, "s", "", "Use Probes to get Service,usage: -s 1")
	flag.IntVar(&timeout, "t", 2, "Setting scanner connection timeouts,Maxtime 30 Second.")
	flag.StringVar(&outFile, "o", "", "Output the scanning information to file.")
	flag.Parse()
}
