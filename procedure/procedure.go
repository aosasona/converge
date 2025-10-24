package procedure

type Handler[In, Out any] func(c *Context, input In) (Out, error)

// TODO: define procedure interface
type Procedure interface{}
