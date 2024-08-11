package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func portScanner(host string, port, timeout int, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Millisecond)
	if err == nil {
		fmt.Printf("Port: %d is open\n", port)
		conn.Close()
	}
}

func main() {
	host := flag.String("host", "localhost", "Host to scan")
	timeout := flag.Int("timeout", 500, "Connection timeout in milliseconds")
	concurrent := flag.Int("concurrent", 100, "Number of concurrent scans")
	flag.Parse()

	fmt.Printf("Scanning host: %s\n", *host)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, *concurrent)

	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(p int) {
			defer func() { <-semaphore }()
			portScanner(*host, p, *timeout, &wg)
		}(port)
	}
}
