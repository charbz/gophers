package sequence

import (
	"slices"
	"testing"
)

func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		base     []int
		toConcat [][]int
		want     []int
	}{
		{
			name:     "single Sequence concat",
			base:     []int{1, 2},
			toConcat: [][]int{{3, 4}},
			want:     []int{1, 2, 3, 4},
		},
		{
			name:     "multiple Sequences concat",
			base:     []int{1, 2},
			toConcat: [][]int{{3, 4}, {5, 6}},
			want:     []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSequence(tt.base)
			var Sequences []Sequence[int]
			for _, slice := range tt.toConcat {
				Sequences = append(Sequences, *NewSequence(slice))
			}

			result := c.Concat(Sequences...)
			if !slices.Equal(result.elements, tt.want) {
				t.Errorf("Concat() = %v, want %v", result.elements, tt.want)
			}
		})
	}
}

func TestDistinct(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "no duplicates",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "with duplicates",
			input: []int{1, 2, 2, 3, 3, 3},
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSequence(tt.input)
			result := c.Distinct(func(a, b int) bool {
				return a == b
			})
			if !slices.Equal(result.elements, tt.want) {
				t.Errorf("Distinct() = %v, want %v", result.elements, tt.want)
			}
		})
	}
}

func TestSequence_At(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		index   int
		want    int
		wantErr bool
	}{
		{
			name:    "valid index - first element",
			slice:   []int{1, 2, 3},
			index:   0,
			want:    1,
			wantErr: false,
		},
		{
			name:    "valid index - middle element",
			slice:   []int{1, 2, 3},
			index:   1,
			want:    2,
			wantErr: false,
		},
		{
			name:    "valid index - last element",
			slice:   []int{1, 2, 3},
			index:   2,
			want:    3,
			wantErr: false,
		},
		{
			name:    "invalid index - out of bounds",
			slice:   []int{1, 2, 3},
			index:   3,
			wantErr: true,
		},
		{
			name:    "invalid index - negative",
			slice:   []int{1, 2, 3},
			index:   -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("At() panicked: %v", r)
					}
				}
			}()
			c := NewSequence(tt.slice)
			got := c.At(tt.index)
			if got != tt.want {
				t.Errorf("At() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestSequence_Contains(t *testing.T) {
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
			c := NewSequence(tt.slice)
			got := c.Contains(tt.predicate)
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSequence_Drop(t *testing.T) {
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
			want:  nil,
		},
		{
			name:  "drop more than length",
			slice: []int{1, 2, 3},
			n:     5,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSequence(tt.slice)
			got := c.Drop(tt.n)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Drop() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestSequence_Filter(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		filter func(int) bool
		want   []int
	}{
		{
			name:   "filter even numbers",
			slice:  []int{1, 2, 3, 4, 5, 6},
			filter: func(i int) bool { return i%2 == 0 },
			want:   []int{2, 4, 6},
		},
		{
			name:   "filter nothing",
			slice:  []int{1, 2, 3},
			filter: func(i int) bool { return false },
			want:   nil,
		},
		{
			name:   "filter everything",
			slice:  []int{1, 2, 3},
			filter: func(i int) bool { return true },
			want:   []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSequence(tt.slice)
			got := c.Filter(tt.filter)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Filter() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestSequence_DropRight(t *testing.T) {
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
			want:  nil,
		},
		{
			name:  "drop right more than length",
			slice: []int{1, 2, 3},
			n:     5,
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSequence(tt.slice)
			got := c.DropRight(tt.n)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("DropRight() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestSequence_FilterNot(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		filter func(int) bool
		want   []int
	}{
		{
			name:   "filter not even numbers",
			slice:  []int{1, 2, 3, 4, 5, 6},
			filter: func(i int) bool { return i%2 == 0 },
			want:   []int{1, 3, 5},
		},
		{
			name:   "filter not nothing",
			slice:  []int{1, 2, 3},
			filter: func(i int) bool { return false },
			want:   []int{1, 2, 3},
		},
		{
			name:   "filter not everything",
			slice:  []int{1, 2, 3},
			filter: func(i int) bool { return true },
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSequence(tt.slice)
			got := c.FilterNot(tt.filter)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("FilterNot() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestSequence_Find(t *testing.T) {
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
			name:      "empty slice",
			slice:     []int{},
			predicate: func(i int) bool { return true },
			value:     0,
			index:     -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSequence(tt.slice)
			index, value := c.Find(tt.predicate)

			if index != tt.index {
				t.Errorf("Find() index = %v, want %v", index, tt.index)
			}
			if value != tt.value {
				t.Errorf("Find() value = %v, want %v", value, tt.value)
			}
		})
	}
}

func TestSequence_Head(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		want    int
		wantErr bool
	}{
		{
			name:    "non-empty slice",
			slice:   []int{1, 2, 3},
			want:    1,
			wantErr: false,
		},
		{
			name:    "empty slice",
			slice:   []int{},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("Head() panicked: %v", r)
					}
				}
			}()
			c := NewSequence(tt.slice)
			got, _ := c.Head()
			if got != tt.want {
				t.Errorf("Head() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestSequence_Pop(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		want    int
		wantErr bool
	}{
		{
			name:    "non-empty slice",
			slice:   []int{1, 2, 3},
			want:    3,
			wantErr: false,
		},
		{
			name:    "empty slice",
			slice:   []int{},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("Pop() panicked: %v", r)
					}
				}
			}()
			c := NewSequence(tt.slice)
			got, _ := c.Pop()
			if !tt.wantErr {
				if got != tt.want {
					t.Errorf("Pop() = %v, want %v", got, tt.want)
				}
				if c.Length() != len(tt.slice)-1 {
					t.Errorf("Pop() length = %v, want %v", c.Length(), len(tt.slice)-1)
				}
			}
		})
	}
}

func TestSequence_Push(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		toPush   int
		expected []int
	}{
		{
			name:     "push to non-empty slice",
			slice:    []int{1, 2},
			toPush:   3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "push to empty slice",
			slice:    []int{},
			toPush:   1,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSequence(tt.slice)
			c.Push(tt.toPush)
			if !slices.Equal(c.ToSlice(), tt.expected) {
				t.Errorf("Push() = %v, want %v", c.ToSlice(), tt.expected)
			}
		})
	}
}

func TestSequence_Dequeue(t *testing.T) {
	tests := []struct {
		name    string
		slice   []int
		want    int
		wantErr bool
	}{
		{
			name:    "non-empty slice",
			slice:   []int{1, 2, 3},
			want:    1,
			wantErr: false,
		},
		{
			name:    "empty slice",
			slice:   []int{},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tt.wantErr {
						t.Errorf("Dequeue() panicked: %v", r)
					}
				}
			}()
			c := NewSequence(tt.slice)
			if !tt.wantErr {
				got, _ := c.Dequeue()
				if got != tt.want {
					t.Errorf("Dequeue() = %v, want %v", got, tt.want)
				}
				if c.Length() != len(tt.slice)-1 {
					t.Errorf("Dequeue() length = %v, want %v", c.Length(), len(tt.slice)-1)
				}
			}
		})
	}
}

func TestSequence_Enqueue(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		toEnqueue int
		expected  []int
	}{
		{
			name:      "enqueue to non-empty slice",
			slice:     []int{1, 2},
			toEnqueue: 3,
			expected:  []int{1, 2, 3},
		},
		{
			name:      "enqueue to empty slice",
			slice:     []int{},
			toEnqueue: 1,
			expected:  []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSequence(tt.slice)
			c.Enqueue(tt.toEnqueue)
			if !slices.Equal(c.ToSlice(), tt.expected) {
				t.Errorf("Enqueue() = %v, want %v", c.ToSlice(), tt.expected)
			}
		})
	}
}

func TestSequence_Length(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{
			name:  "non-empty slice",
			slice: []int{1, 2, 3},
			want:  3,
		},
		{
			name:  "empty slice",
			slice: []int{},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSequence(tt.slice)
			got := c.Length()
			if got != tt.want {
				t.Errorf("Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSequence_Reject(t *testing.T) {
	tests := []struct {
		name   string
		slice  []int
		filter func(int) bool
		want   []int
	}{
		{
			name:   "filter not even numbers",
			slice:  []int{1, 2, 3, 4, 5, 6},
			filter: func(i int) bool { return i%2 == 0 },
			want:   []int{1, 3, 5},
		},
		{
			name:   "filter not nothing",
			slice:  []int{1, 2, 3},
			filter: func(i int) bool { return false },
			want:   []int{1, 2, 3},
		},
		{
			name:   "filter not everything",
			slice:  []int{1, 2, 3},
			filter: func(i int) bool { return true },
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSequence(tt.slice)
			got := c.Reject(tt.filter)
			if !slices.Equal(got.ToSlice(), tt.want) {
				t.Errorf("Reject() = %v, want %v", got.ToSlice(), tt.want)
			}
		})
	}
}

func TestSequence_Slice(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		start int
		end   int
		want  []int
	}{
		{
			name:  "valid slice",
			slice: []int{1, 2, 3, 4, 5},
			start: 1,
			end:   3,
			want:  []int{2, 3},
		},
		{
			name:  "empty range",
			slice: []int{1, 2, 3},
			start: 1,
			end:   1,
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSequence(tt.slice)
			got := c.Slice(tt.start, tt.end)
			asSlice := make([]int, 0)
			for _, v := range got.All() {
				asSlice = append(asSlice, v)
			}
			if !slices.Equal(asSlice, tt.want) {
				t.Errorf("Slice() = %v, want %v", asSlice, tt.want)
			}
		})
	}
}
