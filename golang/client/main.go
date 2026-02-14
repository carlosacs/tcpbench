package main

import (
	"fmt"
	"net"

	"tcptest/common"
)

func main() {

	len := 1920 * 1080 * 4 * 60

	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send some data to the server
	buf := make([]byte, common.BUFF_SIZE)
	for len > 0 {
		w, err := conn.Write(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		len -= w
	}

	// Close the connection
	conn.Close()
}
