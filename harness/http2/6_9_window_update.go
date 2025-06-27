package http2

import (
	"log"
	"net"

	"golang.org/x/net/http2"
)

// Test Case 6.9/1: Sends a WINDOW_UPDATE frame with a flow-control window increment of 0.
// The client is expected to detect a PROTOCOL_ERROR.
func RunTest6_9_1(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.9/1...")

	if err := framer.WriteWindowUpdate(0, 0); err != nil {
		log.Printf("Failed to write WINDOW_UPDATE frame: %v", err)
		return
	}

	log.Println("Sent WINDOW_UPDATE with 0 increment. Test complete.")
}
