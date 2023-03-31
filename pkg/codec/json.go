package codec

import (
	"github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

func init() {
	extra.RegisterFuzzyDecoders()
}

type JsonEncoder interface {
	Marshal(any) ([]byte, error)
	MarshalString(any) (string, error)
}

type JsonDecoder interface {
	Unmarshal([]byte, any) error
	UnmarshalString(string, any) error
}

type JsonCodec interface {
	JsonEncoder
	JsonDecoder
}

var (
	json = jsoniter.ConfigDefault
	Json JsonCodec
)

func init() {
	Json = &defaultJsonCodec{}
}

type defaultJsonCodec struct{}

func (d defaultJsonCodec) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (d defaultJsonCodec) MarshalString(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (d defaultJsonCodec) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (d defaultJsonCodec) UnmarshalString(s string, a any) error {
	return json.Unmarshal([]byte(s), a)
}
