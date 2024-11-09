package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollection_At(t *testing.T) {
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
			c := NewCollection(tt.slice)

			if tt.wantErr {
				assert.Panics(t, func() { c.At(tt.index) })
			} else {
				got := c.At(tt.index)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestCollection_Contains(t *testing.T) {
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
			c := NewCollection(tt.slice)
			got := c.Contains(tt.predicate)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCollection_Drop(t *testing.T) {
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
			c := NewCollection(tt.slice)
			got := c.Drop(tt.n)
			assert.Equal(t, tt.want, got.ToSlice())
		})
	}
}

func TestCollection_Filter(t *testing.T) {
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
			want:   []int{},
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
			c := NewCollection(tt.slice)
			got := c.Filter(tt.filter)
			assert.Equal(t, tt.want, got.ToSlice())
		})
	}
}

func TestCollection_DropRight(t *testing.T) {
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
			c := NewCollection(tt.slice)
			got := c.DropRight(tt.n)
			assert.Equal(t, tt.want, got.ToSlice())
		})
	}
}

func TestCollection_FilterNot(t *testing.T) {
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
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCollection(tt.slice)
			got := c.FilterNot(tt.filter)
			assert.Equal(t, tt.want, got.ToSlice())
		})
	}
}

func TestCollection_Find(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      int
		wantErr   bool
	}{
		{
			name:      "find existing element",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i > 3 },
			want:      4,
			wantErr:   false,
		},
		{
			name:      "element not found",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return i > 5 },
			want:      0, // zero value for int
			wantErr:   true,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(i int) bool { return true },
			want:      0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCollection(tt.slice)
			got, err := c.Find(tt.predicate)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestCollection_FindWhere(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      int
	}{
		{
			name:      "find index of existing element",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(i int) bool { return i > 3 },
			want:      3,
		},
		{
			name:      "element not found",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return i > 5 },
			want:      -1,
		},
		{
			name:      "empty slice",
			slice:     []int{},
			predicate: func(i int) bool { return true },
			want:      -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCollection(tt.slice)
			got := c.FindWhere(tt.predicate)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCollection_ForEach(t *testing.T) {
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
			name:  "empty slice",
			slice: []int{},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCollection(tt.slice)
			sum := 0
			c.ForEach(func(i int) {
				sum += i
			})
			assert.Equal(t, tt.want, sum)
		})
	}
}

func TestCollection_Head(t *testing.T) {
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
			c := NewCollection(tt.slice)
			got, err := c.Head()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// Helper function to compare slices
func sliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
