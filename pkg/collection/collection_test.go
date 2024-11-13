package collection

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockCollection implements the Collection interface for testing purposes
type MockCollection[T any] struct {
	items []T
}

func NewMockCollection[T any](items []T) *MockCollection[T] {
	return &MockCollection[T]{items}
}

// Implementing the Collection interface.

func (m *MockCollection[T]) At(index int) T {
	return m.items[index]
}

func (m *MockCollection[T]) All() iter.Seq2[int, T] {
	return slices.All(m.items)
}

func (m *MockCollection[T]) Append(item T) {
	m.items = append(m.items, item)
}

func (m *MockCollection[T]) Backward() iter.Seq2[int, T] {
	return slices.Backward(m.items)
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
