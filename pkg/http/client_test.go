package http

import (
	"context"
	"net/http"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDefaultHttpClient(t *testing.T) {
	Convey("simple get should success", t, func(c C) {
		client := Global

		c.Convey("simple get should success", func(c C) {
			req := client.NewRequest()
			res := client.NewResponse()

			req.SetRequestURI("https://www.weimobcloud.com/")
			req.GetHeader().SetMethod("GET")
			req.GetHeader().SetContentTypeBytes(ContentTypeHtmlByte)

			err := client.Do(context.Background(), req, res)
			c.So(err, ShouldBeNil)
			c.So(res.StatusCode(), ShouldEqual, http.StatusOK)
			c.Println(string(res.Body())[:100])
		})

		c.Convey("post with null response should success", func(c C) {
			if true {
				// TODO skip
				return
			}
			req := client.NewRequest()
			res := client.NewResponse()

			body := `{"bosId":371629,"event":"onlineStatusUpdate","id":"test_id","model":1,"msgBody":"{\"isOnline\":true,\"goodsIdList\":[100020914999837]}","specsType":2,"test":true,"topic":"weimob_shop.goods","wid":"4464"}`
			req.SetRequestURI("http://127.0.0.1:8080/weimob/cloud/wos/message/receive")
			req.GetHeader().SetMethod("POST")
			req.GetHeader().SetContentTypeBytes(ContentTypeJsonByte)
			req.SetBody([]byte(body))

			err := client.Do(context.Background(), req, res)
			c.So(err, ShouldBeNil)
			c.So(res.StatusCode(), ShouldEqual, http.StatusOK)
			// c.So(res.Body(), ShouldBeEmpty)
			c.Println(string(res.Body()))
		})
	})
}

func TestUploadFile(t *testing.T) {
	Convey("simple get should success", t, func(c C) {
		client := NewHttpClient()

		c.Convey("upload file should success", func(cc C) {
			req := client.NewRequest()
			res := client.NewResponse()

			file, err := os.Open("./cat.jpg")
			cc.So(err, ShouldBeNil)
			defer file.Close()

			req.SetRequestURI("https://dopen.b.weimobqa.com/media/1_0/open/image/uploadImg")
			req.SetQueryString("accesstoken=ba48e3ba-4354-415c-8cac-d5c35e2dbceb")
			req.GetHeader().SetMethod("POST")
			req.GetHeader().SetContentTypeBytes(ContentTypeFromDataByte)
			req.SetFile("file", "cat", file)

			err = client.Do(context.Background(), req, res)
			c.So(err, ShouldBeNil)
			c.So(res.StatusCode(), ShouldEqual, http.StatusOK)
			c.Println(string(res.Body()))
		})
	})
}
