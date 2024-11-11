package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			result := Reduce(NewMockCollection(tt.input), tt.reducer, tt.init)
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
			index, value := Find(NewMockCollection(tt.input), tt.finder)
			assert.Equal(t, tt.expectedIndex, index)
			assert.Equal(t, tt.expectedValue, value)
		})
	}
}
