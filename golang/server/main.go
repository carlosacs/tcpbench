package main

import (
	"fmt"
	"net"
	"time"

	"tcptest/common"

	humanize "github.com/dustin/go-humanize"
)

func main() {
	// Listen for incoming connections on port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Accept incoming connections and handle them
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	begin := time.Now()

	// Close the connection when we're done
	defer conn.Close()

	// Read incoming data
	buf := make([]byte, common.BUFF_SIZE)
	len := 0
	for {
		r, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		len += r
	}

	duration := time.Since(begin)
	hlen := humanize.Bytes(uint64(len))
	bps := uint64(float64(len) / duration.Seconds())
	hps := humanize.Bytes(bps)

	// Print the incoming data
	fmt.Printf("Received: %d (%s) in %s: %d/s (%s/s)", len, hlen, duration, bps, hps)
}
