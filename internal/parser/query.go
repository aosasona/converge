package parser

//go:generate go tool github.com/abice/go-enum --marshal

// ENUM(string,int,float,bool,list)
type FieldType int

const tagQueryParam = "query"

type QueryParam struct {
	name     string
	required bool
	// This will ideally default to our built-in regex validator but can be extended by users.
	validator func(string) error
	fieldType FieldType
}

// TODO: decide if I need this or want to expose it, or maybe just use reflect.StructField directly
type structField struct {
	identifier string
	fieldType  FieldType
	tags       map[string]string
}

// ExtractQueryParams uses reflection to extract query parameters from the given input struct.
// It checks for the `query` struct tag to identify fields that should be treated as query parameters.
// It falls back to the json tag if the query tag is not present.
// And if that is not present, a naive snake-case conversion of the field name is used.
//
// It supports the following field types:
// - string
// - int (no size variants)
// - bool
// - float64
// - slices of the above types (e.g., []string, []int, etc.)
//
// Nested structures are supported but are flattened using dot notation (e.g., `address.street`) to any of the above types.
//
// It returns a slice of QueryParam structs representing the extracted query parameters.
func ExtractQueryParams[T any](input T) ([]QueryParam, error) {
	panic("unimplemented")
}

func ExtractStructFields[T any](input T) ([]structField, error) {
	panic("unimplemented")
}
