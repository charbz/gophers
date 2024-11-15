package collection

import (
	"iter"
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockOrderedCollection implements the OrderedCollection interface for testing purposes
type MockOrderedCollection[T any] struct {
	items []T
}

func NewMockOrderedCollection[T any](items ...[]T) *MockOrderedCollection[T] {
	return &MockOrderedCollection[T]{items: slices.Concat(items...)}
}

// Implementing the Collection & OrderedCollection interfaces.

func (m *MockOrderedCollection[T]) At(index int) T {
	return m.items[index]
}

func (m *MockOrderedCollection[T]) All() iter.Seq2[int, T] {
	return slices.All(m.items)
}

func (m *MockOrderedCollection[T]) Append(item T) {
	m.items = append(m.items, item)
}

func (m *MockOrderedCollection[T]) Backward() iter.Seq2[int, T] {
	return slices.Backward(m.items)
}

func (m *MockOrderedCollection[T]) Values() iter.Seq[T] {
	return slices.Values(m.items)
}

func (m *MockOrderedCollection[T]) Length() int {
	return len(m.items)
}

func (m *MockOrderedCollection[T]) ToSlice() []T {
	return m.items
}

func (c *MockOrderedCollection[T]) Random() T {
	if len(c.items) == 0 {
		return *new(T)
	}
	return c.items[rand.Intn(len(c.items))]
}

func (m *MockOrderedCollection[T]) Slice(start, end int) OrderedCollection[T] {
	return NewMockOrderedCollection(m.items[start:end])
}

func (m *MockOrderedCollection[T]) NewOrdered(s ...[]T) OrderedCollection[T] {
	return NewMockOrderedCollection(s...)
}

func (m *MockOrderedCollection[T]) New(s ...[]T) Collection[T] {
	return NewMockOrderedCollection(s...)
}

func TestMockOrderedCollectionImplementsOrderedCollection(t *testing.T) {
	var m OrderedCollection[string] = NewMockOrderedCollection([]string{"a", "b", "c"})
	assert.Equal(t, 3, m.Length())
}
