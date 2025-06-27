package harness

import (
	"fmt"
	"net"
	"sort"

	"github.com/nomadlabsinc/h2-client-test-harness/harness/cases"
	"golang.org/x/net/http2"
)

type TestFunc func(conn net.Conn, framer *http2.Framer)

var testRegistry = make(map[string]TestFunc)

func init() {
	// 6.5 SETTINGS
	testRegistry["6.5/1"] = cases.RunTest6_5_1
	testRegistry["6.5/2"] = cases.RunTest6_5_2
	testRegistry["6.5/3"] = cases.RunTest6_5_3

	// 6.5.2 Defined SETTINGS Parameters
	testRegistry["6.5.2/1"] = cases.RunTest6_5_2_1
	testRegistry["6.5.2/2"] = cases.RunTest6_5_2_2
	testRegistry["6.5.2/3"] = cases.RunTest6_5_2_3
	testRegistry["6.5.2/4"] = cases.RunTest6_5_2_4
	testRegistry["6.5.2/5"] = cases.RunTest6_5_2_5

	// 6.5.3 Settings Synchronization
	testRegistry["6.5.3/2"] = cases.RunTest6_5_3_2

	// 6.7 PING
	testRegistry["6.7/1"] = cases.RunTest6_7_1
	testRegistry["6.7/2"] = cases.RunTest6_7_2
	testRegistry["6.7/3"] = cases.RunTest6_7_3
	testRegistry["6.7/4"] = cases.RunTest6_7_4

	// 6.8 GOAWAY
	testRegistry["6.8/1"] = cases.RunTest6_8_1

	// 6.9 WINDOW_UPDATE
	testRegistry["6.9/1"] = cases.RunTest6_9_1
	testRegistry["6.9/2"] = cases.RunTest6_9_2
	testRegistry["6.9/3"] = cases.RunTest6_9_3
	testRegistry["6.9.2/3"] = cases.RunTest6_9_2_3

	// 6.10 CONTINUATION
	testRegistry["6.10/2"] = cases.RunTest6_10_2
	testRegistry["6.10/3"] = cases.RunTest6_10_3
	testRegistry["6.10/4"] = cases.RunTest6_10_4
	testRegistry["6.10/5"] = cases.RunTest6_10_5
	testRegistry["6.10/6"] = cases.RunTest6_10_6

	// 8.1 HTTP Request/Response Exchange
	testRegistry["8.1/1"] = cases.RunTest8_1_1

	// 8.1.2 HTTP Header Fields
	testRegistry["8.1.2/1"] = cases.RunTest8_1_2_1

	// 8.1.2.1 Pseudo-Header Fields
	testRegistry["8.1.2.1/1"] = cases.RunTest8_1_2_1_1
	testRegistry["8.1.2.1/2"] = cases.RunTest8_1_2_1_2
	testRegistry["8.1.2.1/3"] = cases.RunTest8_1_2_1_3
	testRegistry["8.1.2.1/4"] = cases.RunTest8_1_2_1_4

	// 8.1.2.2 Connection-Specific Header Fields
	testRegistry["8.1.2.2/1"] = cases.RunTest8_1_2_2_1
	testRegistry["8.1.2.2/2"] = cases.RunTest8_1_2_2_2

	// 8.1.2.3 Request Pseudo-Header Fields
	testRegistry["8.1.2.3/1"] = cases.RunTest8_1_2_3_1
	testRegistry["8.1.2.3/2"] = cases.RunTest8_1_2_3_2
	testRegistry["8.1.2.3/3"] = cases.RunTest8_1_2_3_3
	testRegistry["8.1.2.3/4"] = cases.RunTest8_1_2_3_4
	testRegistry["8.1.2.3/5"] = cases.RunTest8_1_2_3_5
	testRegistry["8.1.2.3/6"] = cases.RunTest8_1_2_3_6
	testRegistry["8.1.2.3/7"] = cases.RunTest8_1_2_3_7

	// 8.1.2.6 Malformed Requests and Responses
	testRegistry["8.1.2.6/1"] = cases.RunTest8_1_2_6_1
	testRegistry["8.1.2.6/2"] = cases.RunTest8_1_2_6_2

	// 8.2 Server Push
	testRegistry["8.2/1"] = cases.RunTest8_2_1

	// HPACK
	testRegistry["hpack/2.3.3/1"] = cases.RunTestHpack2_3_3_1
	testRegistry["hpack/2.3.3/2"] = cases.RunTestHpack2_3_3_2
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
