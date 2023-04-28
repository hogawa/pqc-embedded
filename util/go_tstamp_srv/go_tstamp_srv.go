package main

// Author: William Suzuki (william.suzuki@alumni.usp.br)
// This is a first try.
// The drawback with this implementation is that it does not support multiple
// connections.

// Clock1 is a TCP server that periodically writes the time.
// package main
// =>
// <=
// !=

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	strUrl := "0.0.0.0:5000"
	fmt.Printf("Running on %v\n", strUrl)
	listener, err := net.Listen("tcp", strUrl)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		handleConn(conn) // handle one connection at a time
	}

}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("2006-02-01 15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(100 * time.Millisecond)
	}
}
