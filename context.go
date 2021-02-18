package fastgo3

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type Context struct {
	fastHttpRequestCtx *fasthttp.RequestCtx
	Method             string
	argFuncs           []func() *fasthttp.Args
}

func NewContext(fastHttpRequestCtx *fasthttp.RequestCtx) Context {
	c := Context {
		fastHttpRequestCtx: fastHttpRequestCtx,
		Method:             Upper(string(fastHttpRequestCtx.Method())),
		argFuncs:           make([]func() *fasthttp.Args, 3),
	}
	if c.Method == "GET" {
		c.argFuncs[0] = c.QueryArgs
		c.argFuncs[1] = c.PostArgs
	} else {
		c.argFuncs[0] = c.PostArgs
		c.argFuncs[1] = c.QueryArgs
	}
	return c
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

func(context *Context) RequestURI() string {
	return string(context.fastHttpRequestCtx.RequestURI())
}

func(context *Context) Path() string {
	return string(context.fastHttpRequestCtx.Path())
}

func(context *Context) Host() string {
	return string(context.fastHttpRequestCtx.Host())
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
	return context.fastHttpRequestCtx.Write(p)
}

// WriteString appends s to response body.
func (context *Context) WriteString(s string) (int, error) {
	return context.fastHttpRequestCtx.WriteString(s)
}

func (context *Context) RenderString(data string) {
	context.SetContentType("text/plain; charset=utf8")
	fmt.Fprintf(context, data)
}

func (context *Context) RenderHtml(data string) {
	context.SetContentType("text/html; charset=utf8")
	context.WriteString(data)
}

func (context *Context) RenderJson(data interface{}) {
	context.SetContentType("application/json; charset=utf8")
	json.NewEncoder(context).Encode(data)
}

// QueryArgs returns query arguments from RequestURI.
//
// It doesn't return POST'ed arguments - use PostArgs() for this.
//
// Returned arguments are valid until returning from RequestHandler.
//
// See also PostArgs, FormValue and FormFile.
func (context *Context) QueryArgs() *fasthttp.Args {
	return context.fastHttpRequestCtx.QueryArgs()
}

// PostArgs returns POST arguments.
//
// It doesn't return query arguments from RequestURI - use QueryArgs for this.
//
// Returned arguments are valid until returning from RequestHandler.
//
// See also QueryArgs, FormValue and FormFile.
func (context *Context) PostArgs() *fasthttp.Args {
	return context.fastHttpRequestCtx.PostArgs()
}

func (context *Context) BytesParam(name string, defaultValue []byte) []byte {
	for _, argFunc := range context.argFuncs {
		value := argFunc().Peek(name)
		if value != nil {
			return value
		}
	}
	return defaultValue
}

func (context *Context) StrParam(name string, defaultValue string) string {
	value := context.BytesParam(name, nil)
	if value != nil {
		return string(value)
	}
	return defaultValue
}

// GetBool returns boolean value for the given key.
//
// true is returned for "1", "t", "T", "true", "TRUE", "True", "y", "yes", "Y", "YES", "Yes",
// otherwise false is returned.
func (context *Context) BoolParam(name string, defaultValue bool) bool {
	for _, argFunc := range context.argFuncs {
		args := argFunc()
		if args.Has(name) {
			return args.GetBool(name)
		}
	}
	return defaultValue
}


// GetUfloat returns ufloat value for the given key.
func (context *Context) FloatParam(key string, defaultValue float64) float64 {
	for _, argFunc := range context.argFuncs {
		if value, err := argFunc().GetUfloat(key); err != nil {
			return value
		}
	}
	return defaultValue
}

// GetUint returns uint value for the given key.
func (context *Context) IntParam(key string, defaultValue int) int {
	for _, argFunc := range context.argFuncs {
		if value, err := argFunc().GetUint(key); err != nil {
			return value
		}
	}
	return defaultValue
}
