package gee

import "net/http"

type H map[string]interface {
}

// 上下文context
type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:     w,
		Req:        req,
		Path:       req.URL.Path,
		Method:     req.Method,
		StatusCode: 0,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}
