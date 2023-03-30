package auth

import "github.com/weimob-tech/go-project-base/pkg/config"

type ClientInfo struct {
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
}

// GetClientInfo get client from config by product name,
// for example product name is "shop-a", and the config key would be
//
// > weimob.cloud.shop-a.clientId = foo
// > weimob.cloud.shop-a.clientSecret = bar
func GetClientInfo(product string) *ClientInfo {
	return &ClientInfo{
		ClientId:     config.GetString("weimob.cloud." + product + ".clientId"),
		ClientSecret: config.GetString("weimob.cloud." + product + ".clientSecret"),
	}
}
