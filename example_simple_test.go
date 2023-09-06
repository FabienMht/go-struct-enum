package enum_test

import (
	"encoding/json"
	"fmt"

	enum "github.com/FabienMht/go-struct-enum"
)

var (
	// Define States
	TestStateUnknown = &TestState{enum.New("")}
	TestStatePassed  = &TestState{enum.New("passed")}
	TestStateSkipped = &TestState{enum.New("skipped")}
	TestStateFailed  = &TestState{enum.New("failed")}

	// Define the ordered list of states
	TestStates = []enum.Enummer[string]{
		TestStateUnknown,
		TestStatePassed,
		TestStateSkipped,
		TestStateFailed,
	}

	// Define states parsers
	ParseTestState     = enum.Parse(TestStates)
	MustParseTestState = enum.MustParse(TestStates)
)

// Define the state enum
type TestState struct {
	enum.Enum[string]
}

func Example() {
	// Recover from MustParseTestState panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	// Basic usage
	fmt.Println(TestStatePassed)
	fmt.Println(TestStatePassed.String())
	fmt.Println(TestStatePassed.GetValue())
	fmt.Println(TestStatePassed.EqualValue("passed"))
	fmt.Println(TestStatePassed.EqualValue("failed"))

	// JSON marshaling and unmarshalling
	fmt.Println(json.Marshal(TestStatePassed))
	var result *TestState
	json.Unmarshal([]byte("\"passed\""), &result)
	fmt.Println(result)

	// Parse string into enum
	fmt.Println(ParseTestState("passed"))
	fmt.Println(ParseTestState("xxx"))
	fmt.Println(MustParseTestState("passed"))
	fmt.Println(MustParseTestState("xxx"))

	// Output:
	// passed
	// passed
	// passed
	// true
	// false
	// [34 112 97 115 115 101 100 34] <nil>
	// passed
	// passed
	// <nil>
	// passed
	// recovered from panic: enum: 'xxx' not found
}
