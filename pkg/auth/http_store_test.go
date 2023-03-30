package auth

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-base/pkg/hook"
	"testing"
)

func init() {
	config.SetDefault("client.log.lvl", "body")
	config.SetDefault("client.schema", "http")
	config.SetDefault("client.oauth.domain", "dopen.b.weimobqa.com")
	config.SetDefault("weimob.cloud.foo.clientId", "477FD526146E55C173E128C3596DC790")
	config.SetDefault("weimob.cloud.foo.clientSecret", "D3C10901C731115D8DF6DA38D19CFDD1")

	hook.ApplyPostStartHook()
}

func TestGetCCToken(t *testing.T) {
	Convey("when get access token", t, func(c C) {
		root := context.Background()

		c.Convey("get cc token should success", func(cc C) {
			token := GetCCToken(root, "foo", "", "")
			cc.So(token, ShouldNotBeEmpty)
			_, _ = cc.Println(token)
		})
	})
}
