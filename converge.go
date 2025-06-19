package converge

import (
	"go.trulyao.dev/converge/codegen"
)

type flags struct {
	// Whether debug mode is enabled or not
	debug bool

	// Whether to trap panics or let them bubble up
	trapPanic bool
}

type Converge struct {
	flags flags

	// The codegen options for the preferred targets (e.g. TypeScript)
	codegenOptions []codegen.Options
}

// TODO: add options to pass into the New function
// func New() *Converge {
// 	return &Converge{}
// }
