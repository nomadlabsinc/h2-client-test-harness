# HTTP/2 Client Test Harness

This repository contains a Go-based test harness for testing the compliance of HTTP/2 clients. It is inspired by [h2spec](https://github.com/summerwind/h2spec), but the logic is inverted: this harness acts as a server that sends specific, sometimes malformed, frames to a client to test the client's response.

## Usage

1.  **Build the harness:**
    ```shell
    go build
    ```

2.  **Run a specific test case:**
    ```shell
    ./h2-client-test-harness --test=<test_case_id>
    ```
    For example:
    ```shell
    ./h2-client-test-harness --test=6.5/1
    ```

3.  **List all available test cases:**
    ```shell
    ./h2-client-test-harness --test=
    ```

## Implemented Test Cases

Below is a list of all implemented test cases.

### Section 6.5: SETTINGS

*   [6.5/1: Sends a SETTINGS frame with ACK flag and payload](./harness/http2/6_5_settings.go)
*   [6.5/2: Sends a SETTINGS frame with a stream identifier other than 0x0](./harness/http2/6_5_settings.go)
*   [6.5/3: Sends a SETTINGS frame with a length other than a multiple of 6 octets](./harness/http2/6_5_settings.go)

### Section 6.5.2: Defined SETTINGS Parameters

*   [6.5.2/1: SETTINGS_ENABLE_PUSH (0x2): Sends the value other than 0 or 1](./harness/http2/6_5_2_defined_settings_parameters.go)
*   [6.5.2/2: SETTINGS_INITIAL_WINDOW_SIZE (0x4): Sends the value above the maximum flow control window size](./harness/http2/6_5_2_defined_settings_parameters.go)
*   [6.5.2/3: SETTINGS_MAX_FRAME_SIZE (0x5): Sends the value below the initial value](./harness/http2/6_5_2_defined_settings_parameters.go)
*   [6.5.2/4: SETTINGS_MAX_FRAME_SIZE (0x5): Sends the value above the maximum allowed frame size](./harness/http2/6_5_2_defined_settings_parameters.go)
*   [6.5.2/5: Sends a SETTINGS frame with unknown identifier](./harness/http2/6_5_2_defined_settings_parameters.go)

### Section 6.5.3: Settings Synchronization

*   [6.5.3/2: Sends a SETTINGS frame without ACK flag](./harness/http2/6_5_3_settings_synchronization.go)

### Section 6.7: PING

*   [6.7/1: Sends a PING frame](./harness/http2/6_7_ping.go)
*   [6.7/2: Sends a PING frame with ACK](./harness/http2/6_7_ping.go)
*   [6.7/3: Sends a PING frame with a stream identifier field value other than 0x0](./harness/http2/6_7_ping.go)
*   [6.7/4: Sends a PING frame with a length field value other than 8](./harness/http2/6_7_ping.go)

### Section 6.8: GOAWAY

*   [6.8/1: Sends a GOAWAY frame with a stream identifier other than 0x0](./harness/http2/6_8_goaway.go)

### Section 6.9: WINDOW_UPDATE

*   [6.9/1: Sends a WINDOW_UPDATE frame with a flow control window increment of 0](./harness/http2/6_9_window_update.go)
