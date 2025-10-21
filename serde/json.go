package serde

type Json[T any] struct{}

// Deserialize implements Deserializer.
func (j *Json[T]) Deserialize(data []byte) (*T, error) {
	panic("unimplemented")
}

// Serialize implements Serializer.
func (j *Json[T]) Serialize(v any) ([]byte, error) {
	panic("unimplemented")
}

// ContentType implements Serializer.
func (j *Json[T]) ContentType() string {
	return "application/json"
}

var (
	_ Serializer        = (*Json[any])(nil)
	_ Deserializer[any] = (*Json[any])(nil)
)
