package main

import (
	fmt
	net
)
func checkErr(err) {
	if err != nil {
 		fmt.Println(err)
 		os.Exit(1)
}

func cidrContains (cidr, ip string) {
	ipAddress, ipNet, err := net.ParseCIDR(cidr)
	checkErr(err)
	return ipNet.Contains(net.ParseCIDR(ip))
}

fmt.Println(cidrContains("10.0.0.0/8","10.10.10.255")
fmt.Println(cidrContains("10.0.0.0/8","172.16.32.54")
