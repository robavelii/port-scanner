package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func portScanner(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)
	if err == nil {
		fmt.Printf("Port: %d is open\n", port)
		conn.Close()
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ccscan -host=<hostname>")
		return
	}

	host := os.Args[1][6:] // "-host=" prefix is 6 chars

	fmt.Printf("Scanning host: %s\n", host)

	for port := 1; port <= 65535; port++ {
		portScanner(host, port)
	}
}
