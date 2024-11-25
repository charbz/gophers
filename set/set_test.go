package set

import (
	"cmp"
	"slices"
	"testing"
)

func TestSet_Contains(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate int
		want      bool
	}{
		{
			name:      "contains element matching predicate",
			slice:     []int{1, 2, 3},
			predicate: 2,
			want:      true,
		},
		{
			name:      "does not contain element matching predicate",
			slice:     []int{1, 2, 3},
			predicate: 4,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewSet(tt.slice)
			got := c.Contains(tt.predicate)
			if got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ContainsFunc(t *testing.T) {
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
			s := NewSet(tt.slice)
			got := s.ContainsFunc(tt.predicate)
			if got != tt.want {
				t.Errorf("ContainsFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Union(t *testing.T) {
	tests := []struct {
		name   string
		base   []int
		others []int
		want   []int
	}{
		{
			name:   "union with empty set",
			base:   []int{1, 2, 3},
			others: []int{},
			want:   []int{1, 2, 3},
		},
		{
			name:   "union with non-empty set",
			base:   []int{1, 2, 3},
			others: []int{3, 4, 5},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "union with no change",
			base:   []int{1, 2},
			others: []int{2, 1},
			want:   []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.base)
			result := s.Union(NewSet(tt.others))
			if !assertEqualValues(result.ToSlice(), tt.want) {
				t.Errorf("Union() = %v, want %v", result.ToSlice(), tt.want)
			}
		})
	}
}

func TestSet_UnionIterator(t *testing.T) {
	tests := []struct {
		name   string
		base   []int
		others []int
		want   []int
	}{
		{
			name:   "union with empty set",
			base:   []int{1, 2, 3},
			others: []int{},
			want:   []int{1, 2, 3},
		},
		{
			name:   "union with non-empty set",
			base:   []int{1, 2, 3},
			others: []int{3, 4, 5},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "union with no change",
			base:   []int{1, 2},
			others: []int{2, 1},
			want:   []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			s := NewSet(tt.base)
			for v := range s.UnionIterator(NewSet(tt.others)) {
				got = append(got, v)
			}
			if !assertEqualValues(got, tt.want) {
				t.Errorf("UnionIterator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	tests := []struct {
		name   string
		base   []int
		others []int
		want   []int
	}{
		{
			name:   "intersection with empty set",
			base:   []int{1, 2, 3},
			others: []int{},
			want:   []int{},
		},
		{
			name:   "intersection with non-empty set",
			base:   []int{1, 2, 3},
			others: []int{2, 3, 4},
			want:   []int{2, 3},
		},
		{
			name:   "intersection with non-empty set 2",
			base:   []int{1, 2, 3, 4},
			others: []int{4, 5, 6},
			want:   []int{4},
		},
		{
			name:   "intersection with no overlap",
			base:   []int{1, 2, 3, 4},
			others: []int{5, 6, 7},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.base)
			result := s.Intersection(NewSet(tt.others))
			if !assertEqualValues(result.ToSlice(), tt.want) {
				t.Errorf("Intersection() = %v, want %v", result.ToSlice(), tt.want)
			}
		})
	}
}

func TestSet_IntersectionIterator(t *testing.T) {
	tests := []struct {
		name   string
		base   []int
		others []int
		want   []int
	}{
		{
			name:   "intersection with empty set",
			base:   []int{1, 2, 3},
			others: []int{},
			want:   []int{},
		},
		{
			name:   "intersection with non-empty set",
			base:   []int{1, 2, 3},
			others: []int{2, 3, 4},
			want:   []int{2, 3},
		},
		{
			name:   "intersection with non-empty set 2",
			base:   []int{1, 2, 3, 4},
			others: []int{4, 5, 6},
			want:   []int{4},
		},
		{
			name:   "intersection with no overlap",
			base:   []int{1, 2, 3, 4},
			others: []int{5, 6, 7},
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []int{}
			s := NewSet(tt.base)
			for v := range s.IntersectionIterator(NewSet(tt.others)) {
				got = append(got, v)
			}
			if !assertEqualValues(got, tt.want) {
				t.Errorf("IntersectionIterator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Diff(t *testing.T) {
	tests := []struct {
		name string
		base []int
		diff []int
		want []int
	}{
		{
			name: "diff with empty set",
			base: []int{1, 2, 3},
			diff: []int{},
			want: []int{1, 2, 3},
		},
		{
			name: "diff with non-empty set",
			base: []int{1, 2, 3},
			diff: []int{2, 3},
			want: []int{1},
		},
		{
			name: "diff with no overlap",
			base: []int{1, 2, 3},
			diff: []int{4, 5},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s1 := NewSet(tt.base)
			s2 := NewSet(tt.diff)
			result := s1.Diff(s2)
			if !assertEqualValues(result.ToSlice(), tt.want) {
				t.Errorf("Diff() = %v, want %v", result.ToSlice(), tt.want)
			}
		})
	}
}

func TestSet_Equals(t *testing.T) {
	tests := []struct {
		name string
		s1   []int
		s2   []int
		want bool
	}{
		{
			name: "equal sets",
			s1:   []int{1, 2, 3},
			s2:   []int{1, 2, 3},
			want: true,
		},
		{
			name: "different lengths",
			s1:   []int{1, 2, 3},
			s2:   []int{1, 2},
			want: false,
		},
		{
			name: "different elements",
			s1:   []int{1, 2, 3},
			s2:   []int{1, 2, 4},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s1 := NewSet(tt.s1)
			s2 := NewSet(tt.s2)
			got := s1.Equals(s2)
			if got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  bool
	}{
		{
			name:  "empty set",
			slice: []int{},
			want:  true,
		},
		{
			name:  "non-empty set",
			slice: []int{1, 2, 3},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.slice)
			got := s.IsEmpty()
			if got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Clone(t *testing.T) {
	original := NewSet([]int{1, 2, 3})
	clone := original.Clone()

	// Verify clone has same elements
	if !assertEqualValues(original.ToSlice(), clone.ToSlice()) {
		t.Errorf("Clone() = %v, want %v", clone.ToSlice(), original.ToSlice())
	}

	// Verify modifying clone doesn't affect original
	clone.Add(4)
	if slices.Contains(original.ToSlice(), 4) {
		t.Errorf("Clone() = %v, want %v", clone.ToSlice(), original.ToSlice())
	}
	if !slices.Contains(clone.ToSlice(), 4) {
		t.Errorf("Clone() = %v, want %v", clone.ToSlice(), original.ToSlice())
	}
}

func TestSet_Partition(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		wantLeft  []int
		wantRight []int
	}{
		{
			name:      "partition evens and odds",
			slice:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(i int) bool { return i%2 == 0 },
			wantLeft:  []int{2, 4, 6},
			wantRight: []int{1, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.slice)
			left, right := s.Partition(tt.predicate)
			if !assertEqualValues(left.ToSlice(), tt.wantLeft) {
				t.Errorf("Partition() = %v, want %v", left.ToSlice(), tt.wantLeft)
			}
			if !assertEqualValues(right.ToSlice(), tt.wantRight) {
				t.Errorf("Partition() = %v, want %v", right.ToSlice(), tt.wantRight)
			}
		})
	}
}

func TestSet_NonEmpty(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  bool
	}{
		{
			name:  "empty set",
			slice: []int{},
			want:  false,
		},
		{
			name:  "non-empty set",
			slice: []int{1, 2, 3},
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.slice)
			got := s.NonEmpty()
			if got != tt.want {
				t.Errorf("NonEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_ForAll(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      bool
	}{
		{
			name:      "all elements match predicate",
			slice:     []int{2, 4, 6},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      true,
		},
		{
			name:      "not all elements match predicate",
			slice:     []int{1, 2, 3},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.slice)
			got := s.ForAll(tt.predicate)
			if got != tt.want {
				t.Errorf("ForAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Filter(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "filter even numbers",
			slice:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{2, 4, 6},
		},
		{
			name:      "filter nothing matches",
			slice:     []int{1, 3, 5},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.slice)
			result := s.Filter(tt.predicate)
			if !assertEqualValues(result.ToSlice(), tt.want) {
				t.Errorf("Filter() = %v, want %v", result.ToSlice(), tt.want)
			}
		})
	}
}

func TestSet_FilterNot(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "filter out even numbers",
			slice:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{1, 3, 5},
		},
		{
			name:      "filter out nothing matches",
			slice:     []int{2, 4, 6},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.slice)
			result := s.FilterNot(tt.predicate)
			if !assertEqualValues(result.ToSlice(), tt.want) {
				t.Errorf("FilterNot() = %v, want %v", result.ToSlice(), tt.want)
			}
		})
	}
}

func TestSet_Count(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      int
	}{
		{
			name:      "count even numbers",
			slice:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      3,
		},
		{
			name:      "count nothing matches",
			slice:     []int{1, 3, 5},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.slice)
			got := s.Count(tt.predicate)
			if got != tt.want {
				t.Errorf("Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Random(t *testing.T) {
	s := NewSet([]int{1})
	if got := s.Random(); got != 1 {
		t.Errorf("Random() = %v, want %v", got, 1)
	}
}

func TestSet_Remove(t *testing.T) {
	s := NewSet([]int{1, 2, 3})
	s.Remove(2)
	if !assertEqualValues(s.ToSlice(), []int{1, 3}) {
		t.Errorf("Remove() = %v, want %v", s.ToSlice(), []int{1, 3})
	}
}

func TestSet_Reject(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "filter out even numbers",
			slice:     []int{1, 2, 3, 4, 5, 6},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{1, 3, 5},
		},
		{
			name:      "filter out nothing matches",
			slice:     []int{2, 4, 6},
			predicate: func(i int) bool { return i%2 == 0 },
			want:      []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSet(tt.slice)
			result := s.Reject(tt.predicate)
			if !assertEqualValues(result.ToSlice(), tt.want) {
				t.Errorf("Reject() = %v, want %v", result.ToSlice(), tt.want)
			}
		})
	}
}

func TestDiffIterator(t *testing.T) {
	tests := []struct {
		name string
		a    *Set[int]
		b    *Set[int]
		want []int
	}{
		{
			name: "diff with empty set",
			a:    NewSet([]int{1, 2, 3}),
			b:    NewSet([]int{}),
			want: []int{1, 2, 3},
		},
		{
			name: "diff with non-empty set",
			a:    NewSet([]int{1, 2, 3, 5, 6}),
			b:    NewSet([]int{2, 3, 4}),
			want: []int{1, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collected := []int{}
			for v := range tt.a.DiffIterator(tt.b) {
				collected = append(collected, v)
			}
			if !assertEqualValues(collected, tt.want) {
				t.Errorf("DiffIterator() = %v, want %v", collected, tt.want)
			}
		})
	}
}

func assertEqualValues[T cmp.Ordered](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	slices.Sort(a)
	slices.Sort(b)
	return slices.Equal(a, b)
}
