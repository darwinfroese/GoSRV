package main

import (
	"net"
	"fmt"
	"strings"
	"os"
	"bufio"
)

func process(msg []byte, len int, addr *net.UDPAddr) {
	return
}

func connectionHandler() {
	serverAddr, _ := net.ResolveUDPAddr("udp", ":8081")

	serverConn, _ := net.ListenUDP("udp", serverAddr)
	defer serverConn.Close()

	buf := make([]byte, 1024)

	for {
		msgLen, addr, _ := serverConn.ReadFromUDP(buf)

		process(buf, msgLen, addr)
	}
}

func main() {
	fmt.Println("Launching server...")

	go connectionHandler()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Server > ")
		msg,_ := reader.ReadString('\n')

		if (strings.Index(msg, "exit") == 0) {
			return
		}
	}
}