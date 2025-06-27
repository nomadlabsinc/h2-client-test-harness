package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"

	"golang.org/x/net/http2"
)

func main() {
	testCase := flag.String("test", "", "The test case to run")
	flag.Parse()

	if *testCase == "" {
		fmt.Println("Usage: go run . --test=<test_case_id>")
		fmt.Println("Example: go run . --test=6.5/1")
		os.Exit(1)
	}

	// Ensure certificates exist, or generate them.
	if err := ensureCerts(); err != nil {
		log.Fatalf("Failed to create or find certificates: %v", err)
	}

	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("Failed to load certificates: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{http2.NextProtoTLS},
	}

	listener, err := tls.Listen("tcp", "127.0.0.1:8080", tlsConfig)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	log.Printf("Test harness server listening on %s for test case '%s'", listener.Addr().String(), *testCase)

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalf("Failed to accept connection: %v", err)
	}
	
	handleConnection(conn, *testCase)
}

func handleConnection(conn net.Conn, testCase string) {
	defer conn.Close()
	log.Printf("Accepted connection from %s", conn.RemoteAddr())

	// Before running the test, we must complete the HTTP/2 handshake.
	// 1. Read the client preface.
	preface := make([]byte, len(http2.ClientPreface))
	if _, err := conn.Read(preface); err != nil {
		log.Printf("Failed to read client preface: %v", err)
		return
	}
	if string(preface) != http2.ClientPreface {
		log.Printf("Incorrect client preface received: %s", string(preface))
		return
	}
	log.Println("Client preface received.")

	// 2. Read the client's initial SETTINGS frame.
	framer := http2.NewFramer(conn, conn)
	frame, err := framer.ReadFrame()
	if err != nil {
		log.Printf("Failed to read client's initial SETTINGS frame: %v", err)
		return
	}
	if _, ok := frame.(*http2.SettingsFrame); !ok {
		log.Printf("Expected a SETTINGS frame from client, but got %T", frame)
		return
	}
	log.Println("Client's initial SETTINGS frame received.")

	// 3. Send our own initial SETTINGS frame (it can be empty).
	if err := framer.WriteSettings(); err != nil {
		log.Printf("Failed to write initial server SETTINGS frame: %v", err)
		return
	}
	log.Println("Initial server SETTINGS frame sent.")

	// Handshake complete. Now run the specific test case logic.
	switch testCase {
	case "6.5/1":
		runTest6_5_1(conn, framer)
	default:
		log.Printf("Unknown or unimplemented test case: %s", testCase)
	}
}

// Test Case 6.5/1: Sends a SETTINGS frame with ACK flag and a non-empty payload.
// The client is expected to detect a FRAME_SIZE_ERROR.
func runTest6_5_1(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.5/1...")

	// The h2spec test sends a 1-byte payload with an ACK SETTINGS frame.
	// A valid ACK SETTINGS frame must have a zero-length payload.
	// We must write this as a raw frame, as the library prevents this.
	
	// Frame Header: Length (1), Type (SETTINGS), Flags (ACK), StreamID (0)
	// Payload: One arbitrary byte (e.g., 0xFF)
	malformedFrame := []byte{
		0x00, 0x00, 0x01, // Length: 1
		0x04,             // Type: SETTINGS (0x4)
		0x01,             // Flags: ACK (0x1)
		0x00, 0x00, 0x00, 0x00, // Stream ID: 0
		0xFF, // The illegal payload byte
	}

	if _, err := conn.Write(malformedFrame); err != nil {
		log.Printf("Failed to write malformed SETTINGS ACK frame: %v", err)
		return
	}

	log.Println("Sent malformed SETTINGS ACK frame. Test complete.")
}

// ensureCerts checks for cert.pem and key.pem and generates them if they don't exist.
func ensureCerts() error {
	if _, err := os.Stat("cert.pem"); os.IsNotExist(err) {
		log.Println("Certificate 'cert.pem' not found, generating new one...")
		cmd := exec.Command("openssl", "req", "-x509", "-newkey", "rsa:2048", "-nodes", "-keyout", "key.pem", "-out", "cert.pem", "-days", "365", "-subj", "/CN=localhost")
		cmd.Dir = "." // Run in the test harness directory
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to generate certificate: %s\n%s", err, string(out))
		}
		log.Println("Successfully generated cert.pem and key.pem.")
	}
	return nil
}