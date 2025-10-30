package procedure

type baseProcedure[I, O any] struct {
	name string
	path string
}

// Handle implements Procedure.
func (b *baseProcedure[I, O]) Handle(c *Context, input I) (O, error) {
	panic("unimplemented")
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
func (b *baseProcedure[I, O]) Path() string {
	return b.path
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
	path = normalizePath(path)

	_ = path

	panic("unimplemented")
}

var _ Procedure[any, any] = (*baseProcedure[any, any])(nil)
