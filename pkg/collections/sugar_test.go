package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	tests := []struct {
		name     string
		base     []int
		toConcat [][]int
		want     []int
	}{
		{
			name:     "single collection concat",
			base:     []int{1, 2},
			toConcat: [][]int{{3, 4}},
			want:     []int{1, 2, 3, 4},
		},
		{
			name:     "multiple collections concat",
			base:     []int{1, 2},
			toConcat: [][]int{{3, 4}, {5, 6}},
			want:     []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCollection(tt.base)
			var collections []Collection[int]
			for _, slice := range tt.toConcat {
				collections = append(collections, *NewCollection(slice))
			}

			result := c.Concat(collections...)
			assert.Equal(t, tt.want, result.elements)
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
			c := NewCollection(tt.input)
			result := c.Distinct(func(a, b int) bool {
				return a == b
			})
			assert.Equal(t, tt.want, result.elements)
		})
	}
}
