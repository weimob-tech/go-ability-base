package codec

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

type Md5Codec interface {
	Md5Raw([]byte) [16]byte
	Md5String(string) string
}

var Md5 Md5Codec

type defaultMd5Codec struct{}

func (d defaultMd5Codec) Md5Raw(data []byte) [16]byte {
	return md5.Sum(data)
}

func (d defaultMd5Codec) Md5String(data string) string {
	sum := md5.Sum([]byte(data))
	return base64.StdEncoding.EncodeToString(sum[:])
}

func init() {
	Md5 = &defaultMd5Codec{}
}

func Md5Raw(data []byte) [16]byte {
	return Md5.Md5Raw(data)
}

func Md5String(data string) string {
	return fmt.Sprintf("%x", Md5.Md5Raw([]byte(data)))
}
