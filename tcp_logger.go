// A simple tcp logger that just prints
// out what is coming in through a port
package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// Listen on port 2000 on local system
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Accept as a goroutine so it can accept concurrent connections
		go func(c net.Conn) {
			// Do stuff with this connection
			// Here we are just printing out the data stream
			io.Copy(c, c)
			// Shut down the connection
			c.Close()
		}(conn)
	}
}
