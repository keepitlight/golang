package encoding

import "github.com/keepitlight/golang/encoding/form"

// FormMarshal marshals a value into a URL-encoded form.
//
// See encoding/form.Marshal.
func FormMarshal(v interface{}) ([]byte, error) {
	return form.Marshal(v)
}

// FormUnmarshal unmarshals a URL-encoded form into a value.
//
// See encoding/form.Unmarshal.
func FormUnmarshal(data []byte, v interface{}) error {
	return form.Unmarshal(data, v)
}
