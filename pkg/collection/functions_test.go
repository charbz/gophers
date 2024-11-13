package collection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	countEvens := func(n int) bool { return n%2 == 0 }
	tests := []struct {
		name  string
		input []int
		count int
	}{
		{name: "count evens", input: []int{1, 2, 3, 4, 5, 6}, count: 3},
		{name: "count evens", input: []int{1, 3, 5}, count: 0},
		{name: "count evens", input: []int{2}, count: 1},
		{name: "count evens", input: []int{}, count: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.count, Count(NewMockCollection(tt.input), countEvens))
		})
	}
}

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
			assert.Equal(t, tt.correspond, Corresponds(NewMockCollection(tt.A), NewMockCollection(tt.B), isInverse))
		})
	}
}

func TestDiff(t *testing.T) {
	tests := []struct {
		name string
		A    []int
		B    []int
		diff []int
	}{
		{name: "diff", A: []int{1, 2, 3, 4, 5, 6}, B: []int{2, 4, 6, 8, 10, 12}, diff: []int{1, 3, 5}},
		{name: "diff with empty B", A: []int{1, 2, 3, 4, 5, 6}, B: []int{}, diff: []int{1, 2, 3, 4, 5, 6}},
		{name: "diff with empty A", A: []int{}, B: []int{1, 2, 3, 4, 5, 6}, diff: nil},
		{name: "diff with same elements", A: []int{1, 2, 3, 4, 5, 6}, B: []int{1, 2, 3, 4, 5, 6}, diff: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, NewMockCollection(tt.diff), Diff(NewMockCollection(tt.A), NewMockCollection(tt.B)))
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
			c := NewMockCollection(tt.slice)
			got := Drop(c, tt.n)
			assert.Equal(t, NewMockCollection(tt.want), got)
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
			c := NewMockCollection(tt.slice)
			got := DropRight(c, tt.n)
			assert.Equal(t, NewMockCollection(tt.want), got)
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
			got := DropWhile(NewMockCollection(tt.input), isLessThan4)
			assert.Equal(t, NewMockCollection(tt.want), got)
		})
	}
}

func TestFilter(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "filter evens",
			input: []int{1, 2, 3, 4, 5, 6},
			want:  []int{2, 4, 6},
		},
		{
			name:  "no matches",
			input: []int{1, 3, 5},
			want:  nil,
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(NewMockCollection(tt.input), isEven)
			assert.Equal(t, NewMockCollection(tt.want), got)
		})
	}
}

func TestFilterNot(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "filter not evens",
			input: []int{1, 2, 3, 4, 5, 6},
			want:  []int{1, 3, 5},
		},
		{
			name:  "no matches",
			input: []int{2, 4, 6},
			want:  nil,
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterNot(NewMockCollection(tt.input), isEven)
			assert.Equal(t, NewMockCollection(tt.want), got)
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
			index, value := FindLast(NewMockCollection(tt.input), isLessThan6)
			assert.Equal(t, tt.expectedIndex, index)
			assert.Equal(t, tt.expectedValue, value)
		})
	}
}

func TestForAll(t *testing.T) {
	isLessThan10 := func(n int) bool { return n < 10 }
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "all less than 10",
			input:    []int{1, 2, 3, 4, 5},
			expected: true,
		},
		{
			name:     "not all less than 10",
			input:    []int{1, 5, 10, 15},
			expected: false,
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ForAll(NewMockCollection(tt.input), isLessThan10)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestForEach(t *testing.T) {
	sum := 0
	addToSum := func(n int) { sum += n }

	tests := []struct {
		name        string
		input       []int
		expectedSum int
	}{
		{
			name:        "sum numbers",
			input:       []int{1, 2, 3, 4, 5},
			expectedSum: 15,
		},
		{
			name:        "empty slice",
			input:       []int{},
			expectedSum: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum = 0 // Reset sum for each test
			ForEach(NewMockCollection(tt.input), addToSum)
			assert.Equal(t, tt.expectedSum, sum)
		})
	}
}

func TestGroupBy(t *testing.T) {
	modTwo := func(n int) int { return n % 2 }
	tests := []struct {
		name     string
		input    []int
		expected map[int][]int
	}{
		{
			name:  "group by mod 2",
			input: []int{1, 2, 3, 4, 5, 6},
			expected: map[int][]int{
				0: {2, 4, 6},
				1: {1, 3, 5},
			},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: map[int][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GroupBy(NewMockCollection(tt.input), modTwo)
			for k, v := range tt.expected {
				assert.Equal(t, NewMockCollection(v), result[k])
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
			value, err := Head(NewMockCollection(tt.input))
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
			got := Init(NewMockCollection(tt.input))
			assert.Equal(t, NewMockCollection(tt.want), got)
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		want []int
	}{
		{
			name: "intersect",
			a:    []int{1, 2, 3, 4, 5, 6},
			b:    []int{2, 4, 6, 8, 10},
			want: []int{2, 4, 6},
		},
		{
			name: "no intersection",
			a:    []int{1, 3, 5},
			b:    []int{2, 4, 6},
			want: nil,
		},
		{
			name: "empty slices",
			a:    []int{},
			b:    []int{},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Intersect(NewMockCollection(tt.a), NewMockCollection(tt.b))
			assert.Equal(t, NewMockCollection(tt.want), got)
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
			value, err := Last(NewMockCollection(tt.input))
			assert.Equal(t, tt.expectedValue, value)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestMap(t *testing.T) {
	double := func(n int) int { return n * 2 }
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "double numbers",
			input: []int{1, 2, 3},
			want:  []int{2, 4, 6},
		},
		{
			name:  "empty slice",
			input: []int{},
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(NewMockCollection(tt.input), double)
			assert.Equal(t, NewMockCollection(tt.want), got)
		})
	}
}

func TestMaxBy(t *testing.T) {
	identity := func(a int) int { return a }
	tests := []struct {
		name          string
		input         []int
		expectedValue int
		expectedErr   error
	}{
		{
			name:          "find max",
			input:         []int{1, 2, 3, 4, 5},
			expectedValue: 5,
			expectedErr:   nil,
		},
		{
			name:          "find max in middle",
			input:         []int{1, 2, 5, 3, 4},
			expectedValue: 5,
			expectedErr:   nil,
		},
		{
			name:          "empty collection",
			input:         []int{},
			expectedValue: 0,
			expectedErr:   EmptyCollectionError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := MaxBy(NewMockCollection(tt.input), identity)
			assert.Equal(t, tt.expectedValue, value)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestMinBy(t *testing.T) {
	identity := func(a int) int { return a }
	tests := []struct {
		name          string
		input         []int
		expectedValue int
		expectedErr   error
	}{
		{
			name:          "find min",
			input:         []int{5, 4, 3, 2, 1},
			expectedValue: 1,
			expectedErr:   nil,
		},
		{
			name:          "find min in middle",
			input:         []int{5, 4, 1, 3, 2},
			expectedValue: 1,
			expectedErr:   nil,
		},
		{
			name:          "empty collection",
			input:         []int{},
			expectedValue: 0,
			expectedErr:   EmptyCollectionError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := MinBy(NewMockCollection(tt.input), identity)
			assert.Equal(t, tt.expectedValue, value)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}

func TestPartition(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	tests := []struct {
		name      string
		input     []int
		wantMatch []int
		wantRest  []int
	}{
		{
			name:      "partition evens",
			input:     []int{1, 2, 3, 4, 5, 6},
			wantMatch: []int{2, 4, 6},
			wantRest:  []int{1, 3, 5},
		},
		{
			name:      "empty slice",
			input:     []int{},
			wantMatch: nil,
			wantRest:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			match, rest := Partition(NewMockCollection(tt.input), isEven)
			assert.Equal(t, NewMockCollection(tt.wantMatch), match)
			assert.Equal(t, NewMockCollection(tt.wantRest), rest)
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
			got := Reverse(NewMockCollection(tt.input))
			assert.Equal(t, NewMockCollection(tt.want), got)
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
			result := ReduceRight(NewMockCollection(tt.input), concat, tt.init)
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
			got := ReverseMap(NewMockCollection(tt.input), double)
			assert.Equal(t, NewMockCollection(tt.want), got)
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
			got := Tail(NewMockCollection(tt.input))
			assert.Equal(t, NewMockCollection(tt.want), got)
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
			got := Take(NewMockCollection(tt.input), tt.n)
			assert.Equal(t, NewMockCollection(tt.want), got)
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
			got := TakeRight(NewMockCollection(tt.input), tt.n)
			assert.Equal(t, NewMockCollection(tt.want), got)
		})
	}
}
