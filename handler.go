package fastgo3

import "github.com/valyala/fasthttp"

type HandlerFunc func(ctx *Context)

func Web404Handler(context *Context) {
	httpCode := fasthttp.StatusNotFound
	message := fasthttp.StatusMessage(httpCode)
	WebErrHandler(context, httpCode, message)
}

func Web405Handler(context *Context) {
	httpCode := fasthttp.StatusMethodNotAllowed
	message := fasthttp.StatusMessage(httpCode)
	WebErrHandler(context, httpCode, message)
}

func WebErrHandler(context *Context, httpCode int, message string) {
	context.fastHttpRequestCtx.Response.Reset()
	context.fastHttpRequestCtx.SetStatusCode(httpCode)
	context.fastHttpRequestCtx.SetBodyString(message)
}

