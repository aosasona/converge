package codec

import "io"

type JSON[T any] struct{}

// Decode implements Decoder.
func (j *JSON[T]) Decode(reader io.Reader) (T, error) {
	panic("unimplemented")
}

// Encode implements Encoder.
func (j *JSON[T]) Encode(data T, writer io.Writer) error {
	panic("unimplemented")
}

// ContentType implements Serializer.
func (j *JSON[T]) ContentType() string {
	return "application/json"
}

var (
	_ Encoder[any] = (*JSON[any])(nil)
	_ Decoder[any] = (*JSON[any])(nil)
)
