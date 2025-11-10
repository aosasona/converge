package procedure

type mutationProcedure[I, O any] struct {
	*baseProcedure[I, O]
}

// Handle implements Procedure.
func (q *mutationProcedure[I, O]) Handle(c *Context, input I) (O, error) {
	panic("unimplemented")
}

// Type implements Procedure.
func (q *mutationProcedure[I, O]) Type() ProcedureType {
	return ProcedureTypeMutation
}

var _ Procedure[any, any] = (*mutationProcedure[any, any])(nil)
