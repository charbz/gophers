package sequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	c := NewComparableSequence([]int{1, 2, 3, 4, 5, 6})
	assert.True(t, c.Contains(3))
	assert.False(t, c.Contains(7))
}

func TestExists(t *testing.T) {
	c := NewComparableSequence([]int{1, 2, 3, 4, 5, 6})
	assert.True(t, c.Exists(3))
	assert.False(t, c.Exists(7))
}

func TestEquals(t *testing.T) {
	c1 := NewComparableSequence([]int{1, 2, 3})
	c2 := NewComparableSequence([]int{1, 2, 3})
	c3 := NewComparableSequence([]int{1, 2, 4})

	assert.True(t, c1.Equals(c2))
	assert.False(t, c1.Equals(c3))
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
			assert.Equal(t, tt.want, result.elements)
		})
	}
}

func TestIndexOf(t *testing.T) {
	c := NewComparableSequence([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 2, c.IndexOf(3))
	assert.Equal(t, -1, c.IndexOf(6))
}

func TestMax(t *testing.T) {
	c := NewComparableSequence([]int{1, 5, 3, 9, 2})
	assert.Equal(t, 9, c.Max())
}

func TestMin(t *testing.T) {
	c := NewComparableSequence([]int{4, 2, 7, 1, 9})
	assert.Equal(t, 1, c.Min())
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
			assert.Equal(t, tt.want, c.Sum())
		})
	}
}
