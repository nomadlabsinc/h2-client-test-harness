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
	// Generic tests
	testRegistry["generic/3.1/1"] = cases.RunTestGeneric3_1_1
	testRegistry["generic/3.1/2"] = cases.RunTestGeneric3_1_2
	testRegistry["generic/3.1/3"] = cases.RunTestGeneric3_1_3
	testRegistry["generic/3.2/1"] = cases.RunTestGeneric3_2_1
	testRegistry["generic/3.2/2"] = cases.RunTestGeneric3_2_2
	testRegistry["generic/3.2/3"] = cases.RunTestGeneric3_2_3
	testRegistry["generic/3.3/1"] = cases.RunTestGeneric3_3_1
	testRegistry["generic/3.3/2"] = cases.RunTestGeneric3_3_2
	testRegistry["generic/3.3/3"] = cases.RunTestGeneric3_3_3
	testRegistry["generic/3.3/4"] = cases.RunTestGeneric3_3_4
	testRegistry["generic/3.3/5"] = cases.RunTestGeneric3_3_5
	testRegistry["generic/3.4/1"] = cases.RunTestGeneric3_4_1
	testRegistry["generic/3.5/1"] = cases.RunTestGeneric3_5_1
	testRegistry["generic/3.7/1"] = cases.RunTestGeneric3_7_1
	testRegistry["generic/3.8/1"] = cases.RunTestGeneric3_8_1
	testRegistry["generic/3.9/1"] = cases.RunTestGeneric3_9_1
	testRegistry["generic/3.10/1"] = cases.RunTestGeneric3_10_1
	testRegistry["generic/4/1"] = cases.RunTestGeneric4_1
	testRegistry["generic/4/2"] = cases.RunTestGeneric4_2

	// 3.5 Connection Preface
	testRegistry["3.5/1"] = cases.RunTest3_5_1
	testRegistry["3.5/2"] = cases.RunTest3_5_2

	// 4.1 Frame Format
	testRegistry["4.1/1"] = cases.RunTest4_1_1
	testRegistry["4.1/2"] = cases.RunTest4_1_2
	testRegistry["4.1/3"] = cases.RunTest4_1_3

	// 4.2 Frame Size
	testRegistry["4.2/1"] = cases.RunTest4_2_1
	testRegistry["4.2/2"] = cases.RunTest4_2_2
	testRegistry["4.2/3"] = cases.RunTest4_2_3

	// 5.1 Stream States
	testRegistry["5.1/1"] = cases.RunTest5_1_1
	testRegistry["5.1/2"] = cases.RunTest5_1_2
	testRegistry["5.1/3"] = cases.RunTest5_1_3
	testRegistry["5.1/4"] = cases.RunTest5_1_4
	testRegistry["5.1/5"] = cases.RunTest5_1_5
	testRegistry["5.1/6"] = cases.RunTest5_1_6
	testRegistry["5.1/7"] = cases.RunTest5_1_7
	testRegistry["5.1/8"] = cases.RunTest5_1_8
	testRegistry["5.1/9"] = cases.RunTest5_1_9
	testRegistry["5.1/10"] = cases.RunTest5_1_10
	testRegistry["5.1/11"] = cases.RunTest5_1_11
	testRegistry["5.1/12"] = cases.RunTest5_1_12
	testRegistry["5.1/13"] = cases.RunTest5_1_13

	// 5.1.1 Stream Identifiers
	testRegistry["5.1.1/1"] = cases.RunTest5_1_1_1
	testRegistry["5.1.1/2"] = cases.RunTest5_1_1_2

	// 5.1.2 Stream Concurrency
	testRegistry["5.1.2/1"] = cases.RunTest5_1_2_1

	// 5.3.1 Stream Dependencies
	testRegistry["5.3.1/1"] = cases.RunTest5_3_1_1
	testRegistry["5.3.1/2"] = cases.RunTest5_3_1_2

	// 5.4.1 Connection Error Handling
	testRegistry["5.4.1/1"] = cases.RunTest5_4_1_1
	testRegistry["5.4.1/2"] = cases.RunTest5_4_1_2

	// 6.1 DATA
	testRegistry["6.1/1"] = cases.RunTest6_1_1
	testRegistry["6.1/2"] = cases.RunTest6_1_2
	testRegistry["6.1/3"] = cases.RunTest6_1_3

	// 6.2 HEADERS
	testRegistry["6.2/1"] = cases.RunTest6_2_1
	testRegistry["6.2/2"] = cases.RunTest6_2_2
	testRegistry["6.2/3"] = cases.RunTest6_2_3
	testRegistry["6.2/4"] = cases.RunTest6_2_4

	// 6.3 PRIORITY
	testRegistry["6.3/1"] = cases.RunTest6_3_1
	testRegistry["6.3/2"] = cases.RunTest6_3_2

	// 6.4 RST_STREAM
	testRegistry["6.4/1"] = cases.RunTest6_4_1
	testRegistry["6.4/2"] = cases.RunTest6_4_2
	testRegistry["6.4/3"] = cases.RunTest6_4_3

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

	// 6.9.1 Flow Control Window
	testRegistry["6.9.1/1"] = cases.RunTest6_9_1_1
	testRegistry["6.9.1/2"] = cases.RunTest6_9_1_2
	testRegistry["6.9.1/3"] = cases.RunTest6_9_1_3

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
	testRegistry["hpack/2.3/1"] = cases.RunTestHpack2_3_1
	testRegistry["hpack/2.3.3/1"] = cases.RunTestHpack2_3_3_1
	testRegistry["hpack/2.3.3/2"] = cases.RunTestHpack2_3_3_2
	testRegistry["hpack/4.1/1"] = cases.RunTestHpack4_1_1
	testRegistry["hpack/4.2/1"] = cases.RunTestHpack4_2_1
	testRegistry["hpack/5.2/1"] = cases.RunTestHpack5_2_1
	testRegistry["hpack/5.2/2"] = cases.RunTestHpack5_2_2
	testRegistry["hpack/5.2/3"] = cases.RunTestHpack5_2_3
	testRegistry["hpack/6.1/1"] = cases.RunTestHpack6_1_1
	testRegistry["hpack/6.2/1"] = cases.RunTestHpack6_2_1
	testRegistry["hpack/6.2.2/1"] = cases.RunTestHpack6_2_2_1
	testRegistry["hpack/6.2.3/1"] = cases.RunTestHpack6_2_3_1
	testRegistry["hpack/6.3/1"] = cases.RunTestHpack6_3_1
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
