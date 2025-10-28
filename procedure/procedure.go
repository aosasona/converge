package procedure

import "fmt"

//go:generate go tool github.com/abice/go-enum --marshal

// ENUM(query,mutation,live)
type ProcedureType string

type Handler[In, Out any] func(c *Context, input In) (Out, error)

// TODO: define procedure interface
type Procedure[I, O any] interface {
	fmt.Stringer

	// Name returns the name of the procedure
	Name() string

	// Type returns the type of the procedure (query, mutation, live)
	// A query is a read-only operation
	// A mutation is a write operation that may change state
	// A live procedure is a long-lived operation that streams updates (usually backed by websockets; implementation might differ)
	Type() ProcedureType

	// Handle processes the input and returns the output or an error
	Handle(c *Context, input I) (O, error)

	// The path to use for the REST or websocket endpoint (e.g. "/user/profile")
	Path() string

	// WithPath sets the path for the procedure and returns a modified copy
	WithPath(path string) Procedure[I, O]

	// InputType returns an empty instance of the input type
	InputType() any

	// OutputType returns an empty instance of the output type
	OutputType() any
}
