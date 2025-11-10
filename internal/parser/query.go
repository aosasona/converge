package parser

const tagQueryParam = "query"

type QueryParam struct {
	name     string
	required bool
	// This will ideally default to our built-in regex validator but can be extended by users.
	validator func(string) error
}

// ExtractQueryParams uses reflection to extract query parameters from the given input struct.
// It checks for the `query` struct tag to identify fields that should be treated as query parameters.
// It falls back to the json tag if the query tag is not present.
// And if that is not present, a naive snake-case conversion of the field name is used.
func ExtractQueryParams[T any](input T) ([]QueryParam, error) {
	panic("unimplemented")
}
