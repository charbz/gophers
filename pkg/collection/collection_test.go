package collection

import (
	"iter"
	"math/rand"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockCollection implements the Collection interface for testing purposes
type MockCollection[T any] struct {
	items []T
}

func NewMockCollection[T any](items ...[]T) *MockCollection[T] {
	return &MockCollection[T]{items: slices.Concat(items...)}
}

// Implementing the Collection interface.

func (m *MockCollection[T]) Add(item T) {
	m.items = append(m.items, item)
}

func (m *MockCollection[T]) Values() iter.Seq[T] {
	return slices.Values(m.items)
}

func (m *MockCollection[T]) Length() int {
	return len(m.items)
}

func (m *MockCollection[T]) ToSlice() []T {
	return m.items
}

func (c *MockCollection[T]) Random() T {
	if len(c.items) == 0 {
		return *new(T)
	}
	return c.items[rand.Intn(len(c.items))]
}

func (m *MockCollection[T]) Slice(start, end int) Collection[T] {
	return NewMockCollection(m.items[start:end])
}

func (m *MockCollection[T]) New(s ...[]T) Collection[T] {
	mock := &MockCollection[T]{}
	if len(s) > 0 {
		mock.items = s[0]
	}
	return mock
}

func TestMockCollectionImplementsCollection(t *testing.T) {
	var m Collection[string] = NewMockCollection([]string{"a", "b", "c"})
	assert.Equal(t, 3, m.Length())
}
