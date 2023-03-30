package auth

import (
	"context"
	"fmt"
	"github.com/weimob-tech/go-project-base/pkg/codec"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-base/pkg/hook"
	"github.com/weimob-tech/go-project-base/pkg/http"
	"github.com/weimob-tech/go-project-base/pkg/wlog"
	"net/url"
	"strings"
	"time"
)

var (
	requestUri   string
	blankRequest = []byte("{}")
)

type authRequest struct {
	grantType string
	shopId    string
	shopType  string
}

type OAuthResponse struct {
	ExpiresAt                 time.Time
	ExpiresIn                 int64  `json:"expires_in"`
	AccessToken               string `json:"access_token"`
	TokenType                 string `json:"token_type"`
	Scope                     string `json:"scope"`
	BusinessOperationSystemID string `json:"business_operation_system_id"`
	PublicAccountID           string `json:"public_account_id"`
	BusinessID                string `json:"business_id"`
}

func init() {
	hook.AddPostStartHook(SetupAuthManager)
}

func SetupAuthManager() {
	config.SetDefault("client.oauth.path", "fuwu/b/oauth2/token")

	baseUrl := config.GetString("client.schema") + "://" + config.GetString("client.oauth.domain")
	baseUrl = strings.Trim(baseUrl, "/")
	basePath := config.GetString("client.oauth.path")
	basePath = strings.Trim(basePath, "/")
	requestUri = baseUrl + "/" + basePath
}

func NewHttpStore() Store {
	return &httpStore{http.Global}
}

func NewHttpStoreFrom(client http.Client) Store {
	return &httpStore{client}
}

type httpStore struct {
	client http.Client
}

func (s *httpStore) GetCCToken(c context.Context, product, shopId, shopType string) (response *OAuthResponse, err error) {
	clientInfo := GetClientInfo(product)
	return s.GetProductCCToken(c, clientInfo.ClientId, clientInfo.ClientSecret, shopId, shopType)
}

func (s *httpStore) GetProductCCToken(c context.Context, cid, cse, shopId, shopType string) (response *OAuthResponse, err error) {
	return s.request(c, cid, cse, authRequest{
		grantType: "client_credentials",
		shopId:    shopId,
		shopType:  shopType,
	})
}

func (s *httpStore) request(c context.Context, clientId, clientSecret string, auth authRequest) (response *OAuthResponse, err error) {
	query := url.Values{}
	query.Add("grant_type", auth.grantType)
	query.Add("client_id", clientId)
	query.Add("client_secret", clientSecret)
	if len(auth.shopId) > 0 {
		query.Add("shop_id", auth.shopId)
	}
	if len(auth.shopType) > 0 {
		query.Add("shop_type", auth.shopType)
	}

	req := s.client.NewRequest()
	req.GetHeader().SetMethod(http.MethodPost)
	req.GetHeader().SetContentTypeBytes(http.ContentTypeJsonByte)
	req.SetRequestURI(requestUri)
	req.SetQueryString(query.Encode())
	req.SetBody(blankRequest)

	res := s.client.NewResponse()
	err = s.client.Do(c, req, res)
	if err != nil {
		return
	}
	if res.StatusCode() != 200 {
		wlog.Warnf("failed to get token, status code: %d, body: %s", res.StatusCode(), string(res.Body()))
		return nil, fmt.Errorf("failed to get token, status code: %d", res.StatusCode())
	}

	response = new(OAuthResponse)
	err = codec.Json.Unmarshal(res.Body(), response)
	if err != nil {
		return
	}
	response.ExpiresAt = time.Now().Add(time.Duration(response.ExpiresIn) * time.Second)
	return
}
