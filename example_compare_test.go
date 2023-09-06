package enum_test

import (
	"encoding/json"
	"fmt"

	enum "github.com/FabienMht/go-struct-enum"
)

var (
	// Define States
	TestStateCompareUnknown = &TestStateCompare{enum.New("")}
	TestStateComparePassed  = &TestStateCompare{enum.New("passed")}
	TestStateCompareSkipped = &TestStateCompare{enum.New("skipped")}
	TestStateCompareFailed  = &TestStateCompare{enum.New("failed")}

	// Define the ordered list of states
	// Higher states in the list are considered greater than lower states
	TestStateCompares = []enum.Enummer[string]{
		TestStateCompareUnknown,
		TestStateComparePassed,
		TestStateCompareSkipped,
		TestStateCompareFailed,
	}

	// Define states parsers
	ParseTestStateCompare     = enum.Parse(TestStateCompares)
	MustParseTestStateCompare = enum.MustParse(TestStateCompares)
)

// Define the state enum
type TestStateCompare struct {
	enum.Enum[string]
}

func (ts *TestStateCompare) Equal(other *TestStateCompare) bool {
	return enum.Equal(ts, other)
}

func (ts *TestStateCompare) GreaterThan(other *TestStateCompare) bool {
	return enum.GreaterThan(TestStateCompares)(ts, other)
}

func (ts *TestStateCompare) LessThan(other *TestStateCompare) bool {
	return enum.LessThan(TestStateCompares)(ts, other)
}

func (ts *TestStateCompare) GreaterThanOrEqual(other *TestStateCompare) bool {
	return enum.GreaterThanOrEqual(TestStateCompares)(ts, other)
}

func (ts *TestStateCompare) LessThanOrEqual(other *TestStateCompare) bool {
	return enum.LessThanOrEqual(TestStateCompares)(ts, other)
}

func Example_compare() {
	// Recover from MustParseTestStateCompare panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	// Basic usage
	fmt.Println(TestStateComparePassed)
	fmt.Println(TestStateComparePassed.String())
	fmt.Println(TestStateComparePassed.GetValue())
	fmt.Println(TestStateComparePassed.EqualValue("passed"))
	fmt.Println(TestStateComparePassed.EqualValue("failed"))

	// JSON marshaling and unmarshalling
	fmt.Println(json.Marshal(TestStateComparePassed))
	var result *TestStateCompare
	json.Unmarshal([]byte("\"passed\""), &result)
	fmt.Println(result)

	// Comparison
	fmt.Println(TestStateComparePassed.Equal(TestStateComparePassed))
	fmt.Println(TestStateComparePassed.Equal(TestStateCompareFailed))
	fmt.Println(TestStateComparePassed.GreaterThan(TestStateComparePassed))
	fmt.Println(TestStateComparePassed.GreaterThan(TestStateCompareFailed))
	fmt.Println(TestStateComparePassed.GreaterThanOrEqual(TestStateComparePassed))
	fmt.Println(TestStateComparePassed.GreaterThanOrEqual(TestStateCompareFailed))
	fmt.Println(TestStateComparePassed.LessThan(TestStateComparePassed))
	fmt.Println(TestStateComparePassed.LessThan(TestStateCompareFailed))
	fmt.Println(TestStateComparePassed.LessThanOrEqual(TestStateComparePassed))
	fmt.Println(TestStateComparePassed.LessThanOrEqual(TestStateCompareFailed))

	// Parse string into enum
	fmt.Println(ParseTestStateCompare("passed"))
	fmt.Println(ParseTestStateCompare("xxx"))
	fmt.Println(MustParseTestStateCompare("passed"))
	fmt.Println(MustParseTestStateCompare("xxx"))

	// Output:
	// passed
	// passed
	// passed
	// true
	// false
	// [34 112 97 115 115 101 100 34] <nil>
	// passed
	// true
	// false
	// false
	// false
	// true
	// false
	// false
	// true
	// true
	// true
	// passed
	// <nil>
	// passed
	// recovered from panic: enum: 'xxx' not found
}
