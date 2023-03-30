package codec

func ToJson[T any](v *T) string {
	data, err := Json.MarshalString(v)
	if err != nil {
		return ""
	}
	return data
}

func FromJson[T any](s string) (v *T) {
	v = new(T)
	err := Json.UnmarshalString(s, v)
	if err != nil {
		return nil
	}
	return v
}
func ToJsonByte[T any](v *T) []byte {
	data, err := Json.Marshal(v)
	if err != nil {
		return nil
	}
	return data
}

func FromJsonByte[T any](s []byte) (v *T) {
	v = new(T)
	err := Json.Unmarshal(s, v)
	if err != nil {
		return nil
	}
	return v
}
