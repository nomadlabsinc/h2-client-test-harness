# TODO: Test Cases Not Implemented

This file lists `h2spec` test cases that have not been implemented in the Go test harness, along with the reason for their exclusion.

- [ ] **6.5.3/1: Sends multiple values of SETTINGS_INITIAL_WINDOW_SIZE**
  - **Reason:** This test requires verifying the client's internal state (specifically, the order in which it applies settings). `curl` does not expose this level of detail, making it impossible to distinguish between a client that correctly processes settings in order and one that only applies the last setting received. A custom test client or a client with much more verbose, specific logging would be required to verify this behavior correctly.
