package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
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
	hosts := flag.String("host", "localhost", "Comma-separated list of hosts")
	port := flag.Int("port", 80, "Port to scan")
	timeout := flag.Int("timeout", 500, "Connection timeout in milliseconds")
	concurrent := flag.Int("concurrent", 100, "Number of concurrent scans")
	flag.Parse()

	hostList := strings.Split(*hosts, ",")

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, *concurrent)

	for _, host := range hostList {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(h string) {
			defer func() { <-semaphore }()
			portScanner(h, *port, *timeout, &wg)
		}(host)
	}
}
