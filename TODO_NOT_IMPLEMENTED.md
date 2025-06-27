# TODO: Test Cases Not Implemented

This file lists `h2spec` test cases that have not been implemented in the Go test harness, along with the reason for their exclusion.

- [ ] **6.5.3/1: Sends multiple values of SETTINGS_INITIAL_WINDOW_SIZE**
  - **Reason:** This test requires verifying the client's internal state (specifically, the order in which it applies settings). `curl` does not expose this level of detail, making it impossible to distinguish between a client that correctly processes settings in order and one that only applies the last setting received. A custom test client or a client with much more verbose, specific logging would be required to verify this behavior correctly.

- [ ] **6.9.1/1: Sends SETTINGS frame to set the initial window size to 1 and sends HEADERS frame**
  - **Reason:** This test requires the client to send a request with a body, so the harness can verify that the client respects the small flow control window. `curl` with a simple `GET` request does not send a `DATA` frame, so this behavior cannot be triggered and verified.

- [ ] **6.9.1/2: Sends multiple WINDOW_UPDATE frames increasing the flow control window to above 2^31-1**
  - **Reason:** This test requires the client to send `WINDOW_UPDATE` frames. `curl`'s behavior in sending `WINDOW_UPDATE` is not easily controllable or predictable for the purpose of this test, making it impossible to reliably verify that the harness correctly handles the excessive window size.

- [ ] **6.9.1/3: Sends multiple WINDOW_UPDATE frames increasing the flow control window to above 2^31-1 on a stream**
  - **Reason:** Same as `6.9.1/2`.

- [ ] **6.9.2/1: Changes SETTINGS_INITIAL_WINDOW_SIZE after sending HEADERS frame**
  - **Reason:** Similar to `6.9.1/1`, this test requires the client to have a flow-controlled frame (like `DATA`) ready to send, so we can verify that it correctly adjusts its sending based on the new window size. This is not possible with a simple `curl` `GET` request.

- [ ] **6.9.2/2: Sends a SETTINGS frame for window size to be negative**
  - **Reason:** This is a complex scenario that requires the client to send data, have its window become negative, and then correctly resume sending data after a `WINDOW_UPDATE`. This level of control and observation is not possible with `curl`.

- [ ] **6.10/1: Sends multiple CONTINUATION frames preceded by a HEADERS frame**
  - **Reason:** This test requires the client to successfully parse a fragmented header block. `curl` does not provide a direct way to verify that the headers were received and parsed correctly, only that the request as a whole succeeded or failed.

- [ ] **7/1: Sends a GOAWAY frame with unknown error code**
  - **Reason:** This test expects the client to either gracefully close the connection or ignore the frame. `curl` will likely just close the connection, which is a valid outcome. However, we cannot distinguish this from a crash or other incorrect behavior. A more sophisticated client with introspection capabilities is needed to verify this test.

- [ ] **7/2: Sends a RST_STREAM frame with unknown error code**
  - **Reason:** Similar to the `GOAWAY` test above, this test has ambiguous success conditions that are difficult to distinguish with a simple client like `curl`.

- [ ] **generic/2/1: Sends a PRIORITY frame on idle stream**
  - **Reason:** This test requires the client to correctly handle a `PRIORITY` frame on an idle stream. `curl` does not provide a way to verify that the `PRIORITY` frame was received and processed correctly.

- [ ] **generic/2/2: Sends a WINDOW_UPDATE frame on half-closed (remote) stream**
  - **Reason:** This test requires the client to be in a half-closed (remote) state, which is not easily achievable or verifiable with `curl`.

- [ ] **generic/2/3: Sends a PRIORITY frame on half-closed (remote) stream**
  - **Reason:** This test requires the client to be in a half-closed (remote) state, which is not easily achievable or verifiable with `curl`.

- [ ] **generic/2/4: Sends a RST_STREAM frame on half-closed (remote) stream**
  - **Reason:** This test requires the client to be in a half-closed (remote) state, which is not easily achievable or verifiable with `curl`.

- [ ] **generic/2/5: Sends a PRIORITY frame on closed stream**
  - **Reason:** This test requires the client to have a closed stream, and then correctly handle a `PRIORITY` frame on that stream. This is not easily verifiable with `curl`.

- [ ] **generic/3.1/1: Sends a DATA frame**
  - **Reason:** This test requires the client to send a request with a body. `curl` with a simple `GET` request does not send a `DATA` frame, so this behavior cannot be triggered and verified.

- [ ] **generic/3.1/2: Sends multiple DATA frames**
  - **Reason:** This test requires the client to send a request with a body. `curl` with a simple `GET` request does not send a `DATA` frame, so this behavior cannot be triggered and verified.

- [ ] **generic/3.1/3: Sends a DATA frame with padding**
  - **Reason:** This test requires the client to send a request with a body. `curl` with a simple `GET` request does not send a `DATA` frame, so this behavior cannot be triggered and verified.

- [ ] **generic/3.2/2: Sends a HEADERS frame with padding**
  - **Reason:** This test requires the client to correctly handle a `HEADERS` frame with padding. `curl` does not provide a way to verify that the padding was correctly handled.

- [ ] **generic/3.2/3: Sends a HEADERS frame with priority**
  - **Reason:** This test requires the client to correctly handle a `HEADERS` frame with priority information. `curl` does not provide a way to verify that the priority information was correctly handled.

- [ ] **generic/3.3/1: Sends a PRIORITY frame with priority 1**
  - **Reason:** This test requires the client to correctly handle a `PRIORITY` frame. `curl` does not provide a way to verify that the priority information was correctly handled.

- [ ] **generic/3.3/2: Sends a PRIORITY frame with priority 256**
  - **Reason:** This test requires the client to correctly handle a `PRIORITY` frame. `curl` does not provide a way to verify that the priority information was correctly handled.

- [ ] **generic/3.3/3: Sends a PRIORITY frame with stream dependency**
  - **Reason:** This test requires the client to correctly handle a `PRIORITY` frame with stream dependency. `curl` does not provide a way to verify that the priority information was correctly handled.

- [ ] **generic/3.3/4: Sends a PRIORITY frame with exclusive**
  - **Reason:** This test requires the client to correctly handle a `PRIORITY` frame with the exclusive flag. `curl` does not provide a way to verify that the priority information was correctly handled.

- [ ] **generic/3.3/5: Sends a PRIORITY frame for an idle stream, then send a HEADER frame for a lower stream ID**
  - **Reason:** This test requires the client to correctly handle a `PRIORITY` frame on an idle stream. `curl` does not provide a way to verify that the priority information was correctly handled.

- [ ] **generic/3.4/1: Sends a RST_STREAM frame**
  - **Reason:** This test has ambiguous success conditions (either a PING ACK or a closed connection) that are difficult to distinguish from a crash or other incorrect behavior with a simple client like `curl`.
