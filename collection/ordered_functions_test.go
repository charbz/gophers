package collection

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
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
			got := Corresponds(NewMockOrderedCollection(tt.A), NewMockOrderedCollection(tt.B), isInverse)
			if got != tt.correspond {
				t.Errorf("Corresponds() = %v, want %v", got, tt.correspond)
			}
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
			if !slices.Equal(got.(*MockOrderedCollection[int]).items, tt.want) {
				t.Errorf("Drop() = %v, want %v", got, NewMockOrderedCollection(tt.want))
			}
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
			if index != tt.expectedIndex || value != tt.expectedValue {
				t.Errorf("Find() = %v, want %v", index, tt.expectedIndex)
				t.Errorf("Find() = %v, want %v", value, tt.expectedValue)
			}
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
			if !slices.Equal(got.(*MockOrderedCollection[int]).items, tt.want) {
				t.Errorf("DropRight() = %v, want %v", got, NewMockOrderedCollection(tt.want))
			}
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
			if !slices.Equal(got.(*MockOrderedCollection[int]).items, tt.want) {
				t.Errorf("DropWhile() = %v, want %v", got, NewMockOrderedCollection(tt.want))
			}
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
			if index != tt.expectedIndex || value != tt.expectedValue {
				t.Errorf("FindLast() = %v, want %v", index, tt.expectedIndex)
				t.Errorf("FindLast() = %v, want %v", value, tt.expectedValue)
			}
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
			if value != tt.expectedValue || err != tt.expectedErr {
				t.Errorf("Head() = %v, want %v", value, tt.expectedValue)
				t.Errorf("Head() = %v, want %v", err, tt.expectedErr)
			}
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
			if !slices.Equal(got.(*MockOrderedCollection[int]).items, tt.want) {
				t.Errorf("Init() = %v, want %v", got, NewMockOrderedCollection(tt.want))
			}
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
			if value != tt.expectedValue || err != tt.expectedErr {
				t.Errorf("Last() = %v, want %v", value, tt.expectedValue)
				t.Errorf("Last() = %v, want %v", err, tt.expectedErr)
			}
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
			if !slices.Equal(got.(*MockOrderedCollection[int]).items, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, NewMockOrderedCollection(tt.want))
			}
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
			if result != tt.expected {
				t.Errorf("ReduceRight() = %v, want %v", result, tt.expected)
			}
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
			if !slices.Equal(got.(*MockOrderedCollection[int]).items, tt.want) {
				t.Errorf("ReverseMap() = %v, want %v", got, NewMockOrderedCollection(tt.want))
			}
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
			if !slices.Equal(got.(*MockOrderedCollection[int]).items, tt.want) {
				t.Errorf("Tail() = %v, want %v", got, NewMockOrderedCollection(tt.want))
			}
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
			if !slices.Equal(got.(*MockOrderedCollection[int]).items, tt.want) {
				t.Errorf("Take() = %v, want %v", got, NewMockOrderedCollection(tt.want))
			}
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
			if !slices.Equal(got.(*MockOrderedCollection[int]).items, tt.want) {
				t.Errorf("TakeRight() = %v, want %v", got, NewMockOrderedCollection(tt.want))
			}
		})
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		name       string
		A          []int
		B          []int
		startsWith bool
	}{
		// Core test cases
		{name: "exact match", A: []int{1, 2, 3}, B: []int{1, 2, 3}, startsWith: true},
		{name: "prefix match", A: []int{1, 2, 3, 4}, B: []int{1, 2}, startsWith: true},
		{name: "no match", A: []int{1, 2, 3}, B: []int{2, 3}, startsWith: false},
		{name: "B longer than A", A: []int{1, 2}, B: []int{1, 2, 3}, startsWith: false},

		// Edge cases
		{name: "B is empty", A: []int{1, 2, 3}, B: []int{}, startsWith: true},
		{name: "A is empty", A: []int{}, B: []int{1}, startsWith: false},
		{name: "both empty", A: []int{}, B: []int{}, startsWith: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StartsWith(NewMockOrderedCollection(tt.A), NewMockOrderedCollection(tt.B))
			if got != tt.startsWith {
				t.Errorf("StartsWith() = %v, want %v", got, tt.startsWith)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		name     string
		A        []int
		B        []int
		endsWith bool
	}{
		// Core test cases
		{name: "exact match", A: []int{1, 2, 3}, B: []int{1, 2, 3}, endsWith: true},
		{name: "postfix match", A: []int{1, 2, 3, 4}, B: []int{3, 4}, endsWith: true},
		{name: "no match", A: []int{1, 2, 3}, B: []int{1, 2}, endsWith: false},
		{name: "B longer than A", A: []int{1, 2}, B: []int{1, 2, 3}, endsWith: false},

		// Edge cases
		{name: "B is empty", A: []int{1, 2, 3}, B: []int{}, endsWith: true},
		{name: "A is empty", A: []int{}, B: []int{1}, endsWith: false},
		{name: "both empty", A: []int{}, B: []int{}, endsWith: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EndsWith(NewMockOrderedCollection(tt.A), NewMockOrderedCollection(tt.B))
			if got != tt.endsWith {
				t.Errorf("EndsWith() = %v, want %v", got, tt.endsWith)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{name: "basic shuffle", input: []int{1, 2, 3, 4, 5}},
		{name: "empty collection", input: []int{}},
		{name: "single element", input: []int{42}},
		{name: "duplicate elements", input: []int{1, 1, 2, 2, 3, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewMockOrderedCollection(tt.input)
			shuffled := Shuffle(c)

			if shuffled.Length() != c.Length() {
				t.Errorf("Shuffle() length = %d, want %d", shuffled.Length(), c.Length())
			}

			originalMap := make(map[int]int)
			shuffledMap := make(map[int]int)
			for _, v := range c.All() {
				originalMap[v]++
			}
			for _, v := range shuffled.All() {
				shuffledMap[v]++
			}
			if !reflect.DeepEqual(originalMap, shuffledMap) {
				t.Errorf("Shuffle() elements mismatch, got %v, want %v", shuffledMap, originalMap)
			}
		})
	}
}

func TestShuffleRandomization(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	c := NewMockOrderedCollection(input)
	iterations := 10000

	sameOrderCount := 0
	for i := 0; i < iterations; i++ {
		shuffled := Shuffle(c)
		if reflect.DeepEqual(input, shuffled.(*MockOrderedCollection[int]).items) {
			sameOrderCount++
		}
	}

	// Expect the same order to appear <5% of the time for small collections
	threshold := 0.05 * float64(iterations)
	if float64(sameOrderCount) > threshold {
		t.Errorf("Shuffle() produced the same order %d times, exceeding threshold %f", sameOrderCount, threshold)
	}
}

// This test function ensures that Shuffle() returns randomized
// items in a uniform distribution
func TestShuffleDistribution(t *testing.T) {
	input := []int{1, 2, 3, 4}
	c := NewMockOrderedCollection(input)

	// Map to track the position of each value
	positionCounts := make([]map[int]int, len(input))
	for i := range positionCounts {
		positionCounts[i] = make(map[int]int)
	}

	iterations := 10000
	for i := 0; i < iterations; i++ {
		shuffled := Shuffle(c)
		for pos, val := range shuffled.All() {
			positionCounts[pos][val]++
		}
	}

	// Validate uniform distribution
	expectedCount := iterations / len(input)
	tolerance := 0.1 * float64(expectedCount)
	for pos, counts := range positionCounts {
		for val, count := range counts {
			if float64(count) < float64(expectedCount)-tolerance || float64(count) > float64(expectedCount)+tolerance {
				t.Errorf("Value %d appeared at position %d %d times, which is outside the tolerance range [%f, %f]",
					val, pos, count, float64(expectedCount)-tolerance, float64(expectedCount)+tolerance)
			}
		}
	}
}
