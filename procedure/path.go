package procedure

import "fmt"

// Path is an abstract representation of a procedure Path.
//
// NOTE: in the future, this may be extended to support more complex Path structures, including parameters and wildcards.
type Path struct {
	fmt.Stringer

	value string
}

// NewPath creates a new Path instance from the given string value.
// It will try as much as possible to extract a valid Path from the input or return an error.
func NewPath(value string) (Path, error) {
	// TODO: parse path into segments, parameters, etc.
	return Path{value: normalizePath(value)}, nil
}

func (p Path) String() string {
	return p.value
}

func normalizePath(path string) string {
	if path == "" {
		return "/"
	}

	if path[0] != '/' {
		path = "/" + path
	}

	if len(path) > 1 && path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}

	return path
}
