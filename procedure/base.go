package procedure

type baseProcedure[I, O any] struct{}

// Handle implements Procedure.
func (b *baseProcedure[I, O]) Handle(c *Context, input I) (O, error) {
	panic("unimplemented")
}

// InputType implements Procedure.
func (b *baseProcedure[I, O]) InputType() any {
	panic("unimplemented")
}

// Name implements Procedure.
func (b *baseProcedure[I, O]) Name() string {
	panic("unimplemented")
}

// OutputType implements Procedure.
func (b *baseProcedure[I, O]) OutputType() any {
	panic("unimplemented")
}

// Path implements Procedure.
func (b *baseProcedure[I, O]) Path() string {
	panic("unimplemented")
}

// String implements Procedure.
func (b *baseProcedure[I, O]) String() string {
	panic("unimplemented")
}

// Type implements Procedure.
func (b *baseProcedure[I, O]) Type() ProcedureType {
	panic("unimplemented")
}

// WithPath implements Procedure.
func (b *baseProcedure[I, O]) WithPath(path string) Procedure[I, O] {
	panic("unimplemented")
}

var _ Procedure[any, any] = (*baseProcedure[any, any])(nil)
