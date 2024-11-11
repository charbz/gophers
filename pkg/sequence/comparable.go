package sequence

import (
	"slices"
)

type ComparableSequence[T comparable] struct {
	Sequence[T]
}

// NewComparableSequence is a constructor for a sequence of comparable types.
func NewComparableSequence[T comparable](s ...[]T) *ComparableSequence[T] {
	seq := new(ComparableSequence[T])
	if len(s) == 0 {
		return seq
	}
	return &ComparableSequence[T]{Sequence[T]{elements: slices.Concat(s...)}}
}

// Contains returns true if the sequence contains the given value.
func (c *ComparableSequence[T]) Contains(v T) bool {
	return slices.Contains(c.elements, v)
}

// Distinct returns a new sequence containing only the unique elements from the original sequence.
func (c *ComparableSequence[T]) Distinct() *ComparableSequence[T] {
	m := make(map[T]interface{})
	r := &ComparableSequence[T]{}
	for v := range c.Values() {
		_, ok := m[v]
		if !ok {
			r.Append(v)
			m[v] = true
		}
	}
	return r
}

// Exists returns true if the sequence contains the given value.
func (c *ComparableSequence[T]) Exists(v T) bool {
	return c.Contains(v)
}

// IndexOf returns the index of the first occurrence of the specified element in this sequence,
// or -1 if this sequence does not contain the element.
func (c *ComparableSequence[T]) IndexOf(v T) int {
	return slices.Index(c.elements, v)
}

func (c *ComparableSequence[T]) New(s ...[]T) *ComparableSequence[T] {
	return NewComparableSequence(s...)
}
