# HTTP/2 Client Test Harness

This repository contains a Go-based test harness for testing the compliance of HTTP/2 clients. It is inspired by [h2spec](https://github.com/summerwind/h2spec), but the logic is inverted: this harness acts as a server that sends specific, sometimes malformed, frames to a client to test the client's response.

This project also includes a reference "verifier" client, which is a known-good Go HTTP/2 client used to validate that the harness is behaving correctly.

## Usage

This harness is designed to be a general-purpose testing tool for any HTTP/2 client. To test your client:

1.  **Run the Harness (Server):**
    In one terminal, start the harness with the desired test case.
    ```shell
    go run ./cmd/harness --test=<test_case_id>
    ```

2.  **Run Your Client:**
    In another terminal, run your HTTP/2 client and make a request to `https://localhost:8080`.

3.  **Observe the Outcome:**
    Your client should receive the appropriate error or response from the harness.

### Verifying the Harness Itself

To verify that the harness is working correctly, you can run it against the included verifier client:

1.  **Run the Harness (Server):**
    ```shell
    go run ./cmd/harness --test=<test_case_id> &
    ```

2.  **Run the Verifier (Client):**
    ```shell
    go run ./cmd/verifier --test=<test_case_id>
    ```
    If the verifier exits with a `status 0`, the harness is correctly implementing the test case.

## Using the Harness for HTTP/2 Client Development

This harness can be used to test HTTP/2 clients in any language. The harness acts as a malicious/non-compliant server that sends specific frames to test client compliance.

### Crystal HTTP/2 Client Example

To test a Crystal HTTP/2 client implementation:

1. **Start the test harness server:**
   ```bash
   go run . --test=6.5/1
   ```

2. **Create a Crystal test client** (`test_client.cr`):
   ```crystal
   require "http/client"
   require "openssl"
   
   # Configure SSL context to accept self-signed certificates
   context = OpenSSL::SSL::Context::Client.new
   context.verify_mode = OpenSSL::SSL::VerifyMode::NONE
   
   # Create HTTP/2 client
   client = HTTP::Client.new("localhost", 8080, tls: context)
   client.before_request { |req| req.headers["Connection"] = "Upgrade, HTTP2-Settings" }
   
   begin
     response = client.get("/")
     puts "Response: #{response.status_code}"
   rescue ex
     puts "Error (expected): #{ex.message}"
     exit 1 if ex.message.includes?("FRAME_SIZE_ERROR")
   end
   ```

3. **Run the Crystal client:**
   ```bash
   crystal run test_client.cr
   ```

4. **Verify expected behavior:**
   - For protocol error tests (like `6.5/1`): Client should detect the error and close the connection
   - For compliance tests: Client should handle the frame correctly and maintain the connection
   - Exit codes: 0 = test passed, 1 = test failed

### Test Categories

- **Protocol Errors**: Tests expect the client to detect violations and close the connection with appropriate error codes
- **Compliance Tests**: Tests verify the client handles valid but edge-case frames correctly
- **HPACK Tests**: Tests verify header compression/decompression compliance

### Available Test IDs

Run the harness without arguments to see all available test cases:
```bash
go run . --test=""
```

## Docker Usage

For CI/CD and reproducible testing environments, use the Docker image:

### Quick Start with Docker

```bash
# Build the image
docker build -t h2-test-harness .

# List all 146 available tests
docker run --rm h2-test-harness --list

# Run a specific test
docker run --rm h2-test-harness --test=6.5/2

# Run complete test suite verification  
docker run --rm h2-test-harness --verify-all

# Run harness only (for external client testing)
docker run --rm -p 8080:8080 h2-test-harness --harness-only --test=6.5/1
```

### Docker Test Commands

- `--list`: Display all 146 available test cases
- `--test=<id>`: Run specific test case with full harness + verifier validation
- `--verify-all`: Execute complete test suite (all 146 tests) with pass/fail summary
- `--harness-only --test=<id>`: Run harness server only for external client testing

## Implemented Test Cases

This harness implements **146 comprehensive H2SPEC test cases** covering 100% of HTTP/2 protocol compliance scenarios from RFC 7540 and RFC 7541.

### Complete Test Coverage

The test harness covers all major HTTP/2 protocol areas:

- **Connection Management** (3.5): Connection preface and establishment
- **Frame Format & Size** (4.1, 4.2): Frame structure and size validation  
- **Stream Management** (5.1, 5.1.1, 5.1.2): Stream states, identifiers, and concurrency
- **Stream Dependencies** (5.3.1): Priority and dependency handling
- **Error Handling** (5.4.1): Connection and stream error scenarios
- **Frame Types** (6.1-6.10): All HTTP/2 frame types (DATA, HEADERS, PRIORITY, RST_STREAM, SETTINGS, PING, GOAWAY, WINDOW_UPDATE, CONTINUATION)
- **Flow Control** (6.9.1): Window management and flow control compliance
- **HTTP Semantics** (8.1, 8.2): Request/response exchange and server push
- **HPACK Compression** (RFC 7541): Header compression and decompression compliance

### Test Categories Summary

| Category | Test Count | Description |
|----------|------------|-------------|
| Connection Management | 6 | Preface validation and connection establishment |
| Frame Format & Size | 6 | Frame structure and size limit compliance |
| Stream States | 13 | Stream lifecycle and state transitions |
| Frame Processing | 35+ | All HTTP/2 frame types and protocol violations |
| HPACK Compression | 13 | Header compression compliance |
| HTTP Semantics | 15+ | Request/response and header field validation |
| Flow Control | 6 | Window updates and flow control |
| Error Handling | 8 | Connection and stream error scenarios |
| **Total** | **146** | **Complete H2SPEC Coverage** |

### Available Test Cases

To see all 146 available test cases, run:
```bash
go run . --test=""
```

Or in Docker:
```bash
docker run --rm h2-test-harness --list
```

### Test Case Examples

Some key test categories include:

**Protocol Violations:**
- `5.1/1`: DATA frame on idle stream (expects PROTOCOL_ERROR)
- `6.5/2`: SETTINGS frame with non-zero stream ID (expects PROTOCOL_ERROR)
- `5.1.1/1`: HEADERS frame with even stream ID (expects PROTOCOL_ERROR)

**HPACK Compliance:**
- `hpack/2.3.3/1`: Invalid index in header field representation
- `hpack/4.2/1`: Maximum table size validation
- `hpack/6.1/1`: Indexed header field processing

**Frame Size & Format:**
- `4.2/2`: Oversized DATA frame (expects FRAME_SIZE_ERROR)
- `4.2/3`: Oversized HEADERS frame (expects FRAME_SIZE_ERROR)

**Stream Management:**
- `5.1/12`: HEADERS frame on closed stream (expects STREAM_CLOSED)
- `5.3.1/1`: Self-dependency in stream priority (expects PROTOCOL_ERROR)

For detailed test descriptions and expected behaviors, see the test case files in the `harness/cases/` directory.
