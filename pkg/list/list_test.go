package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_ForEach(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{
			name:  "sum all elements",
			slice: []int{1, 2, 3, 4, 5},
			want:  15,
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
			sum := 0
			l.ForEach(func(i int) {
				sum += i
			})
			assert.Equal(t, tt.want, sum)
		})
	}
}

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
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.want, got.ToSlice())
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
			assert.Equal(t, tt.want, got.ToSlice())
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
			assert.Equal(t, tt.want, got)
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

			assert.Equal(t, tt.index, index)
			assert.Equal(t, tt.value, value)
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
			assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.want, result.ToSlice())
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
			assert.Equal(t, tt.want, got.ToSlice())
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
			assert.Equal(t, tt.want, got.ToSlice())
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
			assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.want, got.ToSlice())
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
			assert.Equal(t, tt.want, got.ToSlice())
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
			assert.Equal(t, tt.wantIndex, gotIndex)
			assert.Equal(t, tt.wantValue, gotValue)
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
			assert.Equal(t, tt.want, got.ToSlice())
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
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
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
			assert.Equal(t, tt.wantLeft, left.ToSlice())
			assert.Equal(t, tt.wantRight, right.ToSlice())
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
			assert.Equal(t, tt.want, got.ToSlice())
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
			n:         3,
			wantLeft:  []int{1, 2, 3},
			wantRight: []int{4, 5},
		},
		{
			name:      "split at start",
			slice:     []int{1, 2, 3},
			n:         0,
			wantLeft:  []int{},
			wantRight: []int{1, 2, 3},
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
			assert.Equal(t, tt.wantLeft, left.ToSlice())
			assert.Equal(t, tt.wantRight, right.ToSlice())
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
			assert.Equal(t, tt.want, got.ToSlice())
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
			assert.Equal(t, tt.want, got.ToSlice())
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
			assert.Equal(t, tt.want, got.ToSlice())
		})
	}
}
