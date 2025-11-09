package procedure

type mutationProcedure[I, O any] struct {
	*baseProcedure[I, O]
}

// Type implements Procedure.
func (q *mutationProcedure[I, O]) Type() ProcedureType {
	return ProcedureTypeMutation
}

var _ Procedure[any, any] = (*mutationProcedure[any, any])(nil)
