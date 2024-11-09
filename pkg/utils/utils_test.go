package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistinct(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "with duplicates",
			input:    []int{1, 2, 2, 3, 3, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Distinct(tt.input)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}

func TestFilter(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	tests := []struct {
		name     string
		input    []int
		filter   func(int) bool
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			filter:   isEven,
			expected: []int{},
		},
		{
			name:     "filter even numbers",
			input:    []int{1, 2, 3, 4, 5, 6},
			filter:   isEven,
			expected: []int{2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Filter(tt.input, tt.filter)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFilterNot(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	tests := []struct {
		name     string
		input    []int
		filter   func(int) bool
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			filter:   isEven,
			expected: []int{},
		},
		{
			name:     "filter not even numbers",
			input:    []int{1, 2, 3, 4, 5, 6},
			filter:   isEven,
			expected: []int{1, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterNot(tt.input, tt.filter)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMap(t *testing.T) {
	double := func(n int) int { return n * 2 }

	tests := []struct {
		name     string
		input    []int
		mapper   func(int) int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			mapper:   double,
			expected: []int{},
		},
		{
			name:     "double numbers",
			input:    []int{1, 2, 3},
			mapper:   double,
			expected: []int{2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Map(tt.input, tt.mapper)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPartition(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }

	tests := []struct {
		name            string
		input           []int
		partition       func(int) bool
		expectedMatch   []int
		expectedNoMatch []int
	}{
		{
			name:            "empty slice",
			input:           []int{},
			partition:       isEven,
			expectedMatch:   []int{},
			expectedNoMatch: []int{},
		},
		{
			name:            "partition even/odd",
			input:           []int{1, 2, 3, 4, 5, 6},
			partition:       isEven,
			expectedMatch:   []int{2, 4, 6},
			expectedNoMatch: []int{1, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match, noMatch := Partition(tt.input, tt.partition)
			assert.Equal(t, tt.expectedMatch, match)
			assert.Equal(t, tt.expectedNoMatch, noMatch)
		})
	}
}

func TestReduce(t *testing.T) {
	sum := func(acc, curr int) int { return acc + curr }

	tests := []struct {
		name     string
		input    []int
		reducer  func(int, int) int
		init     int
		expected int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			reducer:  sum,
			init:     0,
			expected: 0,
		},
		{
			name:     "sum numbers",
			input:    []int{1, 2, 3, 4, 5},
			reducer:  sum,
			init:     0,
			expected: 15,
		},
		{
			name:     "sum with non-zero init",
			input:    []int{1, 2, 3},
			reducer:  sum,
			init:     10,
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reduce(tt.input, tt.reducer, tt.init)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFind(t *testing.T) {
	isThree := func(n int) bool { return n == 3 }

	tests := []struct {
		name          string
		input         []int
		finder        func(int) bool
		expectedIndex int
		expectedValue int
	}{
		{
			name:          "empty slice",
			input:         []int{},
			finder:        isThree,
			expectedIndex: -1,
			expectedValue: 0,
		},
		{
			name:          "value found",
			input:         []int{1, 2, 3, 4, 5},
			finder:        isThree,
			expectedIndex: 2,
			expectedValue: 3,
		},
		{
			name:          "value not found",
			input:         []int{1, 2, 4, 5},
			finder:        isThree,
			expectedIndex: -1,
			expectedValue: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, value := Find(tt.input, tt.finder)
			assert.Equal(t, tt.expectedIndex, index)
			assert.Equal(t, tt.expectedValue, value)
		})
	}
}
