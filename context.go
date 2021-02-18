package fastgo3

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type Context struct {
	fastHttpRequestCtx *fasthttp.RequestCtx
}

func(context *Context) SetHeader(name string, value string) *Context {
	context.fastHttpRequestCtx.Response.Header.Set(name, value)
	return context
}

func(context *Context) SetCookie(name string, value string) *Context {
	var c fasthttp.Cookie
	c.SetKey(name)
	c.SetValue(value)
	context.fastHttpRequestCtx.Response.Header.SetCookie(&c)
	return context
}

func(context *Context) SetContentType(contentType string) *Context {
	context.fastHttpRequestCtx.SetContentType(contentType)
	return context
}

func(context *Context) Method() string {
	return string(context.fastHttpRequestCtx.Method())
}

func(context *Context) RequestURI() string {
	return string(context.fastHttpRequestCtx.RequestURI())
}

func(context *Context) Path() string {
	return string(context.fastHttpRequestCtx.Path())
}

func(context *Context) Host() string {
	return string(context.fastHttpRequestCtx.Host())
}

func(context *Context) QueryArgs() *fasthttp.Args {
	return context.fastHttpRequestCtx.QueryArgs()
}

func(context *Context) UserAgent() string {
	return string(context.fastHttpRequestCtx.UserAgent())
}

func(context *Context) ConnTime() string {
	return context.fastHttpRequestCtx.ConnTime().Format("2006-01-01T00:00:000000")
}

func(context *Context) Time() string {
	return context.fastHttpRequestCtx.Time().Format("2006-01-01T00:00:000000")
}

func(context *Context) ConnRequestNum() uint64 {
	return context.fastHttpRequestCtx.ConnRequestNum()
}

func(context *Context) RemoteIP() string {
	return fmt.Sprintf("%q", context.fastHttpRequestCtx.RemoteIP())
}

func(context *Context) Request() *fasthttp.Request {
	return &context.fastHttpRequestCtx.Request
}

// Write writes p into response body.
func (context *Context) Write(p []byte) (int, error) {
	context.fastHttpRequestCtx.Response.AppendBody(p)
	return len(p), nil
}
