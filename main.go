package main

import (
	"flag"
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
	cidr := flag.String("cidr", "0.0.0.0/0", "CIDR to check")
	ip := flag.String("ip", "127.0.0.1/32", "IP to check")
	flag.Parse()
	fmt.Println(cidrContains(*cidr, *ip))
	// fmt.Println(cidrContains("10.0.0.0/8", "10.10.10.25/32"))
	// fmt.Println(cidrContains("10.0.0.0/8", "172.16.32.54/32"))

}
