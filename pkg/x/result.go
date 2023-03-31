package x

import (
	"fmt"
)

type Code struct {
	Errcode string `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (c Code) GetErrcode() string {
	return c.Errcode
}

func (c Code) GetErrmsg() string {
	return c.Errmsg
}

type Result interface {
	GetCode() Code
	IsSuccess() bool
}

type ErrCode interface {
	GetErrcode() string
	GetErrmsg() string
}

// BizError 是通用业务异常
type BizError struct {
	Errcode string `json:"errcode,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
}

func Error(code, msg string) error {
	return &BizError{code, msg}
}

func ErrorOf(code Code) error {
	return &BizError{code.GetErrcode(), code.GetErrmsg()}
}

func (r BizError) GetErrcode() string {
	return r.Errcode
}

func (r BizError) GetErrmsg() string {
	return r.Errmsg
}

func (r BizError) Error() string {
	return fmt.Sprintf("BizError{errcode: %s, errmsg: %s}", r.Errcode, r.Errmsg)
}

// result 是返回结果的公共数据结构
type result struct {
	Code Code `json:"code,omitempty"`
}

func (r result) GetCode() Code {
	return r.Code
}

func (r result) IsSuccess() bool {
	return r.Code.GetErrcode() == "0"
}

// EmptyResult 只保留 code，没有任何内容
type EmptyResult struct {
	result
}

// AnyResult 可以包含任意结果
type AnyResult[T any] struct {
	result
	Data T `json:"data,omitempty"`
}

var success = &EmptyResult{result{Code{Errmsg: "success"}}}

func Ok() Result {
	return success
}

func OkAnd[T any](data T) Result {
	return &AnyResult[T]{
		result: result{Code{Errcode: "0", Errmsg: "success"}},
		Data:   data,
	}
}

func Fail(code string, msg string) Result {
	return &EmptyResult{result{
		Code{Errcode: code, Errmsg: msg},
	}}
}

func FailAnd(code ErrCode) Result {
	return &EmptyResult{result{
		Code{Errcode: code.GetErrmsg(), Errmsg: code.GetErrmsg()},
	}}
}
