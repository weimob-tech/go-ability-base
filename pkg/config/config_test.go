package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestSyncMapStoreBindEnv(t *testing.T) {
	_ = os.Setenv("foo", "bar")
	Convey("should bind to specified env", t, func(c C) {
		store := NewSyncMapStore()
		err := store.BindEnv("foo")

		So(err, ShouldBeNil)
		So(store.GetString("foo"), ShouldEqual, "bar")

		c.Convey("should set new value", func() {
			store.Set("foo", "baz")
			So(store.GetString("foo"), ShouldEqual, "baz")
		})
	})
}

func TestSyncMapStoreGetStringMap(t *testing.T) {
	Convey("should get string map", t, func(c C) {
		store := NewSyncMapStore()
		store.SetDefault("test.foo", "a")
		store.SetDefault("test.bar", "b")

		var expect = map[string]interface{}{"foo": "a", "bar": "b"}

		c.Convey("should get string map", func(cc C) {
			cc.So(store.GetStringMap("test"), ShouldResemble, expect)
		})

		c.Convey("should get string map cached", func(cc C) {
			cached := Cached(store)
			cc.So(cached.GetStringMap("test"), ShouldResemble, expect)
		})
		c.Convey("should set new value", func(cc C) {
			store.Set("foo", "bar")
			cached := Cached(store)
			cached.Set("foo", "baz")
			So(store.GetString("foo"), ShouldEqual, "baz")
		})
		c.Convey("should override default", func(cc C) {
			store.SetDefault("tar", "ball")
			cached := Cached(store)
			So(cached.GetString("tar"), ShouldEqual, "ball")
			cached.Set("tar", "tall")
			So(store.GetString("tar"), ShouldEqual, "tall")
		})
	})
}
