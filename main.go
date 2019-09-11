package main

import (
	"fmt"
	"net"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func cidrContains(cidr, ip string) bool {
	_, ipNet, err := net.ParseCIDR(cidr)
	checkErr(err)
	checkIP, _, err := net.ParseCIDR(ip)
	checkErr(err)
	return ipNet.Contains(checkIP)
}

func main() {

	fmt.Println(cidrContains("10.0.0.0/8", "10.10.10.25/32"))
	fmt.Println(cidrContains("10.0.0.0/8", "172.16.32.54/32"))

}
