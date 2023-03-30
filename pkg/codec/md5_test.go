package codec

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMd5String(t *testing.T) {
	Convey("should return correct md5 string", t, func() {
		result := Md5String("hello world")
		So(result, ShouldEqual, "5eb63bbbe01eeed093cb22bb8f5acdc3")
	})
}
