package fastgo3

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type ApiResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Context struct {
	fastHttpRequestCtx *fasthttp.RequestCtx
	Method             string
	argFuncs           []func() *fasthttp.Args
	middlewares        *[]HandlerFunc
	middlewareIdx      int
	handler            *HandlerFunc
	SeqId              string
	UriArgs			   map[string]string
}

func NewContext(fastHttpRequestCtx *fasthttp.RequestCtx) Context {
	c := Context{
		fastHttpRequestCtx: fastHttpRequestCtx,
		Method:             Upper(string(fastHttpRequestCtx.Method())),
		middlewareIdx:      0,
		UriArgs: 			make(map[string]string),
	}
	if c.Method == "GET" {
		c.argFuncs = []func() *fasthttp.Args{c.QueryArgs, c.PostArgs}
	} else {
		c.argFuncs = []func() *fasthttp.Args{c.PostArgs, c.QueryArgs}
	}
	return c
}

func (context *Context) Next() {
	l := len(*context.middlewares)
	if context.middlewareIdx < l { // middleware
		middleware := (*context.middlewares)[context.middlewareIdx]
		context.middlewareIdx += 1
		middleware(context)
	} else if context.middlewareIdx == l { // handler
		(*context.handler)(context)
	}
}

func (context *Context) SetHeader(name string, value string) *Context {
	context.fastHttpRequestCtx.Response.Header.Set(name, value)
	return context
}

func (context *Context) SetCookie(name string, value string) *Context {
	var c fasthttp.Cookie
	c.SetKey(name)
	c.SetValue(value)
	context.fastHttpRequestCtx.Response.Header.SetCookie(&c)
	return context
}

func (context *Context) SetContentType(contentType string) *Context {
	context.fastHttpRequestCtx.SetContentType(contentType)
	return context
}

func (context *Context) RequestURI() string {
	return string(context.fastHttpRequestCtx.RequestURI())
}

func (context *Context) Path() string {
	return string(context.fastHttpRequestCtx.Path())
}

func (context *Context) Host() string {
	return string(context.fastHttpRequestCtx.Host())
}

func (context *Context) UserAgent() string {
	return string(context.fastHttpRequestCtx.UserAgent())
}

func (context *Context) ConnTime() string {
	return context.fastHttpRequestCtx.ConnTime().Format("2006-01-01T00:00:000000")
}

func (context *Context) Time() string {
	return context.fastHttpRequestCtx.Time().Format("2006-01-01T00:00:000000")
}

func (context *Context) ConnRequestNum() uint64 {
	return context.fastHttpRequestCtx.ConnRequestNum()
}

func (context *Context) RemoteIP() string {
	return fmt.Sprintf("%q", context.fastHttpRequestCtx.RemoteIP())
}

func (context *Context) Request() *fasthttp.Request {
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

func (context *Context) Ok(data interface{}) {
	context.Finish(0, "", data)
}

func (context *Context) Err(msg string) {
	context.Finish(-1, msg, make(map[string]string))
}

func (context *Context) Finish(code int, msg string, data interface{}) {
	apiResult := ApiResult{Code: code, Msg: msg, Data: data}
	context.RenderJson(apiResult)
}

func (context *Context) Redirect(url string) {
	context.fastHttpRequestCtx.Redirect(url, fasthttp.StatusFound)
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

func (context *Context) FloatParam(key string, defaultValue float64) float64 {
	for _, argFunc := range context.argFuncs {
		if value, err := argFunc().GetUfloat(key); err != nil {
			return value
		}
	}
	return defaultValue
}

func (context *Context) IntParam(key string, defaultValue int) int {
	for _, argFunc := range context.argFuncs {
		if value, err := argFunc().GetUint(key); err != nil {
			return value
		}
	}
	return defaultValue
}
