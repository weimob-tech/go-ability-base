package http

var (
	ContentTypeHtmlByte     = []byte("text/html")
	ContentTypeJsonByte     = []byte("application/json")
	ContentTypeFromDataByte = []byte("multipart/form-data")
)

const (
	MethodGet  = "GET"
	MethodPut  = "PUT"
	MethodDel  = "DELETE"
	MethodHead = "HEAD"
	MethodPost = "POST"
)
