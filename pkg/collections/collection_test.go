package collections

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Foo struct {
	bar  int
	bars []string
}

func TestComparable_NewCollection(t *testing.T) {
	tests := []struct {
		name string
		args [][]int
		want []int
	}{
		{
			name: "initialize empty",
			args: nil,
			want: nil,
		},
		{
			name: "initialize with 1 slice",
			args: [][]int{{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "initialize with 2 slices",
			args: [][]int{{1, 2, 3, 4, 5}, {0, 7, 8}},
			want: []int{1, 2, 3, 4, 5, 0, 7, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCollection(tt.args...)
			assert.Equal(t, tt.want, c.elements)
		})
	}
}

func TestAny_NewCollection(t *testing.T) {
	tests := []struct {
		name string
		args [][]Foo
		want []Foo
	}{
		{
			name: "initialize empty",
			args: nil,
			want: nil,
		},
		{
			name: "initialize with 1 slice",
			args: [][]Foo{{Foo{}, Foo{1, []string{"a"}}, Foo{2, []string{"a", "b"}}}},
			want: []Foo{{}, {1, []string{"a"}}, {2, []string{"a", "b"}}},
		},
		{
			name: "initialize with 2 slices",
			args: [][]Foo{{Foo{1, nil}, Foo{2, nil}}, {Foo{3, nil}, Foo{4, nil}}},
			want: []Foo{{1, nil}, {2, nil}, {3, nil}, {4, nil}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCollection(tt.args...)
			assert.Equal(t, tt.want, c.elements)
		})
	}

}
