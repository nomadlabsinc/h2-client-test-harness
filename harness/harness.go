package harness

import (
	"fmt"
	"net"
	"sort"

	"github.com/nomadlabsinc/h2-client-test-harness/harness/http2"
	"golang.org/x/net/http2"
)

type TestFunc func(conn net.Conn, framer *http2.Framer)

var testRegistry = make(map[string]TestFunc)

func init() {
	// 6.5 SETTINGS
	testRegistry["6.5/1"] = http2.RunTest6_5_1
	testRegistry["6.5/2"] = http2.RunTest6_5_2
	testRegistry["6.5/3"] = http2.RunTest6_5_3

	// 6.5.2 Defined SETTINGS Parameters
	testRegistry["6.5.2/1"] = http2.RunTest6_5_2_1
	testRegistry["6.5.2/2"] = http2.RunTest6_5_2_2
	testRegistry["6.5.2/3"] = http2.RunTest6_5_2_3
	testRegistry["6.5.2/4"] = http2.RunTest6_5_2_4
	testRegistry["6.5.2/5"] = http2.RunTest6_5_2_5

	// 6.5.3 Settings Synchronization
	testRegistry["6.5.3/2"] = http2.RunTest6_5_3_2

	// 6.7 PING
	testRegistry["6.7/1"] = http2.RunTest6_7_1
	testRegistry["6.7/2"] = http2.RunTest6_7_2
	testRegistry["6.7/3"] = http2.RunTest6_7_3
	testRegistry["6.7/4"] = http2.RunTest6_7_4

	// 6.8 GOAWAY
	testRegistry["6.8/1"] = http2.RunTest6_8_1

	// 6.9 WINDOW_UPDATE
	testRegistry["6.9/1"] = http2.RunTest6_9_1
}

func GetTest(id string) (TestFunc, bool) {
	test, ok := testRegistry[id]
	return test, ok
}

func PrintAllTests() {
	fmt.Println("Available test cases:")
	keys := make([]string, 0, len(testRegistry))
	for k := range testRegistry {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("  - %s\n", k)
	}
}
