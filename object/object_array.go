package object

import (
	"bytes"
	"strings"
)

// Array wraps Object array and implements Object interface.
type Array struct {
	// Elements holds the individual members of the array we're wrapping.
	Elements []Object

	// offset holds our iteration-offset.
	offset int
}

// Type returns the type of this object.
func (ao *Array) Type() Type {
	return ARRAY
}

// Inspect returns a string-representation of the given object.
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := make([]string, 0)
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// True returns whether this object wraps a true-like value.
//
// Used when this object is the conditional in a comparison, etc.
func (ao *Array) True() bool {
	return (len(ao.Elements) != 0)
}

// ToInterface converts this object to a go-interface, which will allow
// it to be used naturally in our sprintf/printf primitives.
//
// It might also be helpful for embedded users.
func (ao *Array) ToInterface() interface{} {

	res := make([]interface{}, len(ao.Elements))

	for _, v := range ao.Elements {
		res = append(res, v.ToInterface())
	}

	return res
}

// Reset implements the Iterable interface, and allows the contents
// of the array to be reset to allow re-iteration.
func (ao *Array) Reset() {
	ao.offset = 0
}

// Next implements the Iterable interface, and allows the contents
// of our array to be iterated over.
func (ao *Array) Next() (Object, int, bool) {
	if ao.offset < len(ao.Elements) {
		ao.offset++

		element := ao.Elements[ao.offset-1]
		return element, ao.offset - 1, true
	}

	return nil, 0, false
}
