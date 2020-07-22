package main

import "flag"

var version = "v0.0.1"
var printVersion bool

var hosts = ""
var ports = ""
var model = ""
var hostinfile = ""
var outFile = ""
var timeout int

func init() {
	flag.BoolVar(&printVersion, "v", false, "ServerScan Build Version")
	flag.StringVar(&hosts, "h", "", "Host to be scanned, supports four formats:\n192.168.1.1\n192.168.1.1-10\n192.168.1.*\n192.168.1.0/24.")
	flag.StringVar(&hostinfile, "hf", "", "Hosts list to be scanned, one line one host/ip, support formats are as same as -h")
	flag.StringVar(&ports, "p", "80-99,7000-9000,9001-9999,4430,1433,1521,3306,5000,5432,6379,21,22,100-500,873,4440,6082,3389,5560,5900-5909,1080,1900,10809,50030,50050,50070", "Customize port list, separate with ',' example: 21,22,80-99,8000-8080 ...")
	flag.StringVar(&model, "m", "icmp", "Scan Model [ssh,mysql,ftp,mssql,redis,telnet,postgresql,memcached,mongodb,elasticsearch,icmp,service\ndefault all of them")
	flag.IntVar(&timeout, "t", 2, "Setting scaner connection timeouts,Maxtime 30 Second.")
	flag.StringVar(&outFile, "o", "", "Output the scanning information to file.")
	flag.Parse()
}
