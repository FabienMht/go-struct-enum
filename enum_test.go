package enum

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test type with int enum
type TestTypeInt struct {
	Enum[int]
}

// Test2 type with int enum
type Test2TypeInt struct {
	Enum[int]
}

// Composite type with int enum
type TestCompositeInt struct {
	TestType *TestTypeInt `json:"test_type"`
}

// Test type with string enum
type TestTypeString struct {
	Enum[string]
}

// Test2 type with string enum
type Test2TypeString struct {
	Enum[string]
}

// Composite type with string enum
type TestCompositeString struct {
	TestType *TestTypeString `json:"test_type"`
}

func TestEnum_New(t *testing.T) {
	tests := []struct {
		name  string
		value any
		want  any
	}{
		{
			name:  "int",
			value: 1,
			want:  Enum[int]{1},
		},
		{
			name:  "string",
			value: "hello",
			want:  Enum[string]{"hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.value.(type) {
			case int:
				require.Equal(t, tt.want, New(tt.value.(int)))
			case string:
				require.Equal(t, tt.want, New(tt.value.(string)))
			}
		})
	}
}

func TestEnum_GetValue(t *testing.T) {
	tests := []struct {
		name string
		enum any
		want any
	}{
		{
			name: "int",
			enum: &Enum[int]{1},
			want: 1,
		},
		{
			name: "string",
			enum: &Enum[string]{"hello"},
			want: "hello",
		},
		{
			name: "embedded int",
			enum: &TestTypeInt{Enum[int]{1}},
			want: 1,
		},
		{
			name: "embedded string",
			enum: &TestTypeString{Enum[string]{"hello"}},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.enum.(type) {
			case Enummer[int]:
				require.Equal(t, tt.want, tt.enum.(Enummer[int]).GetValue())
			case Enummer[string]:
				require.Equal(t, tt.want, tt.enum.(Enummer[string]).GetValue())
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}

func TestEnum_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		enum any
		want []byte
	}{
		{
			name: "int",
			enum: &Enum[int]{1},
			want: []byte("1"),
		},
		{
			name: "string",
			enum: &Enum[string]{"hello"},
			want: []byte("\"hello\""),
		},
		{
			name: "embedded int",
			enum: &TestTypeInt{Enum[int]{1}},
			want: []byte("1"),
		},
		{
			name: "embedded string",
			enum: &TestTypeString{Enum[string]{"hello"}},
			want: []byte("\"hello\""),
		},
		{
			name: "composite int",
			enum: TestCompositeInt{
				TestType: &TestTypeInt{
					Enum[int]{1},
				},
			},
			want: []byte("{\"test_type\":1}"),
		},
		{
			name: "composite string",
			enum: TestCompositeString{
				TestType: &TestTypeString{
					Enum[string]{"hello"},
				},
			},
			want: []byte("{\"test_type\":\"hello\"}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.enum)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestEnum_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want any
	}{
		{
			name: "int",
			data: []byte("1"),
			want: &Enum[int]{1},
		},
		{
			name: "string",
			data: []byte("\"hello\""),
			want: &Enum[string]{"hello"},
		},
		{
			name: "embedded int",
			data: []byte("1"),
			want: &TestTypeInt{Enum[int]{1}},
		},
		{
			name: "embedded string",
			data: []byte("\"hello\""),
			want: &TestTypeString{Enum[string]{"hello"}},
		},
		{
			name: "composite int",
			data: []byte("{\"test_type\":1}"),
			want: TestCompositeInt{
				TestType: &TestTypeInt{
					Enum[int]{1},
				},
			},
		},
		{
			name: "composite string",
			data: []byte("{\"test_type\":\"hello\"}"),
			want: TestCompositeString{
				TestType: &TestTypeString{
					Enum[string]{"hello"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.want.(type) {
			case *Enum[int]:
				var got *Enum[int]
				err := json.Unmarshal(tt.data, &got)
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			case *Enum[string]:
				var got *Enum[string]
				err := json.Unmarshal(tt.data, &got)
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			case *TestTypeInt:
				var got *TestTypeInt
				err := json.Unmarshal(tt.data, &got)
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			case *TestTypeString:
				var got *TestTypeString
				err := json.Unmarshal(tt.data, &got)
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			case TestCompositeInt:
				var got TestCompositeInt
				err := json.Unmarshal(tt.data, &got)
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			case TestCompositeString:
				var got TestCompositeString
				err := json.Unmarshal(tt.data, &got)
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}

func TestEnum_Scan(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  any
	}{
		{
			name:  "int",
			value: 1,
			want:  &Enum[int]{1},
		},
		{
			name:  "string",
			value: "hello",
			want:  &Enum[string]{"hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.want.(type) {
			case *Enum[int]:
				got := &Enum[int]{}
				err := got.Scan(tt.value)
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			case *Enum[string]:
				got := &Enum[string]{}
				err := got.Scan(tt.value)
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}

func TestEnum_Value(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  any
	}{
		{
			name:  "int",
			value: &Enum[int]{1},
			want:  1,
		},
		{
			name:  "string",
			value: &Enum[string]{"hello"},
			want:  "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.value.(type) {
			case Enummer[int]:
				value, err := tt.value.(Enummer[int]).Value()
				require.NoError(t, err)
				require.Equal(t, tt.want, value)
			case Enummer[string]:
				value, err := tt.value.(Enummer[string]).Value()
				require.NoError(t, err)
				require.Equal(t, tt.want, value)
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name       string
		list       any
		value      any
		want       any
		wantPanics bool
	}{
		{
			name:  "int",
			list:  []Enummer[int]{&Enum[int]{1}},
			value: 1,
			want:  &Enum[int]{1},
		},
		{
			name:  "string",
			list:  []Enummer[string]{&Enum[string]{"hello"}},
			value: "hello",
			want:  &Enum[string]{"hello"},
		},
		{
			name:  "int not found",
			list:  []Enummer[int]{&Enum[int]{1}},
			value: 2,
			want:  nil,
		},
		{
			name:  "string not found",
			list:  []Enummer[string]{&Enum[string]{"hello"}},
			value: "world",
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.list.(type) {
			case []Enummer[int]:
				require.Equal(t, tt.want, Parse(tt.list.([]Enummer[int]))(tt.value.(int)))
			case []Enummer[string]:
				require.Equal(t, tt.want, Parse(tt.list.([]Enummer[string]))(tt.value.(string)))
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}

func TestMustParse(t *testing.T) {
	tests := []struct {
		name       string
		list       any
		value      any
		want       any
		wantPanics bool
	}{
		{
			name:  "int",
			list:  []Enummer[int]{&Enum[int]{1}},
			value: 1,
			want:  &Enum[int]{1},
		},
		{
			name:  "string",
			list:  []Enummer[string]{&Enum[string]{"hello"}},
			value: "hello",
			want:  &Enum[string]{"hello"},
		},
		{
			name:       "int not found",
			list:       []Enummer[int]{&Enum[int]{1}},
			value:      2,
			wantPanics: true,
		},
		{
			name:       "string not found",
			list:       []Enummer[string]{&Enum[string]{"hello"}},
			value:      "world",
			wantPanics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.list.(type) {
			case []Enummer[int]:
				if tt.wantPanics {
					require.Panics(t, func() { MustParse(tt.list.([]Enummer[int]))(tt.value.(int)) })
				} else {
					require.Equal(t, tt.want, MustParse(tt.list.([]Enummer[int]))(tt.value.(int)))
				}
			case []Enummer[string]:
				if tt.wantPanics {
					require.Panics(t, func() { MustParse(tt.list.([]Enummer[string]))(tt.value.(string)) })
				} else {
					require.Equal(t, tt.want, MustParse(tt.list.([]Enummer[string]))(tt.value.(string)))
				}
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name       string
		enumFirst  any
		enumSecond any
		want       bool
		wantPanics bool
	}{
		{
			name:       "int equal",
			enumFirst:  &Enum[int]{1},
			enumSecond: &Enum[int]{1},
			want:       true,
		},
		{
			name:       "int not equal",
			enumFirst:  &Enum[int]{1},
			enumSecond: &Enum[int]{2},
		},
		{
			name:       "string equal",
			enumFirst:  &Enum[string]{"hello"},
			enumSecond: &Enum[string]{"hello"},
			want:       true,
		},
		{
			name:       "string not equal",
			enumFirst:  &Enum[string]{"hello"},
			enumSecond: &Enum[string]{"world"},
		},
		{
			name:       "int two types panic",
			enumFirst:  &TestTypeInt{Enum[int]{1}},
			enumSecond: &Test2TypeInt{Enum[int]{1}},
			wantPanics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.enumFirst.(type) {
			case Enummer[int]:
				if tt.wantPanics {
					require.Panics(t, func() { Equal(tt.enumFirst.(Enummer[int]), tt.enumSecond.(Enummer[int])) })
				} else {
					require.Equal(t, tt.want, Equal(tt.enumFirst.(Enummer[int]), tt.enumSecond.(Enummer[int])))
				}
			case Enummer[string]:
				if tt.wantPanics {
					require.Panics(t, func() { Equal(tt.enumFirst.(Enummer[string]), tt.enumSecond.(Enummer[string])) })
				} else {
					require.Equal(t, tt.want, Equal(tt.enumFirst.(Enummer[string]), tt.enumSecond.(Enummer[string])))
				}
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}

// TestGreaterThan tests the GreaterThan method
func TestGreaterThan(t *testing.T) {
	tests := []struct {
		name       string
		enums      any
		enumFirst  any
		enumSecond any
		want       bool
		wantPanics bool
	}{
		{
			name: "int lower",
			enums: []Enummer[int]{
				&Enum[int]{1},
				&Enum[int]{2},
			},
			enumFirst:  &Enum[int]{1},
			enumSecond: &Enum[int]{2},
		},
		{
			name: "int greater",
			enums: []Enummer[int]{
				&Enum[int]{1},
				&Enum[int]{2},
			},
			enumFirst:  &Enum[int]{2},
			enumSecond: &Enum[int]{1},
			want:       true,
		},
		{
			name: "string lower",
			enums: []Enummer[string]{
				&Enum[string]{"hello"},
				&Enum[string]{"world"},
			},
			enumFirst:  &Enum[string]{"hello"},
			enumSecond: &Enum[string]{"world"},
		},
		{
			name: "string greater",
			enums: []Enummer[string]{
				&Enum[string]{"hello"},
				&Enum[string]{"world"},
			},
			enumFirst:  &Enum[string]{"world"},
			enumSecond: &Enum[string]{"hello"},
			want:       true,
		},
		{
			name: "int two types panic list",
			enums: []Enummer[int]{
				&TestTypeInt{Enum[int]{1}},
				&Test2TypeInt{Enum[int]{1}},
			},
			enumFirst:  &TestTypeInt{Enum[int]{1}},
			enumSecond: &Test2TypeInt{Enum[int]{1}},
			wantPanics: true,
		},
		{
			name: "int two types panic equal",
			enums: []Enummer[int]{
				&TestTypeInt{Enum[int]{1}},
				&TestTypeInt{Enum[int]{2}},
			},
			enumFirst:  &TestTypeInt{Enum[int]{1}},
			enumSecond: &Test2TypeInt{Enum[int]{1}},
			wantPanics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.enumFirst.(type) {
			case Enummer[int]:
				if tt.wantPanics {
					require.Panics(t, func() {
						GreaterThan(tt.enums.([]Enummer[int]))(
							tt.enumFirst.(Enummer[int]), tt.enumSecond.(Enummer[int]),
						)
					})
				} else {
					require.Equal(t, tt.want, GreaterThan(tt.enums.([]Enummer[int]))(
						tt.enumFirst.(Enummer[int]), tt.enumSecond.(Enummer[int]),
					))
				}
			case Enummer[string]:
				if tt.wantPanics {
					require.Panics(t, func() {
						GreaterThan(tt.enums.([]Enummer[string]))(
							tt.enumFirst.(Enummer[string]), tt.enumSecond.(Enummer[string]),
						)
					})
				} else {
					require.Equal(t, tt.want, GreaterThan(tt.enums.([]Enummer[string]))(
						tt.enumFirst.(Enummer[string]), tt.enumSecond.(Enummer[string]),
					))
				}
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	tests := []struct {
		name       string
		enums      any
		enumFirst  any
		enumSecond any
		want       bool
		wantPanics bool
	}{
		{
			name: "int equal",
			enums: []Enummer[int]{
				&Enum[int]{1},
				&Enum[int]{2},
			},
			enumFirst:  &Enum[int]{1},
			enumSecond: &Enum[int]{1},
			want:       true,
		},
		{
			name: "int lower",
			enums: []Enummer[int]{
				&Enum[int]{1},
				&Enum[int]{2},
			},
			enumFirst:  &Enum[int]{1},
			enumSecond: &Enum[int]{2},
		},
		{
			name: "int greater",
			enums: []Enummer[int]{
				&Enum[int]{1},
				&Enum[int]{2},
			},
			enumFirst:  &Enum[int]{2},
			enumSecond: &Enum[int]{1},
			want:       true,
		},
		{
			name: "string equal",
			enums: []Enummer[string]{
				&Enum[string]{"hello"},
				&Enum[string]{"world"},
			},
			enumFirst:  &Enum[string]{"hello"},
			enumSecond: &Enum[string]{"hello"},
			want:       true,
		},
		{
			name: "string lower",
			enums: []Enummer[string]{
				&Enum[string]{"hello"},
				&Enum[string]{"world"},
			},
			enumFirst:  &Enum[string]{"hello"},
			enumSecond: &Enum[string]{"world"},
		},
		{
			name: "string greater",
			enums: []Enummer[string]{
				&Enum[string]{"hello"},
				&Enum[string]{"world"},
			},
			enumFirst:  &Enum[string]{"world"},
			enumSecond: &Enum[string]{"hello"},
			want:       true,
		},
		{
			name: "int two types panic list",
			enums: []Enummer[int]{
				&TestTypeInt{Enum[int]{1}},
				&Test2TypeInt{Enum[int]{1}},
			},
			enumFirst:  &TestTypeInt{Enum[int]{1}},
			enumSecond: &Test2TypeInt{Enum[int]{1}},
			wantPanics: true,
		},
		{
			name: "int two types panic equal",
			enums: []Enummer[int]{
				&TestTypeInt{Enum[int]{1}},
				&TestTypeInt{Enum[int]{2}},
			},
			enumFirst:  &TestTypeInt{Enum[int]{1}},
			enumSecond: &Test2TypeInt{Enum[int]{1}},
			wantPanics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.enumFirst.(type) {
			case Enummer[int]:
				if tt.wantPanics {
					require.Panics(t, func() {
						GreaterThanOrEqual(tt.enums.([]Enummer[int]))(
							tt.enumFirst.(Enummer[int]), tt.enumSecond.(Enummer[int]),
						)
					})
				} else {
					require.Equal(t, tt.want, GreaterThanOrEqual(tt.enums.([]Enummer[int]))(
						tt.enumFirst.(Enummer[int]), tt.enumSecond.(Enummer[int]),
					))
				}
			case Enummer[string]:
				if tt.wantPanics {
					require.Panics(t, func() {
						GreaterThanOrEqual(tt.enums.([]Enummer[string]))(
							tt.enumFirst.(Enummer[string]), tt.enumSecond.(Enummer[string]),
						)
					})
				} else {
					require.Equal(t, tt.want, GreaterThanOrEqual(tt.enums.([]Enummer[string]))(
						tt.enumFirst.(Enummer[string]), tt.enumSecond.(Enummer[string]),
					))
				}
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}

func TestLessThan(t *testing.T) {
	tests := []struct {
		name       string
		enums      any
		enumFirst  any
		enumSecond any
		want       bool
		wantPanics bool
	}{
		{
			name: "int lower",
			enums: []Enummer[int]{
				&Enum[int]{1},
				&Enum[int]{2},
			},
			enumFirst:  &Enum[int]{1},
			enumSecond: &Enum[int]{2},
			want:       true,
		},
		{
			name: "int greater",
			enums: []Enummer[int]{
				&Enum[int]{1},
				&Enum[int]{2},
			},
			enumFirst:  &Enum[int]{2},
			enumSecond: &Enum[int]{1},
		},
		{
			name: "string lower",
			enums: []Enummer[string]{
				&Enum[string]{"hello"},
				&Enum[string]{"world"},
			},
			enumFirst:  &Enum[string]{"hello"},
			enumSecond: &Enum[string]{"world"},
			want:       true,
		},
		{
			name: "string greater",
			enums: []Enummer[string]{
				&Enum[string]{"hello"},
				&Enum[string]{"world"},
			},
			enumFirst:  &Enum[string]{"world"},
			enumSecond: &Enum[string]{"hello"},
		},
		{
			name: "int two types panic list",
			enums: []Enummer[int]{
				&TestTypeInt{Enum[int]{1}},
				&Test2TypeInt{Enum[int]{1}},
			},
			enumFirst:  &TestTypeInt{Enum[int]{1}},
			enumSecond: &Test2TypeInt{Enum[int]{1}},
			wantPanics: true,
		},
		{
			name: "int two types panic equal",
			enums: []Enummer[int]{
				&TestTypeInt{Enum[int]{1}},
				&TestTypeInt{Enum[int]{2}},
			},
			enumFirst:  &TestTypeInt{Enum[int]{1}},
			enumSecond: &Test2TypeInt{Enum[int]{1}},
			wantPanics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.enumFirst.(type) {
			case Enummer[int]:
				if tt.wantPanics {
					require.Panics(t, func() {
						LessThan(tt.enums.([]Enummer[int]))(
							tt.enumFirst.(Enummer[int]), tt.enumSecond.(Enummer[int]),
						)
					})
				} else {
					require.Equal(t, tt.want, LessThan(tt.enums.([]Enummer[int]))(
						tt.enumFirst.(Enummer[int]), tt.enumSecond.(Enummer[int]),
					))
				}
			case Enummer[string]:
				if tt.wantPanics {
					require.Panics(t, func() {
						LessThan(tt.enums.([]Enummer[string]))(
							tt.enumFirst.(Enummer[string]), tt.enumSecond.(Enummer[string]),
						)
					})
				} else {
					require.Equal(t, tt.want, LessThan(tt.enums.([]Enummer[string]))(
						tt.enumFirst.(Enummer[string]), tt.enumSecond.(Enummer[string]),
					))
				}
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}

func TestLessThanOrEqual(t *testing.T) {
	tests := []struct {
		name       string
		enums      any
		enumFirst  any
		enumSecond any
		want       bool
		wantPanics bool
	}{
		{
			name: "int equal",
			enums: []Enummer[int]{
				&Enum[int]{1},
				&Enum[int]{2},
			},
			enumFirst:  &Enum[int]{1},
			enumSecond: &Enum[int]{1},
			want:       true,
		},
		{
			name: "int lower",
			enums: []Enummer[int]{
				&Enum[int]{1},
				&Enum[int]{2},
			},
			enumFirst:  &Enum[int]{1},
			enumSecond: &Enum[int]{2},
			want:       true,
		},
		{
			name: "int greater",
			enums: []Enummer[int]{
				&Enum[int]{1},
				&Enum[int]{2},
			},
			enumFirst:  &Enum[int]{2},
			enumSecond: &Enum[int]{1},
		},
		{
			name: "string equal",
			enums: []Enummer[string]{
				&Enum[string]{"hello"},
				&Enum[string]{"world"},
			},
			enumFirst:  &Enum[string]{"hello"},
			enumSecond: &Enum[string]{"hello"},
			want:       true,
		},
		{
			name: "string lower",
			enums: []Enummer[string]{
				&Enum[string]{"hello"},
				&Enum[string]{"world"},
			},
			enumFirst:  &Enum[string]{"hello"},
			enumSecond: &Enum[string]{"world"},
			want:       true,
		},
		{
			name: "string greater",
			enums: []Enummer[string]{
				&Enum[string]{"hello"},
				&Enum[string]{"world"},
			},
			enumFirst:  &Enum[string]{"world"},
			enumSecond: &Enum[string]{"hello"},
		},
		{
			name: "int two types panic list",
			enums: []Enummer[int]{
				&TestTypeInt{Enum[int]{1}},
				&Test2TypeInt{Enum[int]{1}},
			},
			enumFirst:  &TestTypeInt{Enum[int]{1}},
			enumSecond: &Test2TypeInt{Enum[int]{1}},
			wantPanics: true,
		},
		{
			name: "int two types panic equal",
			enums: []Enummer[int]{
				&TestTypeInt{Enum[int]{1}},
				&TestTypeInt{Enum[int]{2}},
			},
			enumFirst:  &TestTypeInt{Enum[int]{1}},
			enumSecond: &Test2TypeInt{Enum[int]{1}},
			wantPanics: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.enumFirst.(type) {
			case Enummer[int]:
				if tt.wantPanics {
					require.Panics(t, func() {
						LessThanOrEqual(tt.enums.([]Enummer[int]))(
							tt.enumFirst.(Enummer[int]), tt.enumSecond.(Enummer[int]),
						)
					})
				} else {
					require.Equal(t, tt.want, LessThanOrEqual(tt.enums.([]Enummer[int]))(
						tt.enumFirst.(Enummer[int]), tt.enumSecond.(Enummer[int]),
					))
				}
			case Enummer[string]:
				if tt.wantPanics {
					require.Panics(t, func() {
						LessThanOrEqual(tt.enums.([]Enummer[string]))(
							tt.enumFirst.(Enummer[string]), tt.enumSecond.(Enummer[string]),
						)
					})
				} else {
					require.Equal(t, tt.want, LessThanOrEqual(tt.enums.([]Enummer[string]))(
						tt.enumFirst.(Enummer[string]), tt.enumSecond.(Enummer[string]),
					))
				}
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}

func Test_compareEnummerType(t *testing.T) {
	tests := []struct {
		name  string
		enum1 any
		enum2 any
		want  bool
	}{
		{
			name:  "int",
			enum1: &Enum[int]{1},
			enum2: &Enum[int]{1},
			want:  true,
		},
		{
			name:  "string",
			enum1: &Enum[string]{"hello"},
			enum2: &Enum[string]{"hello"},
			want:  true,
		},
		{
			name:  "int two types",
			enum1: &TestTypeInt{Enum[int]{1}},
			enum2: &Test2TypeInt{Enum[int]{1}},
			want:  false,
		},
		{
			name:  "int pointer two types",
			enum1: &TestTypeInt{Enum[int]{1}},
			enum2: &Test2TypeInt{Enum[int]{1}},
			want:  false,
		},
		{
			name:  "string two types",
			enum1: &TestTypeString{Enum[string]{"hello"}},
			enum2: &Test2TypeString{Enum[string]{"hello"}},
			want:  false,
		},
		{
			name:  "string pointer two types",
			enum1: &TestTypeString{Enum[string]{"hello"}},
			enum2: &Test2TypeString{Enum[string]{"hello"}},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.enum1.(type) {
			case Enummer[int]:
				require.Equal(t, tt.want, compareEnummerType(
					tt.enum1.(Enummer[int]), tt.enum2.(Enummer[int]),
				))
			case Enummer[string]:
				require.Equal(t, tt.want, compareEnummerType(
					tt.enum1.(Enummer[string]), tt.enum2.(Enummer[string]),
				))
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}

func Test_getEnummerType(t *testing.T) {
	tests := []struct {
		name string
		enum any
		want reflect.Type
	}{
		{
			name: "int",
			enum: &Enum[int]{1},
			want: reflect.TypeOf(Enum[int]{}),
		},
		{
			name: "string",
			enum: &Enum[string]{"hello"},
			want: reflect.TypeOf(Enum[string]{}),
		},
		{
			name: "embedded int",
			enum: &TestTypeInt{Enum[int]{1}},
			want: reflect.TypeOf(TestTypeInt{}),
		},
		{
			name: "embedded string",
			enum: &TestTypeString{Enum[string]{"hello"}},
			want: reflect.TypeOf(TestTypeString{}),
		},
		{
			name: "embedded int pointer",
			enum: &TestTypeInt{Enum[int]{1}},
			want: reflect.TypeOf(TestTypeInt{}),
		},
		{
			name: "embedded string pointer",
			enum: &TestTypeString{Enum[string]{"hello"}},
			want: reflect.TypeOf(TestTypeString{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.enum.(type) {
			case Enummer[int]:
				require.Equal(t, tt.want, getEnummerType(tt.enum.(Enummer[int])))
			case Enummer[string]:
				require.Equal(t, tt.want, getEnummerType(tt.enum.(Enummer[string])))
			default:
				require.Fail(t, "unknown type")
			}
		})
	}
}
