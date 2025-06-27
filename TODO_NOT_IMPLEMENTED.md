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
