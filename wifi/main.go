package main

import (
	"fmt"
	"net"
	"os/exec"
	"strings"

	"github.com/mdlayher/arp"
)

func main() {
	// Show current SSID and MAC Address
	ifi, err := net.InterfaceByName("en0")
	if err != nil {
		panic(err)
	}
	c, err := arp.Dial(ifi)
	if err != nil {
		panic(err)
	}

	ip := net.IPv4(192, 168, 10, 1)
	addr, err := c.Resolve(ip)
	if err != nil {
		panic(err)
	}
	fmt.Println("SSID: " + wifiName())
	fmt.Println("MAC: " + addr.String())

	// Ask
	var isExit string
	fmt.Print("Change SSID and Password? (y/N) > ")
	fmt.Scanln(&isExit)
	if isExit == "n" || isExit == "" {
		return
	}

	// Set new SSID and Password
	var ssid string
	var pw string
	fmt.Print("New SSID > ")
	fmt.Scanln(&ssid)
	fmt.Print("New Password > ")
	fmt.Scanln(&pw)

	if ssid == "" {
		fmt.Println("SSID must be more than 1 character.")
		return
	}

	conn, _ := net.Dial("udp", "192.168.10.1:8889")
	defer conn.Close()
	sendCommand(conn, "command")
	sendCommand(conn, "wifi"+" "+ssid+" "+pw)
}

func wifiName() string {
	out, err := exec.Command("sh", "-c", "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport -I | awk -F: '/ SSID/{print $2}'").Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

func sendCommand(conn net.Conn, command string) {
	conn.Write([]byte(command))
}
