# go-struct-enum

![Go Version](https://img.shields.io/github/go-mod/go-version/FabienMht/go-struct-enum.svg)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/FabienMht/go-struct-enum)
[![Go Report Card](https://goreportcard.com/badge/github.com/FabienMht/go-struct-enum)](https://goreportcard.com/report/github.com/FabienMht/go-struct-enum)
[![Sourcegraph](https://sourcegraph.com/github.com/FabienMht/go-struct-enum/-/badge.svg)](https://sourcegraph.com/github.com/FabienMht/go-struct-enum)
[![Tag](https://img.shields.io/github/tag/FabienMht/go-struct-enum.svg)](https://github.com/FabienMht/go-struct-enum/tags)
[![Contributors](https://img.shields.io/github/contributors/FabienMht/go-struct-enum)](https://github.com/FabienMht/go-struct-enum/graphs/contributors)
[![License](https://img.shields.io/github/license/FabienMht/go-struct-enum)](./LICENSE)

A fully featured Golang struct based enums.
It provides a simple way to declare enums and use them in structs.

**Enums are [harden](https://github.com/nikolaydubina/go-enum-example/tree/master) against:**
- Arithmetics operations
- Comparison operators except == and !=
- Implicit cast of [untyped constants](https://medium.com/golangspec/untyped-constants-in-go-2c69eb519b5b) (use the based type instead of the enum type)

**An enum implements the following interfaces:**
- `fmt.Stringer`
- `json.Marshaler`
- `json.Unmarshaler`
- `sql.Scanner`
- `driver.Valuer`

## Install

**Download it:**

```bash
$ go get github.com/FabienMht/go-struct-enum
```

**Add the following import:**

```go
enum "github.com/FabienMht/go-struct-enum"
```

## Usage

### Basic

Checkout the detailed example in the [documentation](https://pkg.go.dev/github.com/FabienMht/go-struct-enum#pkg-examples) for more information.

**Declare enum values:**

```go
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
```

**Use enum values:**

```go
// Parse a state
state := ParseTestState("passed") // TestStatePassed
// Or panic if the state is not valid
state := MustParseTestState("passed") // TestStatePassed

// Get the state value
state.Value() // "passed"
// Check the enum value
state.EqualValue("passed") // true
```

**Use enum in structs:**

```go
type Test struct {
    State *TestState `json:"state"`
}

// Marshal a struct with an enum
json.Marshal(&Test{State: TestStatePassed}) // {"state":"passed"}

// Unmarshal a struct with an enum
var test Test
json.Unmarshal([]byte(`{"state":"passed"}`), &test) // &Test{State: TestStatePassed}
```

### Comparison

Checkout the detailed example in the [documentation](https://pkg.go.dev/github.com/FabienMht/go-struct-enum#pkg-examples) for more information.

**Define comparison functions:**

```go
// Define the state enum
type TestState struct {
	enum.Enum[string]
}

func (ts *TestState) Equal(other *TestState) bool {
	return enum.Equal(ts, other)
}

func (ts *TestState) GreaterThan(other *TestState) bool {
	return enum.GreaterThan(TestStates)(ts, other)
}

func (ts *TestState) LessThan(other *TestState) bool {
	return enum.LessThan(TestStates)(ts, other)
}

func (ts *TestState) GreaterThanOrEqual(other *TestState) bool {
	return enum.GreaterThanOrEqual(TestStates)(ts, other)
}

func (ts *TestState) LessThanOrEqual(other *TestState) bool {
	return enum.LessThanOrEqual(TestStates)(ts, other)
}
```

**Use comparison functions:**

```go
TestStatePassed.Equal(TestStatePassed) // true
TestStatePassed.Equal(TestStateFailed) // false
TestStatePassed.GreaterThan(TestStatePassed) // false
TestStatePassed.GreaterThan(TestStateFailed) // false
TestStatePassed.GreaterThanOrEqual(TestStatePassed) // true
TestStatePassed.GreaterThanOrEqual(TestStateFailed) // false
TestStatePassed.LessThan(TestStatePassed) // false
TestStatePassed.LessThan(TestStateFailed) // true
TestStatePassed.LessThanOrEqual(TestStatePassed) // true
TestStatePassed.LessThanOrEqual(TestStateFailed) // true
```

## Benchmark

```bash
$ task bench
task: [bench] go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/FabienMht/go-struct-enum
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkParse-8                                15764276                77.69 ns/op           48 B/op          1 allocs/op
BenchmarkParsePrealloc-8                        383380515                3.034 ns/op           0 B/op          0 allocs/op
BenchmarkEqual-8                                63047527                16.52 ns/op            0 B/op          0 allocs/op
BenchmarkGreaterThan-8                           6585123               187.9 ns/op            48 B/op          1 allocs/op
BenchmarkGreaterThanPrealloc-8                  10859983               108.9 ns/op             0 B/op          0 allocs/op
BenchmarkGreaterThanOrEqual-8                    6594724               184.5 ns/op            48 B/op          1 allocs/op
BenchmarkGreaterThanOrEqualPrealloc-8           10677360               107.0 ns/op             0 B/op          0 allocs/op
BenchmarkLessThan-8                              6341217               183.0 ns/op            48 B/op          1 allocs/op
BenchmarkLessThanPrealloc-8                     10930384               107.0 ns/op             0 B/op          0 allocs/op
BenchmarkLessThanOrEqualThan-8                   6474268               182.1 ns/op            48 B/op          1 allocs/op
BenchmarkLessThanOrEqualThanPrealloc-8          10892846               107.9 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/FabienMht/go-struct-enum     14.543s
```

## Contributing

Contributions are welcome ! Please open an issue or submit a pull request.

```bash
# Install task
$ go install github.com/go-task/task/v3/cmd/task@v3

# Install dev dependencies
$ task dev

# Run linter
$ task lint

# Run tests
$ task test

# Run benchmarks
$ task bench
```
