package main

import (
	"bufio"
	"flag"
	"github.com/malfunkt/iprange"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var line []string

func ReadFile(textfile string) ([]string, error) {
	file, err := os.Open(textfile)
	if err != nil {
		log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	//fmt.Println(line)
	if err := scanner.Err(); err != nil {
		log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
		return nil, err
	}

	return line, err
}

func CheckHosts(hosts string) bool {
	hostsPattern := `^(([01]?\d?\d|2[0-4]\d|25[0-5])\.){3}([01]?\d?\d|2[0-4]\d|25[0-5])\/(\d{1}|[0-2]{1}\d{1}|3[0-2])$|^(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[0-9]{1,2})(\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[0-9]{1,2})){3}$`
	hostsRegexp := regexp.MustCompile(hostsPattern)
	checkHost := hostsRegexp.MatchString(hosts)

	hostsPattern2 := `\b(?:(?:25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.){3}(((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})\-((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2}))\b`
	hostsRegexp2 := regexp.MustCompile(hostsPattern2)
	checkHost2 := hostsRegexp2.MatchString(hosts)

	hostsPattern3 := `((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.){3}(\*$)`
	hostsRegexp3 := regexp.MustCompile(hostsPattern3)
	checkHost3 := hostsRegexp3.MatchString(hosts)

	if hosts == "" || (checkHost == false && checkHost2 == false && checkHost3 == false) {
		flag.Usage()
		return false
	}
	return true
}

func CheckPorts(ports string) bool {
	portsPattern := `^([0-9]|[1-9]\d|[1-9]\d{2}|[1-9]\d{3}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$|^\d+(-\d+)?(,\d+(-\d+)?)*$`
	portsRegexp := regexp.MustCompile(portsPattern)
	checkPort := portsRegexp.MatchString(ports)
	if ports != "" && checkPort == false {
		flag.Usage()
		return false
	}
	return true
}

/*
 * @param  host string 传入需要处理的ip  t string 传入IP的方式 （文件\命令行输入）
 * @return []string 返回IP列表
 * @auth Annevi
 */

func StandardIPViaFile(host string, t string) []string {
	var standardHostList []string
	if t == "file" {
		hostList, err := ReadFile(host)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range hostList {
			if CheckHosts(v) {
				ipRng, err := iprange.ParseList(v)
				if err != nil {
					log.Fatal(err)
				}
				for _, v := range ipRng.Expand() {
					standardHostList = append(standardHostList, v.String())
				}
			} else {
				log.Fatal("Hosts format error")
			}
		}
	} else if t == "single" {
		if CheckHosts(host) {
			ipRng, err := iprange.ParseList(host)
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range ipRng.Expand() {
				standardHostList = append(standardHostList, v.String())
			}
		} else {
			log.Fatal("Hosts format error")
		}
	}
	return standardHostList
	//fmt.Println(standardHostList)
}

/*
-1 文件路径不正确
-2 文件已存在
-3 文件创建失败
-4 未知错误
1 正常



*/
func PathCheck(files string) int {
	path, _ := filepath.Split(files)
	_, err := os.Stat(path)
	if err == nil {
		_, err2 := os.Stat(files)
		if err2 == nil {
			return -1
		}
		if os.IsNotExist(err2) {
			return 1
		} else {
			return -2
		}
	} else {
		err3 := os.MkdirAll(path, os.ModePerm)
		if err3 == nil {
			return 1
		} else {
			return -3
		}
	}
}
