package codec

import (
	"encoding/json"
	"io"
)

type JSONIndent struct {
	Prefix string
	Indent string
}

type (
	JSONConfig struct {
		EscapeHTML  bool        // default: true
		Indentation *JSONIndent // default: no indentation

		DisallowUnknownFields bool // default: false
		UseNumber             bool // default: false
	}

	JSON[T any] struct {
		escapeHTML            bool
		indent                string
		prefix                string
		disallowUnknownFields bool
		useNumber             bool
	}
)

func NewJSON[T any](config *JSONConfig) *JSON[T] {
	j := &JSON[T]{
		escapeHTML: true,
	}

	if config != nil {
		j.escapeHTML = config.EscapeHTML
		j.disallowUnknownFields = config.DisallowUnknownFields
		j.useNumber = config.UseNumber

		if config.Indentation != nil {
			j.indent = config.Indentation.Indent
			j.prefix = config.Indentation.Prefix
		}
	}

	return j
}

// Decode implements Decoder.
func (j *JSON[T]) Decode(reader io.Reader) (T, error) {
	var result T
	decoder := json.NewDecoder(reader)

	if j.disallowUnknownFields {
		decoder.DisallowUnknownFields()
	}

	if j.useNumber {
		decoder.UseNumber()
	}

	if err := decoder.Decode(&result); err != nil {
		var zeroValue T
		return zeroValue, err
	}

	return result, nil
}

// Encode implements Encoder.
func (j *JSON[T]) Encode(data T, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encoder.SetEscapeHTML(j.escapeHTML)
	if j.indent != "" || j.prefix != "" {
		encoder.SetIndent(j.prefix, j.indent)
	}
	return encoder.Encode(data)
}

// ContentType implements Serializer.
func (j *JSON[T]) ContentType() string {
	return "application/json"
}

var (
	_ Encoder[any] = (*JSON[any])(nil)
	_ Decoder[any] = (*JSON[any])(nil)
)
