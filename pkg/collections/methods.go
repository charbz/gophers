// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// This file contains all the public methods belonging to a Collection.
// Any new functionality "foo" that can be expressed as collection.foo() should be placed here.
//
// As a principle, when a method "foo" updates the underlying slice, it is critical to return
// a new Collection as a result in order to enable function chaining i.e.:
//
//   collection.Take(..).Filter(..).ForEach(..)

package collections

import (
	"github.com/charbz/gophers/pkg/utils"
)

// At returns the element at the specified index.
func (c *Collection[T]) At(index int) T {
	return c.elements[index]
}

// ToSlice returns the underlying slice.
func (c *Collection[T]) ToSlice() []T {
	return c.elements
}

// IsEmpty returns true if the Collection contains 0 elements.
func (c *Collection[T]) IsEmpty() bool {
	return len(c.elements) == 0
}

// NonEmpty returns true if the Collection contains at least 1 element.
func (c *Collection[T]) NonEmpty() bool {
	return len(c.elements) > 0
}

// Length returns the number of elements in the Collection.
func (c *Collection[T]) Length() int {
	return len(c.elements)
}

// Head returns the first element in a Collection and a nil error.
// If the collection is empty, it returns the zero value and an error.
//
// example usage:
//
//	c := NewCollection([]string{"A","B","C"})
//	c.Head()
//
// output:
//
//	"A", nil
func (c *Collection[T]) Head() (T, error) {
	if c.IsEmpty() {
		return c.zeroT(), emptyCollectionError
	}
	return c.elements[0], nil
}

// Last returns the last element in the Collection and a nil error.
// If the collection is empty, it returns the zero value and an error.
//
// example usage:
//
//	c := NewCollection([]string{"A","B","C"})
//	c.Last()
//
// output:
//
//	"C", nil
func (c *Collection[T]) Last() (T, error) {
	if c.IsEmpty() {
		return c.zeroT(), emptyCollectionError
	}
	return c.elements[len(c.elements)-1], nil
}

// Tail returns a new collection containing all elements excluding the first one.
//
// example usage:
//
//	c := NewCollection([]int{1,2,3,4,5,6})
//	c.Tail()
//
// output:
//
//	[2,3,4,5,6]
func (c *Collection[T]) Tail() *Collection[T] {
	if c.IsEmpty() {
		return c
	}
	return &Collection[T]{
		c.elements[1:],
	}
}

// Init returns a collection containing all elements excluding the last one.
//
// example usage:
//
//	c := NewCollection([]int{1,2,3,4,5,6})
//	c.Tail()
//
// output:
//
//	[1,2,3,4,5]
func (c *Collection[T]) Init() *Collection[T] {
	if c.IsEmpty() {
		return c
	}
	return &Collection[T]{
		c.elements[0 : len(c.elements)-1],
	}
}

// Take returns a new collection containing the first n elements.
//
// example usage:
//
//	[c := NewCollection([]int{1,2,3,4,5,6})
//	c.Take(3)
//
// output:
//
//	[1,2,3]
func (c *Collection[T]) Take(n int) *Collection[T] {
	if n <= 0 {
		return c.zero()
	}
	return &Collection[T]{
		c.elements[0:min(n, c.Length())],
	}
}

// TakeRight returns a new collection containing the last n elements.
//
// example usage:
//
//	c := NewCollection([]int{1,2,3,4,5,6})
//	c.TakeRight(3)
//
// output:
//
//	[4,5,6]
func (c *Collection[T]) TakeRight(n int) *Collection[T] {
	if n <= 0 {
		return c.zero()
	}
	return &Collection[T]{
		c.elements[max(c.Length()-n, 0):],
	}
}

// Drop returns a new collection with the first n elements removed.
//
// example usage:
//
//	c := NewCollection([]int{1,2,3,4,5,6})
//	c.Drop(2)
//
// output:
//
//	[3,4,5,6]
func (c *Collection[T]) Drop(n int) *Collection[T] {
	if n <= 0 {
		return c
	} else if n >= c.Length() {
		return c.zero()
	}
	return &Collection[T]{
		c.elements[n:],
	}
}

// DropRight returns a collection with the last n elements removed.
//
// example usage:
//
//	c := NewCollection([]int{1,2,3,4,5,6})
//	c.DropRight(2)
//
// output:
//
//	[1,2,3,4]
func (c *Collection[T]) DropRight(n int) *Collection[T] {
	if n <= 0 {
		return c
	} else if n >= c.Length() {
		return c.zero()
	}
	return &Collection[T]{
		c.elements[0 : c.Length()-n],
	}
}

// Filter takes a filtering function as input and returns a new collection
// containing all the elements that match the filter.
//
// example usage:
//
//	c := NewCollection([]int{1,2,3,4,5,6})
//	c.Filter(func(i int) bool {
//	  return i%2==0
//	})
//
// output:
//
//	[2,4,6]
func (c *Collection[T]) Filter(f func(T) bool) *Collection[T] {
	return &Collection[T]{
		utils.Filter(c.elements, f),
	}
}

// FilterNot takes a filtering function as input and returns a new collection
// containing all the elements that do not match the filter.
//
// example usage:
//
//	c := NewCollection([]int{1,2,3,4,5,6})
//	c.FilterNot(func(i int) bool {
//	  return i%2==0
//	})
//
// output:
//
//	[1,3,5]
func (c *Collection[T]) FilterNot(f func(T) bool) *Collection[T] {
	return &Collection[T]{
		utils.FilterNot(c.elements, f),
	}
}

// Partition takes a partitioning function as input and returns two collections,
// the first one contains the elements that match the partitioning condition,
// the second one contains the rest of the elements.
//
// example usage:
//
//	c := NewCollection([]int{1,2,3,4,5,6})
//	c.Partition(func(i int) bool {
//	  return i%2==0
//	})
//
// output:
//
//	[2,4,6], [1,3,5]
func (c *Collection[T]) Partition(f func(T) bool) (*Collection[T], *Collection[T]) {
	left, right := utils.Partition(c.elements, f)
	return &Collection[T]{left}, &Collection[T]{right}
}

// ForEach takes a function as input and applies the function
// to each element in the collection.
//
// example usage:
//
//	c.ForEach(func(t Task) {
//	  t.run()
//	})
func (c *Collection[T]) ForEach(f func(T)) *Collection[T] {
	for v := range c.Values() {
		f(v)
	}
	return c
}

// Reverse returns a new collection containing all elements in reverse
//
// example usage:
//
//	c := NewCollection([]int{1,2,3,4,5,6})
//	c.Reverse
//
// output:
//
//	[6,5,4,3,2,1]
func (c *Collection[T]) Reverse() *Collection[T] {
	elements := make([]T, 0, len(c.elements))
	for i := len(c.elements) - 1; i >= 0; i-- {
		elements = append(elements, c.elements[i])
	}
	return &Collection[T]{
		elements,
	}
}
