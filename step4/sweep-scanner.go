package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func scanPort(host string, port, timeout int, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, time.Duration(timeout)*time.Millisecond)
	if err == nil {
		fmt.Printf("Host: %s Port: %d is open\n", host, port)
		conn.Close()
	}
}

func main() {
	hosts := flag.String("host", "localhost", "Comma-separated list of hosts or CIDR notation")
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
			scanPort(h, *port, *timeout, &wg)
		}(host)
	}

	wg.Wait()
}
