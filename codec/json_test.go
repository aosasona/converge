package codec_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strings"
	"testing"

	"go.trulyao.dev/converge/codec"
)

type sample struct {
	ID    string `json:"id"`
	Count int    `json:"count"`
}

func TestJSON_Defaults_RoundTrip(t *testing.T) {
	j := codec.NewJSON[sample](nil) // defaults (EscapeHTML=true)

	in := sample{ID: "abc", Count: 42}

	var buf bytes.Buffer
	if err := j.Encode(in, &buf); err != nil {
		t.Fatalf("encode failed: %v", err)
	}

	// json.Encoder.Encode appends a trailing newline.
	if got := buf.String(); !strings.HasSuffix(got, "\n") {
		t.Fatalf("expected trailing newline, got: %q", got)
	}

	out, err := j.Decode(&buf)
	if err != nil {
		t.Fatalf("decode failed: %v", err)
	}
	if out != in {
		t.Fatalf("roundtrip mismatch: got %+v want %+v", out, in)
	}
}

func TestJSON_EscapeHTML_DefaultTrue(t *testing.T) {
	j := codec.NewJSON[map[string]string](nil) // EscapeHTML=true by default

	in := map[string]string{"x": "<script>alert(1)</script>"}
	var buf bytes.Buffer

	if err := j.Encode(in, &buf); err != nil {
		t.Fatalf("encode: %v", err)
	}
	got := buf.String()

	// With EscapeHTML=true, '<' and '>' are escaped.
	if !strings.Contains(got, `\u003cscript\u003e`) {
		t.Fatalf("expected HTML-escaped output, got: %q", got)
	}
}

func TestJSON_Indentation(t *testing.T) {
	j := codec.NewJSON[sample](&codec.JSONConfig{
		Indentation: &codec.JSONIndent{Indent: "  ", Prefix: ""},
	})

	var buf bytes.Buffer
	if err := j.Encode(sample{ID: "x", Count: 1}, &buf); err != nil {
		t.Fatalf("encode: %v", err)
	}
	if !strings.Contains(buf.String(), "\n  ") {
		t.Fatalf("expected indented output, got: %q", buf.String())
	}
}

func TestJSON_DisallowUnknownFields(t *testing.T) {
	j := codec.NewJSON[sample](&codec.JSONConfig{DisallowUnknownFields: true})

	const bad = `{"id":"x","count":7,"extra":"nope"}`
	_, err := j.Decode(strings.NewReader(bad))
	if err == nil {
		t.Fatal("expected error for unknown field, got nil")
	}
	// Optional: error mentions the field name
	if !strings.Contains(err.Error(), "extra") {
		t.Fatalf("expected error to mention 'extra', got: %v", err)
	}
}

func TestJSON_UseNumber(t *testing.T) {
	// Decode into a dynamic map so we can examine number types.
	j := codec.NewJSON[map[string]any](&codec.JSONConfig{UseNumber: true})

	const src = `{"big": 9007199254740993, "flt": 1.25}` // > 2^53
	out, err := j.Decode(strings.NewReader(src))
	if err != nil {
		t.Fatalf("decode: %v", err)
	}

	big, ok := out["big"].(json.Number)
	if !ok {
		t.Fatalf("expected json.Number for 'big', got %T (%v)", out["big"], out["big"])
	}
	if s := big.String(); s != "9007199254740993" {
		t.Fatalf("expected exact numeric string, got %q", s)
	}

	// float is also a json.Number when UseNumber is on
	if _, ok := out["flt"].(json.Number); !ok {
		t.Fatalf("expected json.Number for 'flt', got %T", out["flt"])
	}
}

func TestJSON_StreamRoundTrip_WithPipe(t *testing.T) {
	j := codec.NewJSON[sample](nil)

	in := sample{ID: "pipe", Count: 99}
	pr, pw := io.Pipe()

	encErr := make(chan error, 1)
	go func() {
		encErr <- j.Encode(in, pw)
		_ = pw.Close()
	}()

	out, decErr := j.Decode(pr)
	if err := <-encErr; err != nil {
		t.Fatalf("encode failed: %v", err)
	}
	if decErr != nil {
		t.Fatalf("decode failed: %v", decErr)
	}
	if out != in {
		t.Fatalf("roundtrip mismatch: got %+v want %+v", out, in)
	}
}

func TestJSON_ContentType(t *testing.T) {
	j := codec.NewJSON[sample](nil)
	if ct := j.ContentType(); ct != "application/json" {
		t.Fatalf("content type mismatch: %q", ct)
	}
}

func TestJSON_DecodeSyntaxError(t *testing.T) {
	j := codec.NewJSON[sample](nil)
	_, err := j.Decode(strings.NewReader("{not-json"))
	if err == nil {
		t.Fatal("expected decode error, got nil")
	}
	var syn *json.SyntaxError
	if !errors.As(err, &syn) {
		t.Fatalf("expected *json.SyntaxError, got %T: %v", err, err)
	}
}

func FuzzJSON_RoundTrip(f *testing.F) {
	type S struct {
		A string `json:"a"`
		B int    `json:"b"`
	}
	j := codec.NewJSON[S](nil)

	// Seeds
	f.Add("hello", 123)
	f.Add("<x>", -7)
	f.Add("\xa9", 0)  // raw © byte in string
	f.Add("\\xa9", 0) // the literal backslash sequence
	f.Add("\"q\"", 1)
	f.Add("", 0)

	f.Fuzz(func(t *testing.T, a string, b int) {
		in := S{A: a, B: b}

		// Encode -> Decode via your codec
		var buf bytes.Buffer
		if err := j.Encode(in, &buf); err != nil {
			t.Fatalf("encode: %v", err)
		}
		got, err := j.Decode(&buf)
		if err != nil {
			t.Fatalf("decode: %v (src: %q)", err, buf.String())
		}

		// Sanitize the INPUT the same way encoding/json would:
		// invalid UTF-8 -> U+FFFD, etc.
		var sanitizedIn S
		encIn, _ := json.Marshal(in)
		_ = json.Unmarshal(encIn, &sanitizedIn)

		// Now they should be byte-for-byte equal as Go values.
		if got != sanitizedIn {
			// For easier triage, show canonical JSON too.
			wantJSON, _ := json.Marshal(sanitizedIn)
			gotJSON, _ := json.Marshal(got)
			t.Fatalf("roundtrip mismatch:\n  want=%s\n   got=%s", wantJSON, gotJSON)
		}
	})
}
