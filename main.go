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
	case "6.5/2":
		runTest6_5_2(conn, framer)
	case "6.5/3":
		runTest6_5_3(conn, framer)
	case "6.5.2/1":
		runTest6_5_2_1(conn, framer)
	case "6.5.2/2":
		runTest6_5_2_2(conn, framer)
	case "6.5.2/3":
		runTest6_5_2_3(conn, framer)
	case "6.5.2/4":
		runTest6_5_2_4(conn, framer)
	case "6.5.2/5":
		runTest6_5_2_5(conn, framer)
	case "6.5.3/2":
		runTest6_5_3_2(conn, framer)
	case "6.7/1":
		runTest6_7_1(conn, framer)
	case "6.7/2":
		runTest6_7_2(conn, framer)
	case "6.7/3":
		runTest6_7_3(conn, framer)
	case "6.7/4":
		runTest6_7_4(conn, framer)
	case "6.8/1":
		runTest6_8_1(conn, framer)
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

// Test Case 6.5/2: Sends a SETTINGS frame with a stream identifier other than 0x0.
// The client is expected to detect a PROTOCOL_ERROR.
func runTest6_5_2(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.5/2...")

	// A valid SETTINGS frame MUST have a stream identifier of 0.
	// We will send an empty SETTINGS frame but set the stream ID to 1.
	
	// Frame Header: Length (0), Type (SETTINGS), Flags (0), StreamID (1)
	malformedFrame := []byte{
		0x00, 0x00, 0x00, // Length: 0
		0x04,             // Type: SETTINGS (0x4)
		0x00,             // Flags: 0
		0x00, 0x00, 0x00, 0x01, // Stream ID: 1
	}

	if _, err := conn.Write(malformedFrame); err != nil {
		log.Printf("Failed to write malformed SETTINGS frame: %v", err)
		return
	}

	log.Println("Sent malformed SETTINGS frame with non-zero stream ID. Test complete.")
}

// Test Case 6.5/3: Sends a SETTINGS frame with a length other than a multiple of 6 octets.
// The client is expected to detect a FRAME_SIZE_ERROR.
func runTest6_5_3(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.5/3...")

	// A valid SETTINGS frame's payload must be a multiple of 6 bytes long.
	// We will send a SETTINGS frame with a 5-byte payload.
	
	// Frame Header: Length (5), Type (SETTINGS), Flags (0), StreamID (0)
	malformedFrame := []byte{
		0x00, 0x00, 0x05, // Length: 5
		0x04,             // Type: SETTINGS (0x4)
		0x00,             // Flags: 0
		0x00, 0x00, 0x00, 0x00, // Stream ID: 0
		0x01, 0x02, 0x03, 0x04, 0x05, // 5-byte payload
	}

	if _, err := conn.Write(malformedFrame); err != nil {
		log.Printf("Failed to write malformed SETTINGS frame: %v", err)
		return
	}

	log.Println("Sent malformed SETTINGS frame with invalid length. Test complete.")
}

// Test Case 6.5.2/1: Sends SETTINGS_ENABLE_PUSH with a value other than 0 or 1.
// The client is expected to detect a PROTOCOL_ERROR.
func runTest6_5_2_1(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.5.2/1...")

	if err := framer.WriteSettings(http2.Setting{ID: http2.SettingEnablePush, Val: 2}); err != nil {
		log.Printf("Failed to write SETTINGS frame: %v", err)
		return
	}

	log.Println("Sent SETTINGS_ENABLE_PUSH with invalid value. Test complete.")
}

// Test Case 6.5.2/2: Sends SETTINGS_INITIAL_WINDOW_SIZE with a value > 2^31-1.
// The client is expected to detect a FLOW_CONTROL_ERROR.
func runTest6_5_2_2(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.5.2/2...")

	if err := framer.WriteSettings(http2.Setting{ID: http2.SettingInitialWindowSize, Val: 2147483648}); err != nil {
		log.Printf("Failed to write SETTINGS frame: %v", err)
		return
	}

	log.Println("Sent SETTINGS_INITIAL_WINDOW_SIZE with invalid value. Test complete.")
}

// Test Case 6.5.2/3: Sends SETTINGS_MAX_FRAME_SIZE with a value < 16384.
// The client is expected to detect a PROTOCOL_ERROR.
func runTest6_5_2_3(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.5.2/3...")

	if err := framer.WriteSettings(http2.Setting{ID: http2.SettingMaxFrameSize, Val: 16383}); err != nil {
		log.Printf("Failed to write SETTINGS frame: %v", err)
		return
	}

	log.Println("Sent SETTINGS_MAX_FRAME_SIZE with invalid value. Test complete.")
}

// Test Case 6.5.2/4: Sends SETTINGS_MAX_FRAME_SIZE with a value > 16777215.
// The client is expected to detect a PROTOCOL_ERROR.
func runTest6_5_2_4(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.5.2/4...")

	if err := framer.WriteSettings(http2.Setting{ID: http2.SettingMaxFrameSize, Val: 16777216}); err != nil {
		log.Printf("Failed to write SETTINGS frame: %v", err)
		return
	}

	log.Println("Sent SETTINGS_MAX_FRAME_SIZE with invalid value. Test complete.")
}

// Test Case 6.5.2/5: Sends a SETTINGS frame with an unknown identifier.
// The client is expected to ignore the setting and not terminate the connection.
func runTest6_5_2_5(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.5.2/5...")

	// Send a setting with an unknown ID. The client should ignore this.
	if err := framer.WriteSettings(http2.Setting{ID: 0xFF, Val: 1}); err != nil {
		log.Printf("Failed to write SETTINGS frame with unknown ID: %v", err)
		return
	}
	log.Println("Sent SETTINGS frame with unknown ID.")

	// To verify the connection is still alive, we send a PING...
	pingData := [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
	if err := framer.WritePing(false, pingData); err != nil {
		log.Printf("Failed to write PING frame: %v", err)
		return
	}
	log.Println("Sent PING frame, awaiting ACK.")

	// ...and expect a PING ACK in response.
	// We will loop, ignoring other frames, until we get the PING ACK or an error.
	for {
		frame, err := framer.ReadFrame()
		if err != nil {
			log.Printf("Failed to read frame after PING: %v", err)
			return
		}

		switch f := frame.(type) {
		case *http2.PingFrame:
			if !f.IsAck() {
				log.Println("Received a PING frame, but it was not an ACK.")
				return
			}
			if string(f.Data[:]) != string(pingData[:]) {
				log.Printf("Received PING ACK, but data does not match. Got %v", f.Data)
				return
			}
			log.Println("Received PING ACK with correct data. Test complete.")
			return // Success
		default:
			// Ignore other frames like WINDOW_UPDATE, etc.
			log.Printf("Ignoring frame of type %T while waiting for PING ACK.", f)
		}
	}
}

// Test Case 6.5.3/2: Sends a SETTINGS frame and expects an ACK.
// The client is expected to immediately send a SETTINGS frame with the ACK flag.
func runTest6_5_3_2(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.5.3/2...")

	// Send a valid SETTINGS frame.
	if err := framer.WriteSettings(http2.Setting{ID: http2.SettingEnablePush, Val: 0}); err != nil {
		log.Printf("Failed to write SETTINGS frame: %v", err)
		return
	}
	log.Println("Sent SETTINGS frame, awaiting ACK.")

	// Expect a SETTINGS ACK in response.
	for {
		frame, err := framer.ReadFrame()
		if err != nil {
			log.Printf("Failed to read frame while waiting for SETTINGS ACK: %v", err)
			return
		}

		switch f := frame.(type) {
		case *http2.SettingsFrame:
			if !f.IsAck() {
				log.Println("Received a SETTINGS frame, but it was not an ACK.")
				return
			}
			log.Println("Received SETTINGS ACK. Test complete.")
			return // Success
		default:
			// Ignore other frames.
			log.Printf("Ignoring frame of type %T while waiting for SETTINGS ACK.", f)
		}
	}
}

// Test Case 6.7/1: Sends a PING frame.
// The client is expected to respond with a PING frame with the ACK flag.
func runTest6_7_1(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.7/1...")

	pingData := [8]byte{'h', '2', 's', 'p', 'e', 'c'}
	if err := framer.WritePing(false, pingData); err != nil {
		log.Printf("Failed to write PING frame: %v", err)
		return
	}
	log.Println("Sent PING frame, awaiting ACK.")

	for {
		frame, err := framer.ReadFrame()
		if err != nil {
			log.Printf("Failed to read frame while waiting for PING ACK: %v", err)
			return
		}

		switch f := frame.(type) {
		case *http2.PingFrame:
			if !f.IsAck() {
				log.Println("Received a PING frame, but it was not an ACK.")
				return
			}
			if string(f.Data[:]) != string(pingData[:]) {
				log.Printf("Received PING ACK, but data does not match. Got %v", f.Data)
				return
			}
			log.Println("Received PING ACK with correct data. Test complete.")
			return // Success
		default:
			log.Printf("Ignoring frame of type %T while waiting for PING ACK.", f)
		}
	}
}

// Test Case 6.7/2: Sends a PING frame with ACK flag.
// The client is expected to not respond to the PING ACK, but respond to a subsequent PING.
func runTest6_7_2(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.7/2...")

	// Send a PING with ACK, which the client should ignore.
	if err := framer.WritePing(true, [8]byte{'i', 'g', 'n', 'o', 'r', 'e'}); err != nil {
		log.Printf("Failed to write PING ACK frame: %v", err)
		return
	}
	log.Println("Sent PING ACK, which should be ignored.")

	// Send a normal PING, which the client should respond to.
	pingData := [8]byte{'r', 'e', 's', 'p', 'o', 'n', 'd'}
	if err := framer.WritePing(false, pingData); err != nil {
		log.Printf("Failed to write subsequent PING frame: %v", err)
		return
	}
	log.Println("Sent second PING frame, awaiting ACK.")

	for {
		frame, err := framer.ReadFrame()
		if err != nil {
			log.Printf("Failed to read frame while waiting for PING ACK: %v", err)
			return
		}

		switch f := frame.(type) {
		case *http2.PingFrame:
			if !f.IsAck() {
				log.Println("Received a PING frame, but it was not an ACK.")
				return
			}
			if string(f.Data[:]) != string(pingData[:]) {
				log.Printf("Received PING ACK, but data does not match expected for the second PING. Got %v", f.Data)
				return
			}
			log.Println("Received PING ACK for the second PING. Test complete.")
			return // Success
		default:
			log.Printf("Ignoring frame of type %T while waiting for PING ACK.", f)
		}
	}
}

// Test Case 6.7/3: Sends a PING frame with a non-zero stream identifier.
// The client is expected to detect a PROTOCOL_ERROR.
func runTest6_7_3(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.7/3...")

	// Frame Header: Length (8), Type (PING), Flags (0), StreamID (1)
	malformedFrame := []byte{
		0x00, 0x00, 0x08, // Length: 8
		0x06,             // Type: PING (0x6)
		0x00,             // Flags: 0
		0x00, 0x00, 0x00, 0x01, // Stream ID: 1
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Payload
	}

	if _, err := conn.Write(malformedFrame); err != nil {
		log.Printf("Failed to write malformed PING frame: %v", err)
		return
	}

	log.Println("Sent malformed PING frame with non-zero stream ID. Test complete.")
}

// Test Case 6.7/4: Sends a PING frame with a length other than 8.
// The client is expected to detect a FRAME_SIZE_ERROR.
func runTest6_7_4(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.7/4...")

	// Frame Header: Length (6), Type (PING), Flags (0), StreamID (0)
	malformedFrame := []byte{
		0x00, 0x00, 0x06, // Length: 6
		0x06,             // Type: PING (0x6)
		0x00,             // Flags: 0
		0x00, 0x00, 0x00, 0x00, // Stream ID: 0
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Payload
	}

	if _, err := conn.Write(malformedFrame); err != nil {
		log.Printf("Failed to write malformed PING frame: %v", err)
		return
	}

	log.Println("Sent malformed PING frame with invalid length. Test complete.")
}

// Test Case 6.8/1: Sends a GOAWAY frame with a non-zero stream identifier.
// The client is expected to detect a PROTOCOL_ERROR.
func runTest6_8_1(conn net.Conn, framer *http2.Framer) {
	log.Println("Running test case 6.8/1...")

	// Frame Header: Length (8), Type (GOAWAY), Flags (0), StreamID (1)
	malformedFrame := []byte{
		0x00, 0x00, 0x08, // Length: 8
		0x07,             // Type: GOAWAY (0x7)
		0x00,             // Flags: 0
		0x00, 0x00, 0x00, 0x01, // Stream ID: 1
		0x00, 0x00, 0x00, 0x00, // Last-Stream-ID
		0x00, 0x00, 0x00, 0x00, // Error Code
	}

	if _, err := conn.Write(malformedFrame); err != nil {
		log.Printf("Failed to write malformed GOAWAY frame: %v", err)
		return
	}

	log.Println("Sent malformed GOAWAY frame with non-zero stream ID. Test complete.")
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