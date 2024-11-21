package sequence

import (
	"cmp"
	"slices"

	"github.com/charbz/gophers/collection"
)

// ComparableSequence is a sequence of comparable types.
// it is similar to Sequence, but with additional methods that do not require a
// higher order function comparator to be provided as an argument:
// Max(), Min(), Sum(), Distinct(), Diff(c), and Exists(v).
type ComparableSequence[T cmp.Ordered] struct {
	Sequence[T]
}

func (c *ComparableSequence[T]) New(s ...[]T) collection.Collection[T] {
	return NewComparableSequence(s...)
}

func (c *ComparableSequence[T]) NewOrdered(s ...[]T) collection.OrderedCollection[T] {
	return NewComparableSequence(s...)
}

// NewComparableSequence is a constructor for a sequence of comparable types.
func NewComparableSequence[T cmp.Ordered](s ...[]T) *ComparableSequence[T] {
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
			r.Add(v)
			m[v] = true
		}
	}
	return r
}

func (c *ComparableSequence[T]) Diff(s *ComparableSequence[T]) *ComparableSequence[T] {
	return collection.Diff(c, s).(*ComparableSequence[T])
}

// Equals returns true if the two sequences are equal.
func (c *ComparableSequence[T]) Equals(c2 *ComparableSequence[T]) bool {
	return slices.Equal(c.elements, c2.elements)
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

// LastIndexOf returns the index of the last occurrence of the specified element in this sequence,
// or -1 if this sequence does not contain the element.
func (c *ComparableSequence[T]) LastIndexOf(v T) int {
	for i, val := range c.Backward() {
		if val == v {
			return i
		}
	}
	return -1
}

func (c *ComparableSequence[T]) Max() T {
	return slices.Max(c.elements)
}

func (c *ComparableSequence[T]) Min() T {
	return slices.Min(c.elements)
}

func (c *ComparableSequence[T]) Sum() T {
	var sum T
	for _, v := range c.elements {
		sum += v
	}
	return sum
}

func (c *ComparableSequence[T]) StartsWith(other *ComparableSequence[T]) bool {
	return collection.StartsWith(c, other)
}

func (c *ComparableSequence[T]) EndsWith(other *ComparableSequence[T]) bool {
	return collection.EndsWith(c, other)
}
