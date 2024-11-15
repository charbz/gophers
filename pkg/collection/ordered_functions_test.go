package collection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorresponds(t *testing.T) {
	isInverse := func(i, j int) bool { return i == -j }
	tests := []struct {
		name       string
		A          []int
		B          []int
		correspond bool
	}{
		{name: "correspond", A: []int{1, 2, 3}, B: []int{-1, -2, -3}, correspond: true},
		{name: "not correspond", A: []int{1, 2, 3}, B: []int{-1, -2}, correspond: false},
		{name: "not correspond", A: []int{1, 2, 3}, B: []int{-1, -2, -3, -4}, correspond: false},
		{name: "not correspond", A: []int{1, 2, -3}, B: []int{-1, -2, -3}, correspond: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.correspond, Corresponds(NewMockOrderedCollection(tt.A), NewMockOrderedCollection(tt.B), isInverse))
		})
	}
}

func TestDrop(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		n     int
		want  []int
	}{
		{
			name:  "drop first 2 elements",
			slice: []int{1, 2, 3, 4, 5, 6},
			n:     2,
			want:  []int{3, 4, 5, 6},
		},
		{
			name:  "drop 0 elements",
			slice: []int{1, 2, 3},
			n:     0,
			want:  []int{1, 2, 3},
		},
		{
			name:  "drop all elements",
			slice: []int{1, 2, 3},
			n:     3,
			want:  nil,
		},
		{
			name:  "drop more than length",
			slice: []int{1, 2, 3},
			n:     5,
			want:  nil,
		},
		{
			name:  "drop from empty slice",
			slice: []int{},
			n:     2,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewMockOrderedCollection(tt.slice)
			got := Drop(c, tt.n)
			assert.Equal(t, NewMockOrderedCollection(tt.want), got)
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
			index, value := Find(NewMockOrderedCollection(tt.input), tt.finder)
			assert.Equal(t, tt.expectedIndex, index)
			assert.Equal(t, tt.expectedValue, value)
		})
	}
}

func TestDropRight(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		n     int
		want  []int
	}{
		{
			name:  "drop last 2 elements",
			slice: []int{1, 2, 3, 4, 5, 6},
			n:     2,
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "drop 0 elements",
			slice: []int{1, 2, 3},
			n:     0,
			want:  []int{1, 2, 3},
		},
		{
			name:  "drop all elements",
			slice: []int{1, 2, 3},
			n:     3,
			want:  nil,
		},
		{
			name:  "drop more than length",
			slice: []int{1, 2, 3},
			n:     5,
			want:  nil,
		},
		{
			name:  "drop from empty slice",
			slice: []int{},
			n:     2,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewMockOrderedCollection(tt.slice)
			got := DropRight(c, tt.n)
			assert.Equal(t, NewMockOrderedCollection(tt.want), got)
		})
	}
}

func TestDropWhile(t *testing.T) {
	isLessThan4 := func(n int) bool { return n < 4 }
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "drop while less than 4",
			input: []int{1, 2, 3, 4, 5, 6},
			want:  []int{4, 5, 6},
		},
		{
			name:  "drop none",
			input: []int{4, 5, 6},
			want:  []int{4, 5, 6},
		},
		{
			name:  "drop all",
			input: []int{1, 2, 3},
			want:  []int{},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DropWhile(NewMockOrderedCollection(tt.input), isLessThan4)
			assert.Equal(t, NewMockOrderedCollection(tt.want), got)
		})
	}
}

func TestFindLast(t *testing.T) {
	isLessThan6 := func(n int) bool { return n < 6 }
	tests := []struct {
		name          string
		input         []int
		expectedIndex int
		expectedValue int
	}{
		{
			name:          "find last less than 6",
			input:         []int{1, 2, 3, 4, 5, 6},
			expectedIndex: 4,
			expectedValue: 5,
		},
		{
			name:          "no matches",
			input:         []int{6, 7, 8},
			expectedIndex: -1,
			expectedValue: 0,
		},
		{
			name:          "empty slice",
			input:         []int{},
			expectedIndex: -1,
			expectedValue: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, value := FindLast(NewMockOrderedCollection(tt.input), isLessThan6)
			assert.Equal(t, tt.expectedIndex, index)
			assert.Equal(t, tt.expectedValue, value)
		})
	}
}

func TestHead(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		expectedValue int
		expectedErr   error
	}{
		{
			name:          "get head",
			input:         []int{1, 2, 3},
			expectedValue: 1,
			expectedErr:   nil,
		},
		{
			name:          "empty slice",
			input:         []int{},
			expectedValue: 0,
			expectedErr:   EmptyCollectionError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := Head(NewMockOrderedCollection(tt.input))
			assert.Equal(t, tt.expectedValue, value)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestInit(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "get init",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "single element",
			input: []int{1},
			want:  []int{},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Init(NewMockOrderedCollection(tt.input))
			assert.Equal(t, NewMockOrderedCollection(tt.want), got)
		})
	}
}

func TestLast(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		expectedValue int
		expectedErr   error
	}{
		{
			name:          "get last",
			input:         []int{1, 2, 3},
			expectedValue: 3,
			expectedErr:   nil,
		},
		{
			name:          "empty slice",
			input:         []int{},
			expectedValue: 0,
			expectedErr:   EmptyCollectionError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := Last(NewMockOrderedCollection(tt.input))
			assert.Equal(t, tt.expectedValue, value)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "reverse slice",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{5, 4, 3, 2, 1},
		},
		{
			name:  "single element",
			input: []int{1},
			want:  []int{1},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(NewMockOrderedCollection(tt.input))
			assert.Equal(t, NewMockOrderedCollection(tt.want), got)
		})
	}
}

func TestReduceRight(t *testing.T) {
	concat := func(acc string, curr int) string { return acc + fmt.Sprint(curr) }

	tests := []struct {
		name     string
		input    []int
		init     string
		expected string
	}{
		{
			name:     "concat numbers",
			input:    []int{1, 2, 3},
			init:     "",
			expected: "321",
		},
		{
			name:     "empty slice",
			input:    []int{},
			init:     "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReduceRight(NewMockOrderedCollection(tt.input), concat, tt.init)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestReverseMap(t *testing.T) {
	double := func(n int) int { return n * 2 }
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "double numbers in reverse",
			input: []int{1, 2, 3},
			want:  []int{6, 4, 2},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseMap(NewMockOrderedCollection(tt.input), double)
			assert.Equal(t, NewMockOrderedCollection(tt.want), got)
		})
	}
}

func TestTail(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "get tail",
			input: []int{1, 2, 3, 4, 5},
			want:  []int{2, 3, 4, 5},
		},
		{
			name:  "single element",
			input: []int{1},
			want:  []int{},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tail(NewMockOrderedCollection(tt.input))
			assert.Equal(t, NewMockOrderedCollection(tt.want), got)
		})
	}
}

func TestTake(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		n     int
		want  []int
	}{
		{
			name:  "take first 3",
			input: []int{1, 2, 3, 4, 5},
			n:     3,
			want:  []int{1, 2, 3},
		},
		{
			name:  "take more than length",
			input: []int{1, 2, 3},
			n:     5,
			want:  []int{1, 2, 3},
		},
		{
			name:  "take 0",
			input: []int{1, 2, 3},
			n:     0,
			want:  nil,
		},
		{
			name:  "empty slice",
			input: []int{},
			n:     3,
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Take(NewMockOrderedCollection(tt.input), tt.n)
			assert.Equal(t, NewMockOrderedCollection(tt.want), got)
		})
	}
}

func TestTakeRight(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		n     int
		want  []int
	}{
		{
			name:  "take last 3",
			input: []int{1, 2, 3, 4, 5},
			n:     3,
			want:  []int{3, 4, 5},
		},
		{
			name:  "take more than length",
			input: []int{1, 2, 3},
			n:     5,
			want:  []int{1, 2, 3},
		},
		{
			name:  "take 0",
			input: []int{1, 2, 3},
			n:     0,
			want:  nil,
		},
		{
			name:  "empty slice",
			input: []int{},
			n:     3,
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TakeRight(NewMockOrderedCollection(tt.input), tt.n)
			assert.Equal(t, NewMockOrderedCollection(tt.want), got)
		})
	}
}
