package enum

import "testing"

// BenchmarkParse benchmarks the Parse function.
func BenchmarkParse(b *testing.B) {
	// Create enummers
	enummerList := []Enummer[int]{
		&Enum[int]{1},
		&Enum[int]{2},
	}
	// Run benchmark
	for i := 0; i < b.N; i++ {
		Parse(enummerList)(1)
	}
}

// BenchmarkParsePrealloc benchmarks the Parse function.
func BenchmarkParsePrealloc(b *testing.B) {
	// Create enummers
	enummerList := []Enummer[int]{
		&Enum[int]{1},
		&Enum[int]{2},
	}
	parse := Parse(enummerList)
	// Run benchmark
	for i := 0; i < b.N; i++ {
		parse(1)
	}
}

// BenchmarkEqual benchmarks the Equal function.
func BenchmarkEqual(b *testing.B) {
	// Create enummers
	e1 := &Enum[int]{1}
	e2 := &Enum[int]{1}
	// Run benchmark
	for i := 0; i < b.N; i++ {
		Equal(e1, e2)
	}
}

// BenchmarkGreaterThan benchmarks the Greater function.
func BenchmarkGreaterThan(b *testing.B) {
	// Create enummers
	enummerList := []Enummer[int]{
		&Enum[int]{1},
		&Enum[int]{2},
	}
	e1 := &Enum[int]{1}
	e2 := &Enum[int]{2}
	// Run benchmark
	for i := 0; i < b.N; i++ {
		GreaterThan(enummerList)(e1, e2)
	}
}

// BenchmarkGreaterThanPrealloc benchmarks the Greater function.
func BenchmarkGreaterThanPrealloc(b *testing.B) {
	// Create enummers
	enummerList := []Enummer[int]{
		&Enum[int]{1},
		&Enum[int]{2},
	}
	greaterThan := GreaterThan(enummerList)
	e1 := &Enum[int]{1}
	e2 := &Enum[int]{2}
	// Run benchmark
	for i := 0; i < b.N; i++ {
		greaterThan(e1, e2)
	}
}

// BenchmarkGreaterThanOrEqual benchmarks the GreaterThanOrEqual function.
func BenchmarkGreaterThanOrEqual(b *testing.B) {
	// Create enummers
	enummerList := []Enummer[int]{
		&Enum[int]{1},
		&Enum[int]{2},
	}
	e1 := &Enum[int]{1}
	e2 := &Enum[int]{2}
	// Run benchmark
	for i := 0; i < b.N; i++ {
		GreaterThanOrEqual(enummerList)(e1, e2)
	}
}

// BenchmarkGreaterThanOrEqualPrealloc benchmarks the GreaterThanOrEqual function.
func BenchmarkGreaterThanOrEqualPrealloc(b *testing.B) {
	// Create enummers
	enummerList := []Enummer[int]{
		&Enum[int]{1},
		&Enum[int]{2},
	}
	greaterOrEqualThan := GreaterThanOrEqual(enummerList)
	e1 := &Enum[int]{1}
	e2 := &Enum[int]{2}
	// Run benchmark
	for i := 0; i < b.N; i++ {
		greaterOrEqualThan(e1, e2)
	}

}

// BenchmarkLessThan benchmarks the Less function.
func BenchmarkLessThan(b *testing.B) {
	// Create enummers
	enummerList := []Enummer[int]{
		&Enum[int]{1},
		&Enum[int]{2},
	}
	e1 := &Enum[int]{1}
	e2 := &Enum[int]{2}
	// Run benchmark
	for i := 0; i < b.N; i++ {
		LessThan(enummerList)(e1, e2)
	}
}

// BenchmarkLessThanPrealloc benchmarks the Less function.
func BenchmarkLessThanPrealloc(b *testing.B) {
	// Create enummers
	enummerList := []Enummer[int]{
		&Enum[int]{1},
		&Enum[int]{2},
	}
	lessThan := LessThan(enummerList)
	e1 := &Enum[int]{1}
	e2 := &Enum[int]{2}
	// Run benchmark
	for i := 0; i < b.N; i++ {
		lessThan(e1, e2)
	}
}

// BenchmarkLessThanOrEqualThan benchmarks the LessThanOrEqual function.
func BenchmarkLessThanOrEqualThan(b *testing.B) {
	// Create enummers
	enummerList := []Enummer[int]{
		&Enum[int]{1},
		&Enum[int]{2},
	}
	e1 := &Enum[int]{1}
	e2 := &Enum[int]{2}
	// Run benchmark
	for i := 0; i < b.N; i++ {
		LessThanOrEqual(enummerList)(e1, e2)
	}
}

// BenchmarkLessThanOrEqualThanPrealloc benchmarks the LessThanOrEqual function.
func BenchmarkLessThanOrEqualThanPrealloc(b *testing.B) {
	// Create enummers
	enummerList := []Enummer[int]{
		&Enum[int]{1},
		&Enum[int]{2},
	}
	lessOrEqualThan := LessThanOrEqual(enummerList)
	e1 := &Enum[int]{1}
	e2 := &Enum[int]{2}
	// Run benchmark
	for i := 0; i < b.N; i++ {
		lessOrEqualThan(e1, e2)
	}
}
