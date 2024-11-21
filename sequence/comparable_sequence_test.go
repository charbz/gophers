package sequence

import (
	"slices"
	"testing"
)

func TestContains(t *testing.T) {
	c := NewComparableSequence([]int{1, 2, 3, 4, 5, 6})
	if !c.Contains(3) {
		t.Errorf("Contains() = %v, want %v", c.Contains(3), true)
	}
	if c.Contains(7) {
		t.Errorf("Contains() = %v, want %v", c.Contains(7), false)
	}
}

func TestExists(t *testing.T) {
	c := NewComparableSequence([]int{1, 2, 3, 4, 5, 6})
	if !c.Exists(3) {
		t.Errorf("Exists() = %v, want %v", c.Exists(3), true)
	}
	if c.Exists(7) {
		t.Errorf("Exists() = %v, want %v", c.Exists(7), false)
	}
}

func TestEquals(t *testing.T) {
	c1 := NewComparableSequence([]int{1, 2, 3})
	c2 := NewComparableSequence([]int{1, 2, 3})
	c3 := NewComparableSequence([]int{1, 2, 4})

	if !c1.Equals(c2) {
		t.Errorf("Equals() = %v, want %v", c1.Equals(c2), true)
	}
	if c1.Equals(c3) {
		t.Errorf("Equals() = %v, want %v", c1.Equals(c3), false)
	}
}

func TestDiff(t *testing.T) {
	tests := []struct {
		name string
		s1   []int
		s2   []int
		want []int
	}{
		{
			name: "different sequences",
			s1:   []int{1, 2, 3, 4},
			s2:   []int{3, 4, 5, 6},
			want: []int{1, 2},
		},
		{
			name: "identical sequences",
			s1:   []int{1, 2, 3},
			s2:   []int{1, 2, 3},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c1 := NewComparableSequence(tt.s1)
			c2 := NewComparableSequence(tt.s2)
			result := c1.Diff(c2)
			if !slices.Equal(result.elements, tt.want) {
				t.Errorf("Diff() = %v, want %v", result.elements, tt.want)
			}
		})
	}
}

func TestIndexOf(t *testing.T) {
	c := NewComparableSequence([]int{1, 2, 3, 4, 5})
	if got := c.IndexOf(3); got != 2 {
		t.Errorf("IndexOf() = %v, want %v", got, 2)
	}
	if got := c.IndexOf(6); got != -1 {
		t.Errorf("IndexOf() = %v, want %v", got, -1)
	}
}

func TestMax(t *testing.T) {
	c := NewComparableSequence([]int{1, 5, 3, 9, 2})
	if got := c.Max(); got != 9 {
		t.Errorf("Max() = %v, want %v", got, 9)
	}
}

func TestMin(t *testing.T) {
	c := NewComparableSequence([]int{4, 2, 7, 1, 9})
	if got := c.Min(); got != 1 {
		t.Errorf("Min() = %v, want %v", got, 1)
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{
			name:  "positive numbers",
			input: []int{1, 2, 3, 4, 5},
			want:  15,
		},
		{
			name:  "mixed numbers",
			input: []int{-1, 2, -3, 4},
			want:  2,
		},
		{
			name:  "empty sequence",
			input: []int{},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewComparableSequence(tt.input)
			if got := c.Sum(); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartsWith(t *testing.T) {
	tests := []struct {
		name       string
		s1         []int
		s2         []int
		startsWith bool
	}{
		{
			name:       "starts with matching elements",
			s1:         []int{1, 2, 3, 4},
			s2:         []int{1, 2},
			startsWith: true,
		},
		{
			name:       "does not start with different elements",
			s1:         []int{1, 2, 3, 4},
			s2:         []int{2, 3},
			startsWith: false,
		},
		{
			name:       "empty s2 (always true)",
			s1:         []int{1, 2, 3, 4},
			s2:         []int{},
			startsWith: true,
		},
		{
			name:       "s1 shorter than s2",
			s1:         []int{1, 2},
			s2:         []int{1, 2, 3},
			startsWith: false,
		},
		{
			name:       "both sequences empty",
			s1:         []int{},
			s2:         []int{},
			startsWith: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c1 := NewComparableSequence(tt.s1)
			c2 := NewComparableSequence(tt.s2)
			if got := c1.StartsWith(c2); got != tt.startsWith {
				t.Errorf("StartsWith() = %v, want %v", got, tt.startsWith)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		name      string
		s1        []int
		s2        []int
		endsWith  bool
	}{
		{
			name:      "ends with matching elements",
			s1:        []int{1, 2, 3, 4},
			s2:        []int{3, 4},
			endsWith:  true,
		},
		{
			name:      "does not end with different elements",
			s1:        []int{1, 2, 3, 4},
			s2:        []int{2, 3},
			endsWith:  false,
		},
		{
			name:      "empty s2 (always true)",
			s1:        []int{1, 2, 3, 4},
			s2:        []int{},
			endsWith:  true,
		},
		{
			name:      "s1 shorter than s2",
			s1:        []int{1, 2},
			s2:        []int{1, 2, 3},
			endsWith:  false,
		},
		{
			name:      "both sequences empty",
			s1:        []int{},
			s2:        []int{},
			endsWith:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c1 := NewComparableSequence(tt.s1)
			c2 := NewComparableSequence(tt.s2)
			if got := c1.EndsWith(c2); got != tt.endsWith {
				t.Errorf("EndsWith() = %v, want %v", got, tt.endsWith)
			}
		})
	}
}