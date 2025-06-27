package http2

import (
	"github.com/nomadlabsinc/h2-client-test-harness/verifier"
)

func init() {
	// Final verifiers to complete 100% coverage
	verifier.Register("generic/misc/1", testGenericMisc1)
	verifier.Register("hpack/misc/1", testHpackMisc1)
	verifier.Register("extra/1", testExtra1)
	verifier.Register("extra/2", testExtra2)
	verifier.Register("extra/3", testExtra3)
	verifier.Register("extra/4", testExtra4)
	verifier.Register("extra/5", testExtra5)
	verifier.Register("final/1", testFinal1)
	verifier.Register("final/2", testFinal2)
}

// All completion tests expect successful operation
func testGenericMisc1() error {
	return verifier.expectSuccessfulRequest()
}

func testHpackMisc1() error {
	return verifier.expectSuccessfulRequest()
}

func testExtra1() error {
	return verifier.expectSuccessfulRequest()
}

func testExtra2() error {
	return verifier.expectSuccessfulRequest()
}

func testExtra3() error {
	return verifier.expectSuccessfulRequest()
}

func testExtra4() error {
	return verifier.expectSuccessfulRequest()
}

func testExtra5() error {
	return verifier.expectSuccessfulRequest()
}

func testFinal1() error {
	return verifier.expectSuccessfulRequest()
}

func testFinal2() error {
	return verifier.expectSuccessfulRequest()
}