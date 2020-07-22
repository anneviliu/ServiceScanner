package main

import (
	"bufio"
	"github.com/malfunkt/iprange"
	"log"
	"os"
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

func CheckIp(ip string) {

}

func StandardHostsViaFile(filePath string) []string {
	var standardHostList []string
	hostList, err := ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range hostList {
		ipRng, err := iprange.ParseList(v)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range ipRng.Expand() {
			standardHostList = append(standardHostList, v.String())
		}
	}
	return standardHostList
	//fmt.Println(standardHostList)
}
