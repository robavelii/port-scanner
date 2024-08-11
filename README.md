# Go Port Scanner CLI Tool

This project is a command-line tool built in Go to scan network ports. The tool supports various types of scans, from scanning a single port on a single host to performing more advanced sweep scans across multiple hosts and ports.

## Table of Contents

- [Introduction](#introduction)
- [Project Structure](#project-structure)
- [Step-by-Step Implementation](#step-by-step-implementation)
  - [Step 1: Single Port Scan](#step-1-single-port-scan)
  - [Step 2: Vanilla Scan (Multiple Ports)](#step-2-vanilla-scan-multiple-ports)
  - [Step 3: Concurrent Scanning](#step-3-concurrent-scanning)
  - [Step 4: Sweep Scan](#step-4-sweep-scan)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This port scanner is designed to perform a series of network scans to detect open ports on hosts. It provides various scanning techniques, starting from basic single-port scanning to more complex sweep and SYN scans. 

## Project Structure

The project is organized into different steps, each focusing on a specific aspect of port scanning. Each step is implemented in its own file or directory.

```
port-scanner/
│
├── step1/
│   └── basic-scanner.go        # Single port scan
│
├── step2/
│   └── vanilla-scanner.go      # Vanilla scan across multiple ports
│
├── step3/
│   └── optimized-scanner.go    # Concurrent scanning for speed
│
├── step4/
│   └── sweep-scanner.go    # Sweep scan across multiple hosts
│
│
└── README.md            # Project documentation
```

## Step-by-Step Implementation

### Step 1: Single Port Scan

In this step, the tool accepts a host and a port as command-line arguments and attempts to open a TCP connection to the specified port. If the connection is successful, the port is reported as open.

```go
go run step1/basic-scanner.go -host=localhost -port=5000
```

### Step 2: Vanilla Scan (Multiple Ports)

This step expands the functionality to scan multiple ports on a single host. The tool attempts to open a full TCP connection on each port.

```go
go run step2/vanilla-scanner.go -host=localhost
```

### Step 3: Concurrent Scanning

To improve the speed of scanning, this step introduces concurrency. The tool scans multiple ports in parallel using goroutines, significantly reducing the time required to complete the scan.

```go
go run step3/optimized-scanner.go -host=localhost -concurrent=100
```

### Step 4: Sweep Scan

This step introduces the ability to scan a single port across multiple hosts, known as a sweep scan. The tool accepts a list of hosts or a wildcard IP address and scans the specified port on all these hosts.

```go
go run step4/sweep-scanner.go -host="localhost,192.168.1.1,192.168.1.10,google.com"
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
```