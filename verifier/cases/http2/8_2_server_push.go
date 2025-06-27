package http2

import (
	"github.com/nomadlabsinc/h2-client-test-harness/verifier"
	"golang.org/x/net/http2"
)

func init() {
	verifier.Register("8.2/1", func() error {
		return verifier.ExpectConnectionError("PROTOCOL_ERROR")
	})
}
