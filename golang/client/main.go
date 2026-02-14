package main

import (
	"fmt"
	"net"
	"os"

	"tcptest/common"
)

func main() {

	dataLen := 200 * common.KB // 1920 * 1080 * 4 * 60

	address := "localhost:8080"
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	// Connect to the server
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send some data to the server

	buf := make([]byte, min(common.BUFF_SIZE, dataLen))
	for dataLen > 0 {
		w, err := conn.Write(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		dataLen -= w
	}

	// Close the connection
	conn.Close()
}
