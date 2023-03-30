package http

import (
	"bytes"
	"context"
	"github.com/weimob-tech/go-project-base/pkg/wlog"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// defaultHeader 是简单 header 封装
type defaultHeader struct {
	method      string
	contentType []byte
}

func (header *defaultHeader) SetMethod(method string) {
	header.method = method
}

func (header *defaultHeader) SetContentTypeBytes(content []byte) {
	header.contentType = content
}

// defaultFileInfo 是简单文件信息封装
type defaultFileInfo struct {
	formName string
	fileName string
	file     io.Reader
}

// defaultRequest 包含了最基础的 http 请求信息
type defaultRequest struct {
	header   *defaultHeader
	uri      string
	query    string
	body     []byte
	layer    *http.Request
	fileInfo *defaultFileInfo
}

func (request *defaultRequest) GetRequest() any {
	return request.layer
}

func (request *defaultRequest) SetBody(body []byte) {
	request.body = body
}

func (request *defaultRequest) SetFile(formName, fileName string, file io.Reader) {
	request.fileInfo = &defaultFileInfo{formName, fileName, file}
}

func (request *defaultRequest) GetHeader() Header {
	return request.header
}

func (request *defaultRequest) SetRequestURI(uri string) {
	request.uri = uri
}

func (request *defaultRequest) SetQueryString(query string) {
	request.query = query
}

// defaultResponse 包括了最基础的返回信息封装
type defaultResponse struct {
	layer  *http.Response
	status int
	body   []byte
}

func (response *defaultResponse) StatusCode() int {
	return response.status
}

func (response *defaultResponse) Body() []byte {
	return response.body
}

func (response *defaultResponse) GetResponse() any {
	return response.layer
}

type defaultClient struct {
	*http.Client
	logLvl LogLvl
}

func (client *defaultClient) GetClient() any {
	return client.Client
}

func (client *defaultClient) NewRequest() Request {
	return &defaultRequest{header: &defaultHeader{}}
}

func (client *defaultClient) NewResponse() Response {
	return &defaultResponse{}
}

func (client *defaultClient) Do(ctx context.Context, request Request, response Response) (err error) {
	// prepare request
	req := request.(*defaultRequest)

	var fullUrl = req.uri
	if len(req.query) != 0 {
		fullUrl = fullUrl + "?" + req.query
	}
	parsedUrl, err := url.Parse(fullUrl)
	if err != nil {
		return err
	}

	var body io.Reader = nil
	if req.fileInfo != nil {
		body = &bytes.Buffer{}
		writer := multipart.NewWriter(body.(*bytes.Buffer))

		// create file
		part, err := writer.CreateFormFile(req.fileInfo.formName, req.fileInfo.fileName)
		if err != nil {
			return err
		}
		_, err = io.Copy(part, req.fileInfo.file)
		if err != nil {
			return err
		}

		// create filed
		err = writer.WriteField("name", req.fileInfo.fileName)
		if err != nil {
			return err
		}

		// close writer
		err = writer.Close()
		if err != nil {
			return err
		}
		req.header.contentType = []byte(writer.FormDataContentType())
	} else if len(req.body) != 0 {
		body = bytes.NewReader(req.body)
	}

	// do action
	input, err := http.NewRequestWithContext(ctx, req.header.method, parsedUrl.String(), body)
	if err != nil {
		return err
	}
	input.Header.Set("Content-Type", string(req.header.contentType))
	req.layer = input

	return elapsed(client.logLvl, func() error {
		// log before
		if client.logLvl >= LogLvlBase {
			wlog.Infof("> HTTP Request: %s", input.Method)
			wlog.Infof("> HTTP Request: %s", input.URL)
		}
		if client.logLvl >= LogLvlHeader {
			wlog.Infof("> HTTP Request header:")
			for k, v := range input.Header {
				wlog.Infof("> > %s=%s", k, strings.Join(v, ","))
			}
		}
		if client.logLvl >= LogLvlBody {
			wlog.Infof("> HTTP Request body: %s", string(req.body))
		}

		// do request
		output, err := client.Client.Do(input)
		if err != nil {
			return err
		}
		defer func() {
			e := output.Body.Close()
			if e != nil {
				wlog.Errorf("failed to close response body, err: %v", e)
			}
		}()

		// prepare response
		res := response.(*defaultResponse)
		res.status = output.StatusCode
		res.body, err = io.ReadAll(output.Body)
		res.layer = output

		// log after
		if client.logLvl >= LogLvlBase {
			wlog.Infof(">")
			wlog.Infof("> HTTP Response: %s", output.Status)
		}
		if client.logLvl >= LogLvlHeader {
			wlog.Infof("> HTTP Response header:")
			for k, v := range output.Header {
				wlog.Infof("> > %s=%s", k, strings.Join(v, ","))
			}
		}
		if client.logLvl >= LogLvlBody {
			wlog.Infof("> HTTP Response body: %s", string(res.body))
		}
		return err
	})
}

func elapsed(lvl LogLvl, fn func() error) error {
	if lvl == LogLvlWarn {
		return fn()
	}
	var start = time.Now()
	err := fn()
	wlog.Infof("> HTTP: elapsed %v", time.Since(start))
	wlog.Info(">")
	return err
}
