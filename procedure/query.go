package procedure

import "go.trulyao.dev/converge/internal/parser"

// TODO: pre-parse and cache query params
type queryProcedure[I, O any] struct {
	*baseProcedure[I, O]
	queryParams []parser.QueryParam
}

func Query[I, O any](name string) (*queryProcedure[I, O], error) {
	base := &baseProcedure[I, O]{name: name}
	base.path = base.pathFromName()

	queryParams, err := parser.ExtractQueryParams(new(I))
	if err != nil {
		return nil, err
	}
	return &queryProcedure[I, O]{
		baseProcedure: base,
		queryParams:   queryParams,
	}, nil
}

// Type implements Procedure.
func (q *queryProcedure[I, O]) Type() ProcedureType {
	return ProcedureTypeQuery
}

// Handle implements Procedure.
func (q *queryProcedure[I, O]) Handle(c *Context, input I) (O, error) {
	panic("unimplemented")
}

var _ Procedure[any, any] = (*queryProcedure[any, any])(nil)
