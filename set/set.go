// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// Package set implements support for a generic unordered Set.
// A Set is a Collection that wraps an underlying hash map
// and provides convenience methods and syntatic sugar on top of it.
//
// Set elements are unique and unordered by default. However Sets share
// some methods with other collections and implement the Collection interface.
package set

import (
	"fmt"
	"iter"
	"maps"

	"github.com/charbz/gophers/collection"
)

type Set[T comparable] struct {
	elements map[T]struct{}
}

func NewSet[T comparable](s ...[]T) *Set[T] {
	set := new(Set[T])
	set.elements = make(map[T]struct{})
	for _, slice := range s {
		for _, v := range slice {
			set.elements[v] = struct{}{}
		}
	}
	return set
}

// The following methods implement
// the Collection interface.

func (s *Set[T]) Add(v T) {
	s.elements[v] = struct{}{}
}

func (s *Set[T]) Length() int {
	return len(s.elements)
}

func (s *Set[T]) Random() T {
	for v := range s.elements {
		return v
	}
	panic(collection.EmptyCollectionError)
}

func (s *Set[T]) New(s2 ...[]T) collection.Collection[T] {
	return NewSet(s2...)
}

func (s *Set[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range s.elements {
			if !yield(k) {
				break
			}
		}
	}
}

func (s *Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.elements))
	for v := range s.elements {
		slice = append(slice, v)
	}
	return slice
}

// implement the Stringer interface
func (s *Set[T]) String() string {
	return fmt.Sprintf("Set(%T) %v", *new(T), s.ToSlice())
}

// The following methods are mostly syntatic sugar
// wrapping Collection functions to enable function chaining:
// i.e. set.Filter(f).Foreach(f2)

// Remove removes a value from the set.
func (s *Set[T]) Remove(v T) {
	delete(s.elements, v)
}

// Clone returns a copy of the collection. This is a shallow clone.
func (s *Set[T]) Clone() *Set[T] {
	return &Set[T]{
		elements: maps.Clone(s.elements),
	}
}

func (s *Set[T]) Count(f func(T) bool) int {
	return collection.Count(s, f)
}

// Contains returns true if the set contains the value.
func (s *Set[T]) Contains(v T) bool {
	_, ok := s.elements[v]
	return ok
}

// ContainsFunc returns true if the set contains a value that satisfies the predicate.
func (s *Set[T]) ContainsFunc(f func(T) bool) bool {
	for v := range s.Values() {
		if f(v) {
			return true
		}
	}
	return false
}

// Difference returns a new set containing the difference of the current set and the passed in set.
func (s *Set[T]) Diff(set *Set[T]) *Set[T] {
	newSet := s.Clone()
	for k := range set.Values() {
		delete(newSet.elements, k)
	}
	return newSet
}

// Equals returns true if the two sets contain the same elements.
func (s *Set[T]) Equals(s2 *Set[T]) bool {
	if s.Length() != s2.Length() {
		return false
	}
	for k := range s.Values() {
		if !s2.Contains(k) {
			return false
		}
	}
	return true
}

// Filter is an alias for collection.Filter
func (s *Set[T]) Filter(f func(T) bool) *Set[T] {
	return collection.Filter(s, f).(*Set[T])
}

// FilterNot is an alias for collection.FilterNot
func (s *Set[T]) FilterNot(f func(T) bool) *Set[T] {
	return collection.FilterNot(s, f).(*Set[T])
}

// Apply applies a function to each element in the set.
func (s *Set[T]) Apply(f func(T) T) *Set[T] {
	for k := range s.elements {
		v := f(k)
		s.Remove(k)
		s.Add(v)
	}
	return s
}

// ForAll is an alias for collection.ForAll
func (s *Set[T]) ForAll(f func(T) bool) bool {
	return collection.ForAll(s, f)
}

// IsEmpty returns true if the set is empty.
func (s *Set[T]) IsEmpty() bool {
	return s.Length() == 0
}

// Intersection returns a new set containing the intersection of the current set and the passed in set.
func (s *Set[T]) Intersection(s2 *Set[T]) *Set[T] {
	result := NewSet[T]()
	for k := range s2.elements {
		if _, ok := s.elements[k]; ok {
			result.Add(k)
		}
	}
	return result
}

// NonEmpty returns true if the set is not empty.
func (s *Set[T]) NonEmpty() bool {
	return s.Length() > 0
}

// Partition is an alias for collection.Partition
func (s *Set[T]) Partition(f func(T) bool) (*Set[T], *Set[T]) {
	left, right := collection.Partition(s, f)
	return left.(*Set[T]), right.(*Set[T])
}

// Union returns a new set containing the union of the current set and the passed in set.
func (s *Set[T]) Union(s2 *Set[T]) *Set[T] {
	result := s.Clone()
	for k := range s2.elements {
		result.Add(k)
	}
	return result
}
