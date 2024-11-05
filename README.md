# Gophers - functional utilities for generic collections

A lightweight Go library providing functional programming utilities for working with generic collections. This library offers a clean, type-safe API for common collection operations inspired by functional programming patterns.

## Features

- Generic collection types supporting any data type
- Functional programming utilities (Map, Filter, Reduce, etc.)
- Zero dependencies (uses only Go standard library)
- Fully type-safe operations
- MIT Licensed

## Installation

```bash
go get github.com/charbz/gophers
```

## Quick Start

```go
import (
  "github.com/charbz/gophers/pkg/collections"
  "github.com/charbz/gophers/pkg/utils"
)

// Create a new collection
numbers := collections.NewCollection([]int{1, 2, 3, 4, 5})

// Basic operations
numbers.Head() // 1

numbers.Last() // 5

numbers.Take(3) // [1,2,3]

numbers.TakeRight(2) // [4,5]

numbers.Drop(2) // [3,4,5]

numbers.Filter(func(n int) bool {
  return n%2 == 0
}) // [2,4]

numbers.FilterNot(func(n int) bool {
  return n%2 == 0
}) // [1,3,5]
```

## Core Features

### Collection Operations

- `Head()` - Get the first element
- `Last()` - Get the last element
- `Tail()` - Get all elements except the first
- `Init()` - Get all elements except the last
- `Take(n)` - Get first n elements
- `TakeRight(n)` - Get last n elements
- `Drop(n)` - Drop first n elements
- `DropRight(n)` - Drop last n elements
- `IsEmpty()` - Check if collection is empty
- `Length()` - Get collection size

### Functional Utilities

- `Filter(predicate)` - Filter elements based on predicate
- `FilterNot(predicate)` - Inverse filter operation
- `Map(function)` - Transform elements using provided function
- `Reduce(function, initial)` - Reduce collection to single value
- `Partition(predicate)` - Split collection into two based on predicate

## Examples

### Map, Reduce, Partition

```go
utils.Map(numbers.ToSlice(), func(n int) int {
  return n * 2
}) // [2,4,6,8,10]

utils.Reduce(numbers.ToSlice(), func(acc int, n int) int {
    return acc + n
}, 0) // 15

utils.Partition(numbers.ToSlice(), func(n int) bool {
    return n%2 == 0
}) // [2, 4]  [1, 3, 5]
```

## Contributing

Contributions are welcome! Feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
