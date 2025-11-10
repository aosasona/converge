package procedure

import (
	"regexp"
	"strings"
)

var (
	// Valid procedure name regex
	ReValidProcedureName = regexp.MustCompile(`(?m)^([a-zA-Z0-9]+)([_\.\-]?[a-zA-Z0-9]+)+$`)

	// Invalid characters in a procedure name
	ReAlphaNumeric = regexp.MustCompile(`[^a-zA-Z0-9]+`)

	// Multiple slashes regex
	ReIllegalSlash = regexp.MustCompile(`/+`)

	// Valid/common words associated with queries
	ReQueryWords = regexp.MustCompile(
		`(?i)(^(get|fetch|list|lookup|search|find|query|retrieve|show|view|read)\.)`,
	)

	// Valid/common words associated with mutations
	ReMutationWords = regexp.MustCompile(
		`(?i)(^(create|add|insert|update|upsert|edit|modify|change|delete|remove|destroy)\.)`,
	)
)

type baseProcedure[I, O any] struct {
	name string
	path Path
}

// Name implements Procedure.
func (b *baseProcedure[I, O]) Name() string {
	return b.name
}

// InputType implements Procedure.
func (b *baseProcedure[I, O]) InputType() any {
	return (*new(I))
}

// OutputType implements Procedure.
func (b *baseProcedure[I, O]) OutputType() any {
	return (*new(O))
}

// Path implements Procedure.
func (b *baseProcedure[I, O]) Path() Path {
	return b.path
}

// String implements Procedure.
func (b *baseProcedure[I, O]) String() string {
	return b.name
}

// WithPath implements Procedure.
func (b *baseProcedure[I, O]) WithPath(path string) Procedure[I, O] {
	path = normalizePath(path)

	_ = path

	panic("unimplemented")
}

// Generate a valid path from the procedure name/identifier.
// This is ideally used as a default path if none is provided by the user.
func (b *baseProcedure[I, O]) pathFromName() Path {
	var alias string

	// Replace all non-alphanumeric characters with slashes
	alias = ReAlphaNumeric.ReplaceAllString(b.name, "/")

	// Replace all multiple slashes with a single slash
	alias = ReIllegalSlash.ReplaceAllString(alias, "/")

	// Remove all words that are associable with the query type
	alias = ReQueryWords.ReplaceAllString(alias, "")

	// Remove all leading and trailing slashes
	alias = strings.TrimSpace(alias)
	alias = strings.TrimSuffix(alias, "/")
	if !strings.HasPrefix(alias, "/") {
		alias = "/" + alias
	}

	return Path{value: alias}
}
