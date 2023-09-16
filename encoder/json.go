package encoder

import "encoding/json"

type encoder struct {
}

func New() *encoder {

	return &encoder{}
}

func (e *encoder) Encode(jsonString string, object any) error {

	return json.Unmarshal([]byte(jsonString), object)
}
