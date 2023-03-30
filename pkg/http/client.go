package http

import (
	"context"
	"crypto/tls"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"github.com/weimob-tech/go-project-base/pkg/hook"
	"io"
	"net/http"
	"strings"
)

type LogLvl int

// LogLvlWarn
const (
	LogLvlWarn LogLvl = iota
	LogLvlBase
	LogLvlHeader
	LogLvlBody
)

func GetLevel(lvl string) LogLvl {
	switch strings.ToLower(lvl) {
	case "warn":
		return LogLvlWarn
	case "base":
		return LogLvlBase
	case "header":
		return LogLvlHeader
	case "body":
		return LogLvlBody
	default:
		return LogLvlWarn
	}
}

var (
	Global        Client
	NewHttpClient HttpClientFactory
)

type HttpClientFactory func() Client

func init() {
	config.SetDefault("client.log.lvl", "warn")
	if NewHttpClient == nil {
		NewHttpClient = func() Client {
			noTls := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			return &defaultClient{
				Client: &http.Client{Transport: noTls},
				logLvl: GetLevel(config.GetString("client.log.lvl")),
			}
		}
	}
	hook.AddPostStartHook(func() {
		if Global == nil {
			Global = NewHttpClient()
		}
	})
}

type (
	DelegateRequest interface {
		GetRequest() any
	}

	DelegateResponse interface {
		GetResponse() any
	}

	Request interface {
		DelegateRequest
		SetBody(payload []byte)
		SetFile(formName, fileName string, file io.Reader)
		GetHeader() Header
		SetRequestURI(string)
		SetQueryString(string)
	}

	Header interface {
		SetMethod(string)
		SetContentTypeBytes([]byte)
	}

	Response interface {
		DelegateResponse
		StatusCode() int
		Body() []byte
	}

	Client interface {
		GetClient() any
		NewRequest() Request
		NewResponse() Response
		Do(ctx context.Context, request Request, response Response) error
	}
)
