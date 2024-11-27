package list

import (
	"reflect"
	"slices"
	"testing"
)

func TestList_Head(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		want    int
		wantErr bool
	}{
		{
			name:    "non-empty list",
			slice:   []int{1, 2, 3},
			want:    1,
			wantErr: false,
		},
		{
			name:    "empty list",
			slice:   []int{},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got, err := l.Head()
			if tt.wantErr {
				if err == nil {
					t.Errorf("Head() = %v, want error", got)
				}
			} else {
				if err != nil {
					t.Errorf("Head() = %v, want no error", got)
				}
				if got != tt.want {
					t.Errorf("Head() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestList_Drop(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		n     int
		want  []int
	}{
		{
			name:  "drop positive number",
			slice: []int{1, 2, 3, 4, 5},
			n:     2,
			want:  []int{3, 4, 5},
		},
		{
			name:  "drop zero",
			slice: []int{1, 2, 3},
			n:     0,
			want:  []int{1, 2, 3},
		},
		{
			name:  "drop all elements",
			slice: []int{1, 2, 3},
			n:     3,
			want:  []int{},
		},
		{
			name:  "drop more than length",
			slice: []int{1, 2, 3},
			n:     5,
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.Drop(tt.n)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Drop() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_DropRight(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		n     int
		want  []int
	}{
		{
			name:  "drop right positive number",
			slice: []int{1, 2, 3, 4, 5},
			n:     2,
			want:  []int{1, 2, 3},
		},
		{
			name:  "drop right zero",
			slice: []int{1, 2, 3},
			n:     0,
			want:  []int{1, 2, 3},
		},
		{
			name:  "drop right all elements",
			slice: []int{1, 2, 3},
			n:     3,
			want:  []int{},
		},
		{
			name:  "drop right more than length",
			slice: []int{1, 2, 3},
			n:     5,
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.DropRight(tt.n)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("DropRight() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_Contains(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      bool
	}{
		{
			name:      "contains element matching predicate",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return i == 2 },
			want:      true,
		},
		{
			name:      "does not contain element matching predicate",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return i == 4 },
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.Contains(tt.predicate)
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Find(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		value     int
		index     int
	}{
		{
			name:      "find existing element",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i > 3 },
			value:     4,
			index:     3,
		},
		{
			name:      "element not found",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return i > 5 },
			value:     0,
			index:     -1,
		},
		{
			name:      "empty list",
			slice:     []int{},
			predicate: func(i int) bool { return true },
			value:     0,
			index:     -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			index, value := l.Find(tt.predicate)

			if index != tt.index {
				t.Errorf("Find() index = %v, want %v", index, tt.index)
			}
			if value != tt.value {
				t.Errorf("Find() value = %v, want %v", value, tt.value)
			}
		})
	}
}

func TestList_Length(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{
			name:  "non-empty list",
			slice: []int{1, 2, 3},
			want:  3,
		},
		{
			name:  "empty list",
			slice: []int{},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.Length()
			if got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		base     []int
		toConcat [][]int
		want     []int
	}{
		{
			name:     "single list concat",
			base:     []int{1, 2},
			toConcat: [][]int{{3, 4}},
			want:     []int{1, 2, 3, 4},
		},
		{
			name:     "multiple lists concat",
			base:     []int{1, 2},
			toConcat: [][]int{{3, 4}, {5, 6}},
			want:     []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.base)
			var lists []*List[int]
			for _, slice := range tt.toConcat {
				lists = append(lists, NewList(slice))
			}

			result := l.Concat(lists...)
			if !slices.Equal(result.ToSlice(), tt.want) {
				t.Errorf("Concat() = %v, want %v", result.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_Distinct(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{
			name:  "distinct elements",
			slice: []int{1, 2, 2, 3, 3, 4},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "all unique elements",
			slice: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "empty list",
			slice: []int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.Distinct(func(a, b int) bool { return a == b })
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Distinct() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_DropWhile(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "drop while less than 3",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i < 3 },
			want:      []int{3, 4, 5},
		},
		{
			name:      "drop while false",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return false },
			want:      []int{1, 2, 3},
		},
		{
			name:      "drop all elements",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return true },
			want:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.DropWhile(tt.predicate)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("DropWhile() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_Equals(t *testing.T) {
	tests := []struct {
		name   string
		list1  []int
		list2  []int
		equals func(int, int) bool
		want   bool
	}{
		{
			name:   "equal lists",
			list1:  []int{1, 2, 3},
			list2:  []int{1, 2, 3},
			equals: func(a, b int) bool { return a == b },
			want:   true,
		},
		{
			name:   "different lengths",
			list1:  []int{1, 2},
			list2:  []int{1, 2, 3},
			equals: func(a, b int) bool { return a == b },
			want:   false,
		},
		{
			name:   "different elements",
			list1:  []int{1, 2, 3},
			list2:  []int{1, 2, 4},
			equals: func(a, b int) bool { return a == b },
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := NewList(tt.list1)
			l2 := NewList(tt.list2)
			got := l1.Equals(l2, tt.equals)
			if got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Filter(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "filter even numbers",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{2, 4},
		},
		{
			name:      "filter none",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return false },
			want:      []int{},
		},
		{
			name:      "filter all",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return true },
			want:      []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.Filter(tt.predicate)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Filter() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_FilterNot(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "filter not even numbers",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{1, 3, 5},
		},
		{
			name:      "filter not none",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return false },
			want:      []int{1, 2, 3},
		},
		{
			name:      "filter not all",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return true },
			want:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.FilterNot(tt.predicate)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("FilterNot() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_FindLast(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		wantIndex int
		wantValue int
	}{
		{
			name:      "find last even number",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i%2 == 0 },
			wantIndex: 3,
			wantValue: 4,
		},
		{
			name:      "no match",
			slice:     []int{1, 3, 5},
			predicate: func(i int) bool { return i%2 == 0 },
			wantIndex: -1,
			wantValue: 0,
		},
		{
			name:      "empty list",
			slice:     []int{},
			predicate: func(i int) bool { return true },
			wantIndex: -1,
			wantValue: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			gotIndex, gotValue := l.FindLast(tt.predicate)
			if gotIndex != tt.wantIndex {
				t.Errorf("FindLast() index = %v, want %v", gotIndex, tt.wantIndex)
			}
			if gotValue != tt.wantValue {
				t.Errorf("FindLast() value = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestList_Init(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{
			name:  "non-empty list",
			slice: []int{1, 2, 3, 4},
			want:  []int{1, 2, 3},
		},
		{
			name:  "single element",
			slice: []int{1},
			want:  []int{},
		},
		{
			name:  "empty list",
			slice: []int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.Init()
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Init() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_Last(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		want    int
		wantErr bool
	}{
		{
			name:    "non-empty list",
			slice:   []int{1, 2, 3},
			want:    3,
			wantErr: false,
		},
		{
			name:    "empty list",
			slice:   []int{},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got, err := l.Last()
			if tt.wantErr {
				if err == nil {
					t.Errorf("Last() error = nil, want error")
				}
			} else {
				if err != nil {
					t.Errorf("Last() error = %v, want no error", err)
				}
				if got != tt.want {
					t.Errorf("Last() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestList_Partition(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		wantLeft  []int
		wantRight []int
	}{
		{
			name:      "partition even and odd",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i%2 == 0 },
			wantLeft:  []int{2, 4},
			wantRight: []int{1, 3, 5},
		},
		{
			name:      "all elements satisfy predicate",
			slice:     []int{2, 4, 6},
			predicate: func(i int) bool { return i%2 == 0 },
			wantLeft:  []int{2, 4, 6},
			wantRight: []int{},
		},
		{
			name:      "no elements satisfy predicate",
			slice:     []int{1, 3, 5},
			predicate: func(i int) bool { return i%2 == 0 },
			wantLeft:  []int{},
			wantRight: []int{1, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			left, right := l.Partition(tt.predicate)
			if !slices.Equal(left.ToSlice(), tt.wantLeft) {
				t.Errorf("Partition() left = %v, want %v", left.ToSlice(), tt.wantLeft)
			}
			if !slices.Equal(right.ToSlice(), tt.wantRight) {
				t.Errorf("Partition() right = %v, want %v", right.ToSlice(), tt.wantRight)
			}
		})
	}
}

func TestList_Reverse(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{
			name:  "non-empty list",
			slice: []int{1, 2, 3, 4},
			want:  []int{4, 3, 2, 1},
		},
		{
			name:  "single element",
			slice: []int{1},
			want:  []int{1},
		},
		{
			name:  "empty list",
			slice: []int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.Reverse()
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Reverse() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_Reject(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "filter not even numbers",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{1, 3, 5},
		},
		{
			name:      "filter not none",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return false },
			want:      []int{1, 2, 3},
		},
		{
			name:      "filter not all",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return true },
			want:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.Reject(tt.predicate)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Reject() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_SplitAt(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		n         int
		wantLeft  []int
		wantRight []int
	}{
		{
			name:      "split in middle",
			slice:     []int{1, 2, 3, 4, 5},
			n:         2,
			wantLeft:  []int{1, 2, 3},
			wantRight: []int{4, 5},
		},
		{
			name:      "split at start",
			slice:     []int{1, 2, 3},
			n:         0,
			wantLeft:  []int{1},
			wantRight: []int{2, 3},
		},
		{
			name:      "split at end",
			slice:     []int{1, 2, 3},
			n:         3,
			wantLeft:  []int{1, 2, 3},
			wantRight: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			left, right := l.SplitAt(tt.n)
			if !slices.Equal(left.ToSlice(), tt.wantLeft) {
				t.Errorf("SplitAt() left = %v, want %v", left.ToSlice(), tt.wantLeft)
			}
			if !slices.Equal(right.ToSlice(), tt.wantRight) {
				t.Errorf("SplitAt() right = %v, want %v", right.ToSlice(), tt.wantRight)
			}
		})
	}
}

func TestList_Take(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		n     int
		want  []int
	}{
		{
			name:  "take positive number",
			slice: []int{1, 2, 3, 4, 5},
			n:     3,
			want:  []int{1, 2, 3},
		},
		{
			name:  "take zero",
			slice: []int{1, 2, 3},
			n:     0,
			want:  []int{},
		},
		{
			name:  "take more than length",
			slice: []int{1, 2, 3},
			n:     5,
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.Take(tt.n)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Take() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_TakeRight(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		n     int
		want  []int
	}{
		{
			name:  "take right positive number",
			slice: []int{1, 2, 3, 4, 5},
			n:     3,
			want:  []int{3, 4, 5},
		},
		{
			name:  "take right zero",
			slice: []int{1, 2, 3},
			n:     0,
			want:  []int{},
		},
		{
			name:  "take right more than length",
			slice: []int{1, 2, 3},
			n:     5,
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.TakeRight(tt.n)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("TakeRight() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_Tail(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{
			name:  "non-empty list",
			slice: []int{1, 2, 3, 4},
			want:  []int{2, 3, 4},
		},
		{
			name:  "single element",
			slice: []int{1},
			want:  []int{},
		},
		{
			name:  "empty list",
			slice: []int{},
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewList(tt.slice)
			got := l.Tail()
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Tail() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestList_Shuffle(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{name: "basic shuffle", input: []int{1, 2, 3, 4, 5}},
		{name: "empty list", input: []int{}},
		{name: "single element", input: []int{22}},
		{name: "duplicate elements", input: []int{1, 1, 2, 2, 3, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := NewList(tt.input)
			shuffled := list.Shuffle()

			// Verify length
			if shuffled.Length() != list.Length() {
				t.Errorf("Shuffle() length = %d, want %d", shuffled.Length(), list.Length())
			}

			// Verify element preservation
			originalMap := make(map[int]int)
			shuffledMap := make(map[int]int)
			for _, v := range list.ToSlice() {
				originalMap[v]++
			}
			for _, v := range shuffled.ToSlice() {
				shuffledMap[v]++
			}
			if !reflect.DeepEqual(originalMap, shuffledMap) {
				t.Errorf("Shuffle() elements mismatch, got %v, want %v", shuffledMap, originalMap)
			}
		})
	}
}

func TestList_ShuffleRandomization(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	list := NewList(input)
	iterations := 10000

	sameOrderCount := 0
	for i := 0; i < iterations; i++ {
		shuffled := list.Shuffle()
		if reflect.DeepEqual(input, shuffled.ToSlice()) {
			sameOrderCount++
		}
	}

	// Expect the same order to appear <5% of the time for small lists
	threshold := 0.05 * float64(iterations)
	if float64(sameOrderCount) > threshold {
		t.Errorf("Shuffle() produced the same order %d times, exceeding threshold %f", sameOrderCount, threshold)
	}
}

func TestList_ShuffleDistribution(t *testing.T) {
	input := []int{1, 2, 3, 4}
	list := NewList(input)

	// Map to track the position of each value
	positionCounts := make([]map[int]int, len(input))
	for i := range positionCounts {
		positionCounts[i] = make(map[int]int)
	}

	iterations := 10000
	for i := 0; i < iterations; i++ {
		shuffled := list.Shuffle()
		for pos, val := range shuffled.ToSlice() {
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
