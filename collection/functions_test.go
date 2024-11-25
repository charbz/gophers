package collection

import (
	"slices"
	"testing"
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
			if got := Count(NewMockCollection(tt.input), countEvens); got != tt.count {
				t.Errorf("Count() = %v, want %v", got, tt.count)
			}
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
		{name: "diff with same elements", A: []int{1, 2, 3, 4, 3, 6}, B: []int{1, 2, 3, 4, 5, 6}, diff: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Diff(NewMockCollection(tt.A), NewMockCollection(tt.B)).(*MockCollection[int]).items
			want := NewMockCollection(tt.diff).items
			if !slices.Equal(got, want) {
				t.Errorf("Diff() = %v, want %v", got, want)
			}
		})
	}
}

func TestDiffFunc(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		b    []int
		diff []int
	}{
		{name: "diff", a: []int{1, 2, 3, 4, 5, 6}, b: []int{2, 4, 6, 8, 10, 12}, diff: []int{1, 3, 5}},
		{name: "diff with empty b", a: []int{1, 2, 3, 4, 5, 6}, b: []int{}, diff: []int{1, 2, 3, 4, 5, 6}},
		{name: "diff with empty a", a: []int{}, b: []int{1, 2, 3, 4, 5, 6}, diff: nil},
		{name: "diff with same elements", a: []int{1, 2, 3, 4, 3, 6}, b: []int{1, 2, 3, 4, 5, 6}, diff: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DiffFunc(NewMockCollection(tt.a), NewMockCollection(tt.b), func(a, b int) bool { return a == b }).(*MockCollection[int]).items
			want := NewMockCollection(tt.diff).items
			if !slices.Equal(got, want) {
				t.Errorf("DiffFunc() = %v, want %v", got, want)
			}
		})
	}
}

func TestDistinct(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		want []int
	}{
		{name: "distinct", a: []int{1, 1, 1, 2, 2, 3}, want: []int{1, 2, 3}},
		{name: "distinct with no duplicates", a: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{name: "distinct with empty collection", a: []int{}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Distinct(NewMockCollection(tt.a), func(a, b int) bool { return a == b }).(*MockCollection[int]).items
			if !slices.Equal(got, tt.want) {
				t.Errorf("DistinctFunc() = %v, want %v", got, tt.want)
			}
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
			if got := Reduce(NewMockCollection(tt.input), tt.reducer, tt.init); got != tt.expected {
				t.Errorf("Reduce() = %v, want %v", got, tt.expected)
			}
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
			want := NewMockCollection(tt.want)
			if !slices.Equal(got.(*MockCollection[int]).items, want.items) {
				t.Errorf("Filter() = %v, want %v", got, want)
			}
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
			want := NewMockCollection(tt.want)
			if !slices.Equal(got.(*MockCollection[int]).items, want.items) {
				t.Errorf("FilterNot() = %v, want %v", got, want)
			}
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
			if got := ForAll(NewMockCollection(tt.input), isLessThan10); got != tt.expected {
				t.Errorf("ForAll() = %v, want %v", got, tt.expected)
			}
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
				want := NewMockCollection(v)
				got := result[k]
				if !slices.Equal(got.(*MockCollection[int]).items, want.items) {
					t.Errorf("GroupBy()[%v] = %v, want %v", k, got, want)
				}
			}
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
			want := NewMockCollection(tt.want)
			if !slices.Equal(got.(*MockCollection[int]).items, want.items) {
				t.Errorf("Intersect() = %v, want %v", got, want)
			}
		})
	}
}

func TestIntersectFunc(t *testing.T) {
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
			got := IntersectFunc(NewMockCollection(tt.a), NewMockCollection(tt.b), func(a, b int) bool { return a == b })
			want := NewMockCollection(tt.want)
			if !slices.Equal(got.(*MockCollection[int]).items, want.items) {
				t.Errorf("IntersectFunc() = %v, want %v", got, want)
			}
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
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(NewMockCollection(tt.input), double)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
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
			if value != tt.expectedValue {
				t.Errorf("MaxBy() value = %v, want %v", value, tt.expectedValue)
			}
			if err != tt.expectedErr {
				t.Errorf("MaxBy() error = %v, want %v", err, tt.expectedErr)
			}
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
			if value != tt.expectedValue {
				t.Errorf("MinBy() value = %v, want %v", value, tt.expectedValue)
			}
			if err != tt.expectedErr {
				t.Errorf("MinBy() error = %v, want %v", err, tt.expectedErr)
			}
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
			wantMatch := NewMockCollection(tt.wantMatch)
			wantRest := NewMockCollection(tt.wantRest)

			if len(match.(*MockCollection[int]).items) != len(wantMatch.items) {
				t.Errorf("Partition() match = %v, want %v", match, wantMatch)
			}
			for i := range wantMatch.items {
				if match.(*MockCollection[int]).items[i] != wantMatch.items[i] {
					t.Errorf("Partition() match = %v, want %v", match, wantMatch)
				}
			}

			if len(rest.(*MockCollection[int]).items) != len(wantRest.items) {
				t.Errorf("Partition() rest = %v, want %v", rest, wantRest)
			}
			for i := range wantRest.items {
				if rest.(*MockCollection[int]).items[i] != wantRest.items[i] {
					t.Errorf("Partition() rest = %v, want %v", rest, wantRest)
				}
			}
		})
	}
}
