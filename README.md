# Gophers - generic collection utils for Go
[![Go Reference](https://pkg.go.dev/badge/github.com/charbz/gophers.svg)](https://pkg.go.dev/github.com/charbz/gophers)
![gophers](https://github.com/user-attachments/assets/8ae33a16-e66a-43af-9412-f6bbd192d647)

Gophers is an awesome collections library for Go offering tons of utilities right out of the box.

Collection Types:
- **Sequence** : An ordered collection wrapping a Go slice. Great for fast random access.
- **ComparableSequence** : A Sequence of comparable elements. Offers extra functionality.
- **List** : An ordered collection wrapping a linked list. Great for fast insertion, removal, and implementing stacks and queues.
- **ComparableList** : A List of comparable elements. Offers extra functionality.
- **Set** : A hash set of unique elements.

Here's a few examples of what you can do:

## Quick Start

### Installation
```bash
go get github.com/charbz/gophers
```

### Using Generic Data Types

```go
import (
  "github.com/charbz/gophers/list"
)

type Foo struct {
  a int
  b string
}

foos := list.NewList([]Foo{
  {a: 1, b: "one"}, 
  {a: 2, b: "two"}, 
  {a: 3, b: "three"}, 
  {a: 4, b: "four"}, 
  {a: 5, b: "five"},
})

foos.Filter(func(f Foo) bool { return f.a%2 == 0 }) 
// List[Foo] {{2 two} {4 four}}

foos.Reject(func(f Foo) bool { return f.a%2 == 0 }) 
// List[Foo] {{1 one} {3 three} {5 five}}

foos.Find(func(f Foo) bool { return f.a == 3 }) 
// {a: 3, b: "three"}

foos.Partition(func(f Foo) bool { return len(f.b) == 3 })
// List[Foo] {{1 one} {2 two}} , List[Foo] {{3 three} {4 four} {5 five}}

foos.SplitAt(3) 
// List[Foo] {{1 one} {2 two} {3 three} {4 four}} , List[Foo] {{5 five}}

foos.Count(func(f Foo) bool { return f.a < 3 }) 
// 2

bars := foos.Concat(list.NewList([]Foo{{a: 1, b: "one"}, {a: 2, b: "two"}})) 
// List[Foo] {{1 one} {2 two} {3 three} {4 four} {5 five} {1 one} {2 two}}

bars.Distinct(func(i Foo, j Foo) bool { return i.a == j.a }) 
// List[Foo] {{1 one} {2 two} {3 three} {4 four} {5 five}}

foos.Apply(
  func(f Foo) Foo {
    f.a *= 2
    f.b += " * two"
    return f
  }
)
// List[Foo] {{2 one * two} {4 two * two} {6 three * two} {8 four * two} {10 five * two}}
```

### Comparable Collections

Comparable collections are collections with elements that can be compared to each other.
They offer all the same functionality as an ordered collection but provide additional convenience methods.

```go
import (
  "github.com/charbz/gophers/list"
)

nums := list.NewComparableList([]int{1, 1, 2, 2, 2, 3, 4, 5})

nums.Max() // 5

nums.Min() // 1

nums.Sum() // 20

nums.Distinct() // List[int] {1,2,3,4,5}

nums.Reverse() // List[int] {5,4,3,2,2,1,1}

nums.SplitAt(3) // List[int] {1,1,2,2}, List[int] {2,3,4,5}

nums.Take(3) // List[int] {1,1,2}

nums.TakeRight(3) // List[int] {3,4,5}

nums.Drop(3) // List[int] {2,2,3,4,5}

nums.DropRight(3) // List[int] {1,1,2,2,2}

nums.DropWhile(
  func(i int) bool { return i < 3 }, // List[int] {3,4,5}
)

nums.Diff(
  list.NewComparableList([]int{1, 2, 3}), // List[int] {4,5}
)

nums.Count(
  func(i int) bool { return i > 2 }, // 3
)
```

### Sets

Sets are collections of unique elements. They offer all the same functionality as an unordered collection
but provide additional methods for set operations.

```go
import (
  "github.com/charbz/gophers/set"
)

setA := set.NewSet([]string{"A", "B", "C", "A", "C", "A"}) // Set[string] {"A", "B", "C"}
setB := set.NewSet([]string{"A", "B", "D"})

setA.Intersection(setB) // Set[string] {"A", "B"}

setA.Union(setB) // Set[string] {"A", "B", "C", "D"}

setA.Diff(setB) // Set[string] {"C"}

setA.Apply(
  func(s string) string {
    s += "!"
    return s
  }, // Set[string] {"A!", "B!", "C!"}
)
```

### Map, Reduce, GroupBy...

You can use package functions such as Map, Reduce, GroupBy, and many more on any concrete collection type.

```go
import (
  "github.com/charbz/gophers/collection"
  "github.com/charbz/gophers/list"
  "github.com/charbz/gophers/sequence"
)

foos := sequence.NewSequence([]Foo{
	{a: 1, b: "one"},
	{a: 2, b: "two"},
	{a: 3, b: "three"},
	{a: 4, b: "four"},
	{a: 5, b: "five"},
})

collection.Map(foos, func(f Foo) string { return f.b }) //  ["one", "two", "three", "four", "five"] 

collection.Reduce(foos, func(acc string, f Foo) string { return acc + f.b }, "") // "onetwothreefourfive"

collection.Reduce(foos, func(acc int, f Foo) int { return acc + f.a }, 0) // 15

collection.GroupBy(foos, func(f Foo) int { return f.a % 2 }) // Map[int][]Foo { 0: [{2 two}, {4 four}], 1: [{1 one}, {3 three}, {5 five}]}
```

**Note:** Given that methods cannot define new type parameters in Go, any function that produces a new type, for example `Map(List[T], func(T) K) -> List[K]`, 
must be called as a function similar to the examples above and cannot be made into a method of the collection type i.e. `List[T].Map(func(T) K) -> List[K]`.
This is a minor inconvenience as it breaks the consistency of the API but is a limitation of the language.

### Iterator Methods

All collections implement methods that return iterators over the result as opposed to returning the result itself.
This is useful when combined with the `for ... range` syntax to iterate over the result.

```go
import (
  "github.com/charbz/gophers/list"
)

A := list.NewComparableList([]int{1, 2, 2, 3})
B := list.NewComparableList([]int{4, 5, 2})

for i := range A.Filtered(func(v int) bool { return v%2 == 0 }) {
  fmt.Printf("filtered %v\n", i)
}
// filtered 2
// filtered 2

for i := range A.Rejected(func(v int) bool { return v%2 == 0 }) {
  fmt.Printf("rejected %v\n", i)
}
// rejected 1
// rejected 3

for i := range A.Distincted() {
  fmt.Printf("distincted %v\n", i)
}
// distincted 1
// distincted 2
// distincted 3

for i := range A.Concatenated(B) {
  fmt.Printf("concatenated %v\n", i)
}
// concatenated 1
// concatenated 2
// concatenated 2
// concatenated 3
// concatenated 4
// concatenated 5
// concatenated 2

for i := range A.Diffed(B) {
  fmt.Printf("diffed %v\n", i)
}
// diffed 1
// diffed 3
```

### Sequence Operations

- `Add(element)` - Append element to sequence
- `All()` - Get iterator over all elements
- `At(index)` - Get element at index
- `Apply(function)` - Apply function to each element (mutates the original collection)
- `Backward()` - Get reverse iterator over elements
- `Clone()` - Create shallow copy of sequence
- `Concat(sequences...)` - Concatenates any passed sequences
- `Concatenated(sequence)` - Get iterator over concatenated sequence
- `Contains(predicate)` - Test if any element matches predicate
- `Corresponds(sequence, function)` - Test element-wise correspondence
- `Count(predicate)` - Count elements matching predicate
- `Dequeue()` - Remove and return first element
- `Diff(sequence, function)` - Get elements in first sequence but not in second
- `Diffed(sequence, function)` - Get iterator over elements in first sequence but not in second
- `Distinct(function)` - Get unique elements using equality function
- `Distincted()` - Get unique elements using equality comparison
- `Drop(n)` - Drop first n elements
- `DropRight(n)` - Drop last n elements
- `DropWhile(predicate)` - Drop elements while predicate is true
- `Enqueue(element)` - Add element to end
- `Equals(sequence, function)` - Test sequence equality using function
- `Exists(predicate)` - Test if any element matches predicate
- `Filter(predicate)` - Filter elements based on predicate
- `FilterNot(predicate)` - Inverse filter operation
- `Find(predicate)` - Find first matching element
- `FindLast(predicate)` - Find last matching element
- `ForAll(predicate)` - Test if predicate holds for all elements
- `Head()` - Get first element
- `Init()` - Get all elements except last
- `Intersect(sequence, function)` - Get elements present in both sequences
- `Intersected(sequence, function)` - Get iterator over elements present in both sequences
- `IsEmpty()` - Test if sequence is empty
- `Last()` - Get last element
- `Length()` - Get number of elements
- `New(slices...)` - Create new sequence
- `NewOrdered(slices...)` - Create new ordered sequence
- `NonEmpty()` - Test if sequence is not empty
- `Partition(predicate)` - Split sequence based on predicate
- `Pop()` - Remove and return last element
- `Push(element)` - Add element to end
- `Random()` - Get random element
- `Reverse()` - Reverse order of elements
- `Reject(predicate)` - Inverse filter operation
- `Rejected(predicate)` - Get iterator over elements rejected by predicate
- `Slice(start, end)` - Get subsequence from start to end
- `SplitAt(n)` - Split sequence at index n
- `String()` - Get string representation
- `Take(n)` - Get first n elements
- `TakeRight(n)` - Get last n elements
- `Tail()` - Get all elements except first
- `ToSlice()` - Convert to Go slice
- `Values()` - Get iterator over values

### ComparableSequence Operations

Inherits all operations from Sequence, but with the following additional operations:

- `Contains(element)` - Test if sequence contains element
- `Distinct()` - Get unique elements using equality comparison
- `Diff(sequence)` - Get elements in first sequence but not in second
- `Equals(sequence)` - Test sequence equality using equality comparison
- `Exists(element)` - Test if sequence contains element
- `IndexOf(element)` - Get index of first occurrence of element
- `LastIndexOf(element)` - Get index of last occurrence of element
- `Max()` - Get maximum element
- `Min()` - Get minimum element
- `Sum()` - Get sum of all elements

### List Operations

- `Add(element)` - Add element to end
- `All()` - Get iterator over index/value pairs
- `Apply(function)` - Apply function to each element
- `At(index)` - Get element at index
- `Backward()` - Get reverse iterator over index/value pairs
- `Clone()` - Create shallow copy
- `Concat(lists...)` - Concatenate multiple lists
- `Concatenated(list)` - Get iterator over concatenated list
- `Contains(predicate)` - Test if any element matches predicate
- `Corresponds(list, function)` - Test element-wise correspondence
- `Count(predicate)` - Count elements matching predicate
- `Dequeue()` - Remove and return first element
- `Diff(list, function)` - Get elements in first list but not in second
- `Diffed(list, function)` - Get iterator over elements in first list but not in second
- `Distinct(function)` - Get unique elements using equality function
- `Distincted()` - Get unique elements using equality comparison
- `Drop(n)` - Drop first n elements
- `DropRight(n)` - Drop last n elements
- `DropWhile(predicate)` - Drop elements while predicate is true
- `Enqueue(element)` - Add element to end
- `Equals(list, function)` - Test list equality using function
- `Exists(predicate)` - Test if any element matches predicate
- `Filter(predicate)` - Filter elements based on predicate
- `FilterNot(predicate)` - Inverse filter operation
- `Find(predicate)` - Find first matching element
- `FindLast(predicate)` - Find last matching element
- `ForAll(predicate)` - Test if predicate holds for all elements
- `Head()` - Get first element
- `Init()` - Get all elements except last
- `Intersect(list, function)` - Get elements present in both lists
- `Intersected(list, function)` - Get iterator over elements present in both lists
- `IsEmpty()` - Test if list is empty
- `Last()` - Get last element
- `Length()` - Get number of elements
- `New(slices...)` - Create new list
- `NewOrdered(slices...)` - Create new ordered list
- `NonEmpty()` - Test if list is not empty
- `Partition(predicate)` - Split list based on predicate
- `Pop()` - Remove and return last element
- `Push(element)` - Add element to end
- `Random()` - Get random element
- `Reverse()` - Reverse order of elements
- `Reject(predicate)` - Inverse filter operation
- `Rejected(predicate)` - Get iterator over elements rejected by predicate
- `Slice(start, end)` - Get sublist from start to end
- `SplitAt(n)` - Split list at index n
- `String()` - Get string representation
- `Take(n)` - Get first n elements
- `TakeRight(n)` - Get last n elements
- `Tail()` - Get all elements except first
- `ToSlice()` - Convert to Go slice
- `Values()` - Get iterator over values

### ComparableList Operations

Inherits all operations from List, but with the following additional operations:

- `Contains(value)` - Test if list contains value
- `Distinct()` - Get unique elements
- `Diff(list)` - Get elements in first list but not in second
- `Exists(value)` - Test if list contains value (alias for Contains)
- `Equals(list)` - Test list equality
- `IndexOf(value)` - Get index of first occurrence of value
- `LastIndexOf(value)` - Get index of last occurrence of value
- `Max()` - Get maximum element
- `Min()` - Get minimum element
- `Sum()` - Get sum of all elements


### Set Operations

- `Add(element)` - Add element to set
- `Apply(function)` - Apply function to each element
- `Clone()` - Create shallow copy of set
- `Contains(value)` - Test if set contains value
- `ContainsFunc(predicate)` - Test if set contains element matching predicate
- `Count(predicate)` - Count elements matching predicate
- `Diff(set)` - Get elements in first set but not in second
- `Diffed(set)` - Get iterator over elements in first set but not in second
- `Equals(set)` - Test set equality
- `Filter(predicate)` - Filter elements based on predicate
- `FilterNot(predicate)` - Inverse filter operation
- `ForAll(predicate)` - Test if predicate holds for all elements
- `Intersection(set)` - Get elements present in both sets
- `Intersected(set)` - Get iterator over elements present in both sets
- `IsEmpty()` - Test if set is empty
- `Length()` - Get number of elements
- `New(slices...)` - Create new set
- `NonEmpty()` - Test if set is not empty
- `Partition(predicate)` - Split set based on predicate
- `Random()` - Get random element
- `Remove(element)` - Remove element from set
- `Reject(predicate)` - Inverse filter operation
- `Rejected(predicate)` - Get iterator over elements rejected by predicate
- `String()` - Get string representation
- `ToSlice()` - Convert to Go slice
- `Union(set)` - Get elements present in either set
- `Unioned(set)` - Get iterator over elements present in either set
- `Values()` - Get iterator over values


### Collection Functions

The following package functions can be called on any collection, including Sequence, ComparableSequence, List, ComparableList, and Set.
- `Count(collection, predicate)` - Count elements matching predicate
- `Diff(collection)` - Get elements in first collection but not in second
- `Distinct(collection, function)` - Get unique elements
- `Filter(collection, predicate)` - Filter elements based on predicate
- `FilterNot(collection, predicate)` - Inverse filter operation
- `ForAll(collection, predicate)` - Test if predicate holds for all elements
- `GroupBy(collection, function)` - Group elements by key function
- `Intersect(collection1, collection2)` - Get elements present in both collections
- `Map(collection, function)` - Transform elements using function
- `MaxBy(collection, function)` - Get maximum element by comparison function
- `MinBy(collection, function)` - Get minimum element by comparison function
- `Partition(collection, predicate)` - Split collection based on predicate
- `Reduce(collection, function, initial)` - Reduce collection to single value

The package functions below can be called on ordered collections (Sequence, ComparableSequence, List, and ComparableList):
- `Corresponds(collection1, collection2, function)` - test whether values in collection1 map into values in collection2 by the given function
- `Drop(collection, n)` - Drop first n elements
- `DropRight(collection, n)` - Drop last n elements
- `DropWhile(collection, predicate)` - Drop elements while predicate is true
- `Find(collection, predicate)` - returns the index and value of the first element matching predicate
- `FindLast(collection, predicate)` - returns the index and value of the last element matching predicate
- `Head(collection)` - returns the first element in a collection
- `Init(collection)` - returns all elements excluding the last one
- `Last(collection)` - Get last element
- `ReduceRight(collection, function, initial)` - Right-to-left reduction
- `Reverse(collection)` - Reverse order of elements
- `ReverseMap(collection, function)` - Map elements in reverse order
- `SplitAt(collection, n)` - Split collection at index n
- `Tail(collection)` - Get all elements except first
- `Take(collection, n)` - Get first n elements
- `TakeRight(collection, n)` - Get last n elements

The following package functions return an iterator for the result:
- `Concatenated(collection1, collection2)` - Get iterator over concatenated collection
- `Diffed(collection1, collection2, function)` - Get iterator over elements in first collection but not in second
- `Intersected(collection1, collection2, function)` - Get iterator over elements present in both collections
- `Mapped(collection, function)` - Get iterator over elements transformed by function
- `Rejected(collection, predicate)` - Get iterator over elements rejected by predicate


## Contributing

Contributions are welcome! Feel free to submit a Pull Request.

If you have any ideas for new features or improvements, or would like to chat,

Feel free to reach out on [Discord: Gophers Project](https://discord.gg/vQ2dqQU6ve) or on [Reddit: r/gopherslib](https://www.reddit.com/r/gopherslib)
