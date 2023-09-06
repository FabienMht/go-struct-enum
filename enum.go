package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
)

// Enummer is an interface that represents an enum.
// Type ~int is the set of all types whose underlying type is int.
// Type ~string is the set of all types whose underlying type is string.
type Enummer[T ~int | ~string] interface {
	String() string
	GetValue() T
	EqualValue(other T) bool
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
	Scan(value interface{}) error
	Value() (driver.Value, error)
}

// Check that Enum implements Enummer.
var _ Enummer[int] = (*Enum[int])(nil)

// Enum is a generic type used to create enums.
type Enum[T ~int | ~string] struct {
	// The value of the enum
	val T
}

// New creates a new enum with the given value.
// The result must be embedded into a struct.
// The embedding struct must be a pointer to
// implement the Enummer interface.
//
// Example:
//
//	type TestState struct {
//		enum.Enum[string]
//	}
//	TestStatePassed = &TestState{enum.New("passed")}
func New[T ~int | ~string](val T) Enum[T] {
	return Enum[T]{val}
}

// GetValue returns the enum underlying value.
func (e Enum[T]) GetValue() T {
	return e.val
}

// String returns the string representation of the enum underlying value.
func (e Enum[T]) String() string {
	return fmt.Sprintf("%v", e.val)
}

// EqualValue returns true if the enums underlying value are equal.
func (e Enum[T]) EqualValue(other T) bool {
	return e.val == other
}

// MarshalJSON implements the json.Marshaler interface.
func (e Enum[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.val)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *Enum[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &e.val)
}

// Scan implements the sql.Scanner interface.
func (e *Enum[T]) Scan(value interface{}) error {
	switch v := value.(type) {
	case T:
		e.val = v
	default:
		return fmt.Errorf("enum: cannot convert '%T' to '%T'", value, e)
	}
	return nil
}

// Value implements the driver.Valuer interface.
func (e Enum[T]) Value() (driver.Value, error) {
	return e.val, nil
}

// Parse parses the given string/int into an Enummer.
// It takes an Enummer list and returns a function
// that takes the given string/int and returns the Enummer.
// It panics if the Enummer in list are not of the same type
// or if the list is empty. If the string/int is not found, it returns nil.
func Parse[T ~int | ~string](list []Enummer[T]) func(T) Enummer[T] {
	checkEnummerListType(list)
	return func(val T) Enummer[T] {
		for _, e := range list {
			if e.EqualValue(val) {
				return e
			}
		}
		return nil
	}
}

// MustParse parses the given string/int into an Enummer.
// It takes an Enummer list and returns a function
// that takes the given string/int and returns the Enummer.
// It panics if the Enummer in list are not of the same type
// or if the list is empty. If the string/int is not found, it panics.
func MustParse[T ~int | ~string](list []Enummer[T]) func(T) Enummer[T] {
	checkEnummerListType(list)
	return func(val T) Enummer[T] {
		for _, e := range list {
			if e.EqualValue(val) {
				return e
			}
		}
		panic(fmt.Sprintf("enum: '%v' not found", val))
	}
}

// Equal returns true if the first Enummer is equal to the second Enummer.
// It panics if the Enummer are not of the same type.
func Equal[T ~int | ~string](a, b Enummer[T]) bool {
	if !compareEnummerType(a, b) {
		panic(fmt.Sprintf("enum: different types '%T' and '%T'", a, b))
	}
	return a.EqualValue(b.GetValue())
}

// GreaterThan returns true if the first Enummer is greater than the second Enummer.
// It takes the Enummer list that defines the order and returns a function.
// Higher indices are considered higher than lower indices.
// It panics if Enummers are not of the same type or if Enummers are not in the list.
func GreaterThan[T ~int | ~string](list []Enummer[T]) func(Enummer[T], Enummer[T]) bool {
	checkEnummerListType(list)
	return func(a, b Enummer[T]) bool {
		// Get list index for values
		ai := compareGetIndex(list, a)
		bi := compareGetIndex(list, b)
		// Compare indexes
		return ai > bi
	}
}

// GreaterThanOrEqual returns true if the first Enummer is greater than or equal to the second Enummer.
// It takes the Enummer list that defines the order and returns a function.
// Higher indices are considered higher than lower indices.
// It panics if Enummers are not of the same type or if Enummers are not in the list.
func GreaterThanOrEqual[T ~int | ~string](list []Enummer[T]) func(Enummer[T], Enummer[T]) bool {
	checkEnummerListType(list)
	return func(a, b Enummer[T]) bool {
		// Get list index for values
		ai := compareGetIndex(list, a)
		bi := compareGetIndex(list, b)
		// Compare indexes
		return ai >= bi
	}
}

// LessThan returns true if the first Enummer is less than the second Enummer.
// It takes the Enummer list that defines the order and returns a function.
// Higher indices are considered higher than lower indices.
// It panics if Enummers are not of the same type or if Enummers are not in the list.
func LessThan[T ~int | ~string](list []Enummer[T]) func(Enummer[T], Enummer[T]) bool {
	checkEnummerListType(list)
	return func(a, b Enummer[T]) bool {
		// Get list index for values
		ai := compareGetIndex(list, a)
		bi := compareGetIndex(list, b)
		// Compare indexes
		return ai < bi
	}
}

// LessThanOrEqual returns true if the first Enummer is less than or equal to the second Enummer.
// It takes the Enummer list that defines the order and returns a function.
// Higher indices are considered higher than lower indices.
// It panics if Enummers are not of the same type or if Enummers are not in the list.
func LessThanOrEqual[T ~int | ~string](list []Enummer[T]) func(Enummer[T], Enummer[T]) bool {
	checkEnummerListType(list)
	return func(a, b Enummer[T]) bool {
		// Get list index for values
		ai := compareGetIndex(list, a)
		bi := compareGetIndex(list, b)
		// Compare indexes
		return ai <= bi
	}
}

// compareGetIndex returns the index of the Enummer in the list.
// It panics if the Enummer is not in the list.
func compareGetIndex[T ~int | ~string](list []Enummer[T], e Enummer[T]) int {
	// Check if e is in the list
	if !existInEnummerList(list, e) {
		panic(fmt.Sprintf("enum: '%v' not found in list", e))
	}
	// Get list index for value
	return slices.IndexFunc(list, func(elem Enummer[T]) bool {
		return Equal(e, elem)
	})
}

// checkEnummerListType panics if the list is empty or
// if the Enummer in the list have different types.
func checkEnummerListType[T ~int | ~string](list []Enummer[T]) {
	if len(list) == 0 {
		panic("enum: list is empty")
	}
	for i, e := range list {
		for j, other := range list {
			// Check the type of the Enummer using reflection
			if i != j && !compareEnummerType(e, other) {
				panic(fmt.Sprintf("enum: different types '%T' and '%T'", e, other))
			}
		}
	}
}

// existInEnummerList returns false if the Enummer is not in the list.
func existInEnummerList[T ~int | ~string](list []Enummer[T], e Enummer[T]) bool {
	for _, other := range list {
		if Equal(other, e) {
			return true
		}
	}
	return false
}

// compareEnummerType returns true if the Enummer are of the same type.
func compareEnummerType[T ~int | ~string](a, b Enummer[T]) bool {
	return getEnummerType(a) == getEnummerType(b)
}

// getEnummerType returns the type of the enum.
// Enum as a pointer: returns the type of the pointer.
// Enum embedded in a struct: returns the type of the struct.
// Enum embedded in a pointer of struct: returns the type of the pointer.
//
// Example:
//
//	type TestInt struct {
//		*Enum[int]
//	}
//	getEnummerType(&Enum[int]{1}) // returns Enum[int]
//	getEnummerType(TestInt{&Enum[int]{1}}) // returns TestInt
//	getEnummerType(&TestInt{&Enum[int]{1}}) // returns TestInt
func getEnummerType[T ~int | ~string](value Enummer[T]) reflect.Type {
	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Ptr {
		return t.Elem()
	}
	return t
}
