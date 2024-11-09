package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc_Contains(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		search   int
		expected bool
	}{
		{
			name:     "contains value",
			input:    []int{1, 2, 3, 4, 5},
			search:   3,
			expected: true,
		},
		{
			name:     "does not contain value",
			input:    []int{1, 2, 3, 4, 5},
			search:   6,
			expected: false,
		},
		{
			name:     "empty collection",
			input:    []int{},
			search:   1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collection := NewCollection(tt.input)
			result := Contains(collection, tt.search)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFunc_Distinct(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "all unique elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "with duplicates",
			input:    []int{1, 2, 2, 3, 3, 3, 4, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "empty collection",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collection := NewCollection(tt.input)
			result := Distinct(collection)
			assert.ElementsMatch(t, tt.expected, result.ToSlice())
		})
	}
}

func TestFunc_Map(t *testing.T) {
	t.Run("string length mapping", func(t *testing.T) {
		input := NewCollection([]string{"Alice", "Bob", "Charlie"})
		result := Map(input, func(s string) int {
			return len(s)
		})
		assert.ElementsMatch(t, []int{5, 3, 7}, result.ToSlice())
	})

	t.Run("integer doubling", func(t *testing.T) {
		input := NewCollection([]int{1, 2, 3, 4, 5})
		result := Map(input, func(i int) int {
			return i * 2
		})
		assert.ElementsMatch(t, []int{2, 4, 6, 8, 10}, result.ToSlice())
	})

	t.Run("empty collection", func(t *testing.T) {
		input := NewCollection([]int{})
		result := Map(input, func(i int) string {
			return "test"
		})
		assert.Empty(t, result.ToSlice())
	})
}

func TestFunc_Reduce(t *testing.T) {
	t.Run("sum reduction", func(t *testing.T) {
		input := NewCollection([]int{1, 2, 3, 4, 5, 6})
		result := Reduce(input, func(acc, curr int) int {
			return acc + curr
		}, 0)
		assert.Equal(t, 21, result)
	})

	t.Run("string concatenation", func(t *testing.T) {
		input := NewCollection([]string{"Hello", " ", "World", "!"})
		result := Reduce(input, func(acc, curr string) string {
			return acc + curr
		}, "")
		assert.Equal(t, "Hello World!", result)
	})

	t.Run("empty collection", func(t *testing.T) {
		input := NewCollection([]int{})
		result := Reduce(input, func(acc, curr int) int {
			return acc + curr
		}, 0)
		assert.Equal(t, 0, result)
	})
}
