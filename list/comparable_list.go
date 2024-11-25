// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package list

import (
	"cmp"
	"iter"

	"github.com/charbz/gophers/collection"
)

// ComparableList is a list of comparable types.
// it is similar to List, but with additional methods that do not require a
// higher order function comparator to be provided as an argument:
// Max(), Min(), Sum(), Distinct(), Diff(c), and Exists(v).
type ComparableList[T cmp.Ordered] struct {
	List[T]
}

func (l *ComparableList[T]) New(s ...[]T) collection.Collection[T] {
	return NewComparableList(s...)
}

func (l *ComparableList[T]) NewOrdered(s ...[]T) collection.OrderedCollection[T] {
	return NewComparableList(s...)
}

func NewComparableList[T cmp.Ordered](s ...[]T) *ComparableList[T] {
	list := new(ComparableList[T])
	if len(s) == 0 {
		return list
	}
	for _, slice := range s {
		for _, v := range slice {
			list.Add(v)
		}
	}
	return list
}

// Clone returns a copy of the list. This is a shallow clone.
func (l *ComparableList[T]) Clone() *ComparableList[T] {
	clone := &ComparableList[T]{}
	for v := range l.Values() {
		clone.Add(v)
	}
	return clone
}

// Concat returns a new list concatenating the passed in lists.
func (l *ComparableList[T]) Concat(lists ...*ComparableList[T]) *ComparableList[T] {
	clone := l.Clone()
	for _, list := range lists {
		for v := range list.Values() {
			clone.Add(v)
		}
	}
	return clone
}

// Concatenated is an alias for collection.Concatenated
func (l *ComparableList[T]) Concatenated(l2 *ComparableList[T]) iter.Seq[T] {
	return collection.Concatenated(l, l2)
}

// Contains returns true if the list contains the given value.
func (l *ComparableList[T]) Contains(v T) bool {
	for val := range l.Values() {
		if val == v {
			return true
		}
	}
	return false
}

// Corresponds is an alias for collection.Corresponds
func (l *ComparableList[T]) Corresponds(s *ComparableList[T], f func(T, T) bool) bool {
	return collection.Corresponds(l, s, f)
}

// Distinct returns a new list containing only the unique elements from the original list.
func (l *ComparableList[T]) Distinct() *ComparableList[T] {
	m := make(map[T]struct{})
	r := &ComparableList[T]{}
	for v := range l.Values() {
		_, ok := m[v]
		if !ok {
			r.Add(v)
			m[v] = struct{}{}
		}
	}
	return r
}

// Distincted is an alias for collection.Distincted
func (l *ComparableList[T]) Distincted() iter.Seq[T] {
	return collection.Distincted(l)
}

// Diff returns a new list containing the elements of the original list that are not in the other list.
func (l *ComparableList[T]) Diff(s *ComparableList[T]) *ComparableList[T] {
	return collection.Diff(l, s).(*ComparableList[T])
}

// Diffed is an alias for collection.Diffed
func (l *ComparableList[T]) Diffed(s *ComparableList[T]) iter.Seq[T] {
	return collection.Diffed(l, s)
}

// Exists is an alias for Contains
func (l *ComparableList[T]) Exists(v T) bool {
	return l.Contains(v)
}

// Equals returns true if the two lists are equal.
func (l *ComparableList[T]) Equals(s *ComparableList[T]) bool {
	if l.size != s.size {
		return false
	}
	n1 := l.head
	n2 := s.head
	for n1 != nil && n2 != nil {
		if n1.value != n2.value {
			return false
		}
		n1 = n1.next
		n2 = n2.next
	}
	return true
}

// IndexOf returns the index of the first occurrence of the specified element in this list,
func (l *ComparableList[T]) IndexOf(v T) int {
	for i, val := range l.All() {
		if val == v {
			return i
		}
	}
	return -1
}

// Intersect returns a new list containing the elements that are present in both lists.
func (l *ComparableList[T]) Intersect(s *ComparableList[T]) *ComparableList[T] {
	return collection.Intersect(l, s).(*ComparableList[T])
}

// Intersected is an alias for collection.Intersected
func (l *ComparableList[T]) Intersected(s *ComparableList[T]) iter.Seq[T] {
	return collection.Intersected(l, s)
}

// LastIndexOf returns the index of the last occurrence of the specified element in this list,
func (l *ComparableList[T]) LastIndexOf(v T) int {
	for i, val := range l.Backward() {
		if val == v {
			return i
		}
	}
	return -1
}

// Max returns the maximum element in the list.
func (l *ComparableList[T]) Max() (T, error) {
	return collection.MaxBy(l, func(v T) T { return v })
}

// Min returns the minimum element in the list.
func (l *ComparableList[T]) Min() (T, error) {
	return collection.MinBy(l, func(v T) T { return v })
}

// Sum returns the sum of the elements in the list.
func (l *ComparableList[T]) Sum() T {
	var sum T
	for v := range l.Values() {
		sum += v
	}
	return sum
}

// StartsWith returns true if the list starts with the given list.
func (l *ComparableList[T]) StartsWith(other *ComparableList[T]) bool {
	return collection.StartsWith(l, other)
}

// EndsWith returns true if the list ends with the given list.
func (l *ComparableList[T]) EndsWith(other *ComparableList[T]) bool {
	return collection.EndsWith(l, other)
}
