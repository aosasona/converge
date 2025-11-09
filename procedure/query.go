package procedure

type queryProcedure[I, O any] struct {
	*baseProcedure[I, O]
}

// Type implements Procedure.
func (q *queryProcedure[I, O]) Type() ProcedureType {
	return ProcedureTypeQuery
}

var _ Procedure[any, any] = (*queryProcedure[any, any])(nil)
