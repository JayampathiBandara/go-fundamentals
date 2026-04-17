package aggregatedatatype

import (
	"fmt"
	"go-fundamentals/common"
)

func MapDeclaration() {
	// Maps in Go are unordered collections of key-value pairs.
	// They are implemented as hash tables and provide efficient lookups, insertions, and deletions.
	// The keys in a map must be of a type that is comparable (e.g., strings, integers),
	// while the values can be of any type.

	var m map[string]int
	fmt.Println(m)

	m = map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map ->", m)
	fmt.Println(`m["foo"]:`, m["foo"])

	m["bar"] = 99
	delete(m, "foo")
	m["baz"] = 418
	fmt.Println("map ->", m)
	fmt.Println(`m["foo"]:`, m["foo"])
	v, ok := m["foo"]

	fmt.Println(`Accessing not existing Key- m["foo"]:`, v, ok)

	v, ok = m["bar"]
	fmt.Println(`Accessing existing Key- m["bar"]:`, v, ok)
	common.PrintLine("=")
}
