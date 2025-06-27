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

## Implemented Test Cases

Below is a list of all implemented test cases.

### Section 6.5: SETTINGS

*   [6.5/1: Sends a SETTINGS frame with ACK flag and payload](./harness/cases/6_5_settings.go)
*   [6.5/2: Sends a SETTINGS frame with a stream identifier other than 0x0](./harness/cases/6_5_settings.go)
*   [6.5/3: Sends a SETTINGS frame with a length other than a multiple of 6 octets](./harness/cases/6_5_settings.go)

### Section 6.5.2: Defined SETTINGS Parameters

*   [6.5.2/1: SETTINGS_ENABLE_PUSH (0x2): Sends the value other than 0 or 1](./harness/cases/6_5_2_defined_settings_parameters.go)
*   [6.5.2/2: SETTINGS_INITIAL_WINDOW_SIZE (0x4): Sends the value above the maximum flow control window size](./harness/cases/6_5_2_defined_settings_parameters.go)
*   [6.5.2/3: SETTINGS_MAX_FRAME_SIZE (0x5): Sends the value below the initial value](./harness/cases/6_5_2_defined_settings_parameters.go)
*   [6.5.2/4: SETTINGS_MAX_FRAME_SIZE (0x5): Sends the value above the maximum allowed frame size](./harness/cases/6_5_2_defined_settings_parameters.go)
*   [6.5.2/5: Sends a SETTINGS frame with unknown identifier](./harness/cases/6_5_2_defined_settings_parameters.go)

### Section 6.5.3: Settings Synchronization

*   [6.5.3/2: Sends a SETTINGS frame without ACK flag](./harness/cases/6_5_3_settings_synchronization.go)

### Section 6.7: PING

*   [6.7/1: Sends a PING frame](./harness/cases/6_7_ping.go)
*   [6.7/2: Sends a PING frame with ACK](./harness/cases/6_7_ping.go)
*   [6.7/3: Sends a PING frame with a stream identifier field value other than 0x0](./harness/cases/6_7_ping.go)
*   [6.7/4: Sends a PING frame with a length field value other than 8](./harness/cases/6_7_ping.go)

### Section 6.8: GOAWAY

*   [6.8/1: Sends a GOAWAY frame with a stream identifier other than 0x0](./harness/cases/6_8_goaway.go)

### Section 6.9: WINDOW_UPDATE

*   [6.9/1: Sends a WINDOW_UPDATE frame with a flow control window increment of 0](./harness/cases/6_9_window_update.go)
*   [6.9/2: Sends a WINDOW_UPDATE frame with a flow control window increment of 0 on a stream](./harness/cases/6_9_window_update.go)
*   [6.9/3: Sends a WINDOW_UPDATE frame with a length other than 4 octets](./harness/cases/6_9_window_update.go)
*   [6.9.2/3: Sends a SETTINGS_INITIAL_WINDOW_SIZE settings with an exceeded maximum window size value](./harness/cases/6_9_2_initial_flow_control_window_size.go)

### Section 6.10: CONTINUATION

*   [6.10/2: Sends a CONTINUATION frame followed by any frame other than CONTINUATION](./harness/cases/6_10_continuation.go)
*   [6.10/3: Sends a CONTINUATION frame with 0x0 stream identifier](./harness/cases/6_10_continuation.go)
*   [6.10/4: Sends a CONTINUATION frame preceded by a HEADERS frame with END_HEADERS flag](./harness/cases/6_10_continuation.go)
*   [6.10/5: Sends a CONTINUATION frame preceded by a CONTINUATION frame with END_HEADERS flag](./harness/cases/6_10_continuation.go)
*   [6.10/6: Sends a CONTINUATION frame preceded by a DATA frame](./harness/cases/6_10_continuation.go)

### Section 8.1: HTTP Request/Response Exchange

*   [8.1/1: Sends a second HEADERS frame without the END_STREAM flag](./harness/cases/8_1_http_request_response_exchange.go)
*   [8.1.2/1: Sends a HEADERS frame that contains the header field name in uppercase letters](./harness/cases/8_1_2_http_header_fields.go)
*   [8.1.2.1/1: Sends a HEADERS frame that contains a unknown pseudo-header field](./harness/cases/8_1_2_1_pseudo_header_fields.go)
*   [8.1.2.1/2: Sends a HEADERS frame that contains the pseudo-header field defined for response](./harness/cases/8_1_2_1_pseudo_header_fields.go)
*   [8.1.2.1/3: Sends a HEADERS frame that contains a pseudo-header field as trailers](./harness/cases/8_1_2_1_pseudo_header_fields.go)
*   [8.1.2.1/4: Sends a HEADERS frame that contains a pseudo-header field that appears in a header block after a regular header field](./harness/cases/8_1_2_1_pseudo_header_fields.go)
*   [8.1.2.2/1: Sends a HEADERS frame that contains the connection-specific header field](./harness/cases/8_1_2_2_connection_specific_header_fields.go)
*   [8.1.2.2/2: Sends a HEADERS frame that contains the TE header field with any value other than "trailers"](./harness/cases/8_1_2_2_connection_specific_header_fields.go)
*   [8.1.2.3/1: Sends a HEADERS frame with empty ":path" pseudo-header field](./harness/cases/8_1_2_3_request_pseudo_header_fields.go)
*   [8.1.2.3/2: Sends a HEADERS frame that omits ":method" pseudo-header field](./harness/cases/8_1_2_3_request_pseudo_header_fields.go)
*   [8.1.2.3/3: Sends a HEADERS frame that omits ":scheme" pseudo-header field](./harness/cases/8_1_2_3_request_pseudo_header_fields.go)
*   [8.1.2.3/4: Sends a HEADERS frame that omits ":path" pseudo-header field](./harness/cases/8_1_2_3_request_pseudo_header_fields.go)
*   [8.1.2.3/5: Sends a HEADERS frame with duplicated ":method" pseudo-header field](./harness/cases/8_1_2_3_request_pseudo_header_fields.go)
*   [8.1.2.3/6: Sends a HEADERS frame with duplicated ":scheme" pseudo-header field](./harness/cases/8_1_2_3_request_pseudo_header_fields.go)
*   [8.1.2.3/7: Sends a HEADERS frame with duplicated ":path" pseudo-header field](./harness/cases/8_1_2_3_request_pseudo_header_fields.go)
*   [8.1.2.6/1: Sends a HEADERS frame with the "content-length" header field which does not equal the DATA frame payload length](./harness/cases/8_1_2_6_malformed_requests_and_responses.go)
*   [8.1.2.6/2: Sends a HEADERS frame with the "content-length" header field which does not equal the sum of the multiple DATA frames payload length](./harness/cases/8_1_2_6_malformed_requests_and_responses.go)

### Section 8.2: Server Push

*   [8.2/1: Sends a PUSH_PROMISE frame](./harness/cases/8_2_server_push.go)

### HPACK

*   [hpack/2.3.3/1: Sends a indexed header field representation with invalid index](./harness/cases/hpack/2_3_3_index_address_space.go)
*   [hpack/2.3.3/2: Sends a literal header field representation with invalid index](./harness/cases/hpack/2_3_3_index_address_space.go)
