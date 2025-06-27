package cases

import (
	"bytes"
	"log"
	"net"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
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

// Test Case 6.9/2: Sends a WINDOW_UPDATE frame with a flow-control window increment of 0 on a stream.
// The client is expected to detect a PROTOCOL_ERROR.
func RunTest6_9_2(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.9/2...")

	// To test a stream-specific error, we first need to create a stream.
	// We can do this by sending a HEADERS frame.
	var streamID uint32 = 1
	var buf bytes.Buffer
	hpackEncoder := hpack.NewEncoder(&buf)
	hpackEncoder.WriteField(hpack.HeaderField{Name: ":status", Value: "200"})
	
	if err := framer.WriteHeaders(http2.HeadersFrameParam{
		StreamID:      streamID,
		BlockFragment: buf.Bytes(),
		EndStream:     false,
		EndHeaders:    true,
	}); err != nil {
		log.Printf("Failed to write HEADERS frame: %v", err)
		return
	}
	log.Println("Sent HEADERS frame to create stream 1.")

	if err := framer.WriteWindowUpdate(streamID, 0); err != nil {
		log.Printf("Failed to write WINDOW_UPDATE frame: %v", err)
		return
	}

	log.Println("Sent WINDOW_UPDATE with 0 increment on stream 1. Test complete.")
}
