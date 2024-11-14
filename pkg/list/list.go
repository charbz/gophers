// Copyright (c) 2024 Gophers. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package list

import (
	"fmt"
	"iter"

	"github.com/charbz/gophers/pkg/collection"
)

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewList[T any](s ...[]T) *List[T] {
	list := new(List[T])
	if len(s) == 0 {
		return list
	}
	for _, slice := range s {
		for _, v := range slice {
			list.Append(v)
		}
	}
	return list
}

// Implement the Collection interface.

func (l *List[T]) All() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := 0
		for node := l.head; node != nil; node = node.next {
			if !yield(i, node.value) {
				break
			}
			i++
		}
	}
}

func (l *List[T]) At(index int) T {
	if index < 0 || index >= l.size {
		panic(collection.IndexOutOfBoundsError)
	}
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.value
}

func (l *List[T]) Append(v T) {
	node := &Node[T]{value: v}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}
	l.size++
}

func (l *List[T]) Backward() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		i := l.size - 1
		for node := l.tail; node != nil; node = node.prev {
			if !yield(i, node.value) {
				break
			}
			i--
		}
	}
}

func (l *List[T]) Length() int {
	return l.size
}

func (l *List[T]) New(s ...[]T) collection.Collection[T] {
	return NewList(s...)
}

func (l *List[T]) Slice(start, end int) collection.Collection[T] {
	if start < 0 || end > l.size || start > end {
		panic(collection.IndexOutOfBoundsError)
	}
	list := &List[T]{}
	for i, v := range l.All() {
		if i < start {
			continue
		}
		if i >= start && i < end {
			list.Append(v)
		}
		if i >= end {
			break
		}
	}
	return list
}

func (l *List[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for node := l.head; node != nil; node = node.next {
			if !yield(node.value) {
				break
			}
		}
	}
}

func (l *List[T]) ToSlice() []T {
	slice := make([]T, 0, l.size)
	for v := range l.Values() {
		slice = append(slice, v)
	}
	return slice
}

// Implement the Stringer interface.

func (l *List[T]) String() string {
	return fmt.Sprintf("List(%T) %v", *new(T), l.ToSlice())
}
