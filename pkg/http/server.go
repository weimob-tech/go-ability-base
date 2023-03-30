package http

import "context"

type DelegateServer interface {
	GetServer() any
}

type Server interface {
	DelegateServer
	Start()
	AddExtendCallback(*ExtendCallbackConfig)
}

type ExtendCallback func(ctx *ExtendCallbackContext) (any, error)

type ExtendCallbackConfig struct {
	Method      string
	Path        string
	Params      []string
	Headers     []string
	Queries     []string
	Callback    ExtendCallback
	NullToEmpty bool
}

type ExtendCallbackContext struct {
	Context context.Context
	Path    string
	Method  string
	Params  map[string]string
	Headers map[string]string
	Queries map[string]string
	Body    []byte
}
