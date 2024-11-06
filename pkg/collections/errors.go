package collections

import "fmt"

type CollectionError struct {
	Code int
	Msg  string
}

func (e *CollectionError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Msg)
}

var (
	EmptyCollectionError = &CollectionError{Code: 100, Msg: "invalid operation on an empty collection"}
)
