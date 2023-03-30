package auth

import (
	"context"
	"github.com/weimob-tech/go-project-base/pkg/hook"
	"github.com/weimob-tech/go-project-base/pkg/wlog"
)

type Store interface {
	GetCCToken(ctx context.Context, product, shopId, shopType string) (*OAuthResponse, error)
	GetProductCCToken(ctx context.Context, cid, cse, shopId, shopType string) (*OAuthResponse, error)
}

var DefaultStore Store

func init() {
	hook.AddPostStartHook(func() {
		if DefaultStore == nil {
			DefaultStore = NewHttpStore()
		}
	})
}

func GetCCToken(ctx context.Context, product, shopId, shopType string) (token string) {
	resp, err := DefaultStore.GetCCToken(ctx, product, shopId, shopType)
	if err != nil {
		wlog.Errorf("get access token failed, err: %v", err)
		return ""
	}
	return resp.AccessToken
}

func GetClientCCToken(ctx context.Context, clientId, clientSecret, shopId, shopType string) (token string) {
	resp, err := DefaultStore.GetProductCCToken(ctx, clientId, clientSecret, shopId, shopType)
	if err != nil {
		wlog.Errorf("get access token failed, err: %v", err)
		return ""
	}
	return resp.AccessToken
}
