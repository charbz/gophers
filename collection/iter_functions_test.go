package collection

import (
	"slices"
	"testing"
)

func TestConcatenated(t *testing.T) {
	tests := []struct {
		name string
		a    OrderedCollection[int]
		b    OrderedCollection[int]
		want []int
	}{
		{
			name: "concat two lists",
			a:    NewMockOrderedCollection([]int{1, 2}),
			b:    NewMockOrderedCollection([]int{3, 4}),
			want: []int{1, 2, 3, 4},
		},
		{
			name: "concat empty list A",
			a:    NewMockOrderedCollection([]int{}),
			b:    NewMockOrderedCollection([]int{1, 2, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "concat empty list B",
			a:    NewMockOrderedCollection([]int{1, 2, 3}),
			b:    NewMockOrderedCollection([]int{}),
			want: []int{1, 2, 3},
		},
		{
			name: "concat empty lists A and B",
			a:    NewMockOrderedCollection([]int{}),
			b:    NewMockOrderedCollection([]int{}),
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collected := []int{}
			for v := range Concatenated(tt.a, tt.b) {
				collected = append(collected, v)
			}
			if !slices.Equal(collected, tt.want) {
				t.Errorf("Concatenated() = %v, want %v", collected, tt.want)
			}
		})
	}
}

func TestDiffed(t *testing.T) {
	tests := []struct {
		name string
		a    OrderedCollection[int]
		b    OrderedCollection[int]
		want []int
	}{
		{
			name: "diff same elements",
			a:    NewMockOrderedCollection([]int{1, 2, 3}),
			b:    NewMockOrderedCollection([]int{1, 2, 3}),
			want: []int{},
		},
		{
			name: "diff different elements",
			a:    NewMockOrderedCollection([]int{1, 2, 3, 6, 5}),
			b:    NewMockOrderedCollection([]int{2, 4, 6}),
			want: []int{1, 3, 5},
		},
		{
			name: "diff empty B",
			a:    NewMockOrderedCollection([]int{1, 2, 3}),
			b:    NewMockOrderedCollection([]int{}),
			want: []int{1, 2, 3},
		},
		{
			name: "diff empty A",
			a:    NewMockOrderedCollection([]int{}),
			b:    NewMockOrderedCollection([]int{1, 2, 3}),
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collected := []int{}
			for v := range Diffed(tt.a, tt.b) {
				collected = append(collected, v)
			}
			if !slices.Equal(collected, tt.want) {
				t.Errorf("Diffed() = %v, want %v", collected, tt.want)
			}
		})
	}
}

func TestDiffedFunc(t *testing.T) {
	tests := []struct {
		name string
		a    OrderedCollection[int]
		b    OrderedCollection[int]
		want []int
	}{
		{
			name: "diff same elements",
			a:    NewMockOrderedCollection([]int{1, 2, 3}),
			b:    NewMockOrderedCollection([]int{1, 2, 3}),
			want: []int{},
		},
		{
			name: "diff different elements",
			a:    NewMockOrderedCollection([]int{1, 2, 3, 6, 5}),
			b:    NewMockOrderedCollection([]int{2, 4, 6}),
			want: []int{1, 3, 5},
		},
		{
			name: "diff empty B",
			a:    NewMockOrderedCollection([]int{1, 2, 3}),
			b:    NewMockOrderedCollection([]int{}),
			want: []int{1, 2, 3},
		},
		{
			name: "diff empty A",
			a:    NewMockOrderedCollection([]int{}),
			b:    NewMockOrderedCollection([]int{1, 2, 3}),
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collected := []int{}
			for v := range DiffedFunc(tt.a, tt.b, func(a, b int) bool { return a == b }) {
				collected = append(collected, v)
			}
			if !slices.Equal(collected, tt.want) {
				t.Errorf("DiffedFunc() = %v, want %v", collected, tt.want)
			}
		})
	}
}

func TestDistincted(t *testing.T) {
	tests := []struct {
		name string
		a    OrderedCollection[int]
		want []int
	}{
		{
			name: "distinct elements",
			a:    NewMockOrderedCollection([]int{1, 1, 2, 2, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "distinct with no duplicates",
			a:    NewMockOrderedCollection([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "distinct with empty collection",
			a:    NewMockOrderedCollection([]int{}),
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collected := []int{}
			for v := range Distincted(tt.a) {
				collected = append(collected, v)
			}
			if !slices.Equal(collected, tt.want) {
				t.Errorf("Distincted() = %v, want %v", collected, tt.want)
			}
		})
	}
}

func TestDistinctedFunc(t *testing.T) {
	tests := []struct {
		name string
		a    OrderedCollection[int]
		want []int
	}{
		{
			name: "distinct elements",
			a:    NewMockOrderedCollection([]int{1, 1, 2, 2, 3}),
			want: []int{1, 2, 3},
		},
		{
			name: "distinct with no duplicates",
			a:    NewMockOrderedCollection([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "distinct with empty collection",
			a:    NewMockOrderedCollection([]int{}),
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collected := []int{}
			for v := range DistinctedFunc(tt.a, func(a, b int) bool { return a == b }) {
				collected = append(collected, v)
			}
			if !slices.Equal(collected, tt.want) {
				t.Errorf("DistinctedFunc() = %v, want %v", collected, tt.want)
			}
		})
	}
}

func TestIntersected(t *testing.T) {
	tests := []struct {
		name string
		a    OrderedCollection[int]
		b    OrderedCollection[int]
		want []int
	}{
		{
			name: "intersect",
			a:    NewMockOrderedCollection([]int{1, 2, 3, 4, 5, 6, 6, 4}),
			b:    NewMockOrderedCollection([]int{2, 4, 6, 8, 10, 12}),
			want: []int{2, 4, 6, 6, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collected := []int{}
			for v := range Intersected(tt.a, tt.b) {
				collected = append(collected, v)
			}
			if !slices.Equal(collected, tt.want) {
				t.Errorf("Intersected() = %v, want %v", collected, tt.want)
			}
		})
	}
}

func TestIntersectedFunc(t *testing.T) {
	tests := []struct {
		name string
		a    OrderedCollection[int]
		b    OrderedCollection[int]
		want []int
	}{
		{
			name: "intersect",
			a:    NewMockOrderedCollection([]int{1, 2, 3, 4, 5, 6, 6, 4}),
			b:    NewMockOrderedCollection([]int{2, 4, 6, 8, 10, 12}),
			want: []int{2, 4, 6, 6, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collected := []int{}
			for v := range IntersectedFunc(tt.a, tt.b, func(a, b int) bool { return a == b }) {
				collected = append(collected, v)
			}
			if !slices.Equal(collected, tt.want) {
				t.Errorf("IntersectedFunc() = %v, want %v", collected, tt.want)
			}
		})
	}
}

func TestFiltered(t *testing.T) {
	tests := []struct {
		name string
		a    OrderedCollection[int]
		want []int
	}{
		{
			name: "filter even numbers",
			a:    NewMockOrderedCollection([]int{1, 2, 3, 4, 5, 6}),
			want: []int{2, 4, 6},
		},
		{
			name: "filter no matches",
			a:    NewMockOrderedCollection([]int{1, 3, 5, 7, 9}),
			want: []int{},
		},
		{
			name: "filter empty collection",
			a:    NewMockOrderedCollection([]int{}),
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collected := []int{}
			for v := range Filtered(tt.a, func(i int) bool { return i%2 == 0 }) {
				collected = append(collected, v)
			}
			if !slices.Equal(collected, tt.want) {
				t.Errorf("Filtered() = %v, want %v", collected, tt.want)
			}
		})
	}
}

func TestMapped(t *testing.T) {
	tests := []struct {
		name string
		a    OrderedCollection[int]
		want []int
	}{
		{
			name: "double numbers",
			a:    NewMockOrderedCollection([]int{1, 2, 3}),
			want: []int{2, 4, 6},
		},
		{
			name: "empty collection",
			a:    NewMockOrderedCollection([]int{}),
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collected := []int{}
			for v := range Mapped(tt.a, func(i int) int { return i * 2 }) {
				collected = append(collected, v)
			}
			if !slices.Equal(collected, tt.want) {
				t.Errorf("Mapped() = %v, want %v", collected, tt.want)
			}
		})
	}
}

func TestRejected(t *testing.T) {
	tests := []struct {
		name string
		a    OrderedCollection[int]
		want []int
	}{
		{
			name: "reject even numbers",
			a:    NewMockOrderedCollection([]int{1, 2, 3, 4, 5, 6}),
			want: []int{1, 3, 5},
		},
		{
			name: "reject no matches",
			a:    NewMockOrderedCollection([]int{2, 4, 6, 8}),
			want: []int{},
		},
		{
			name: "reject empty collection",
			a:    NewMockOrderedCollection([]int{}),
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collected := []int{}
			for v := range Rejected(tt.a, func(i int) bool { return i%2 == 0 }) {
				collected = append(collected, v)
			}
			if !slices.Equal(collected, tt.want) {
				t.Errorf("Rejected() = %v, want %v", collected, tt.want)
			}
		})
	}
}
