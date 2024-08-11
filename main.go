package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ccscan -host=<hostname> -port=<port>")
		return
	}

	host := os.Args[1][6:] // "-host=" prefix is 6 chars
	port := os.Args[2][6:] // "-port=" prefix is 6 chars

	address := net.JoinHostPort(host, port)
	fmt.Printf("Scanning host: %s port: %s\n", host, port)

	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		fmt.Printf("Port: %s is closed or unreachable\n", port)
		return
	}
	conn.Close()
	fmt.Printf("Port: %s is open\n", port)
}
