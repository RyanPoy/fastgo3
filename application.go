package fastgo3

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

type Application struct {
	ip string
	port int
	router *Router
}

func NewApplication() Application {
	router := newRouter()
	return Application{ router: &router }
}

func (app *Application) Run(ip string, port int) {
	app.ip = ip
	app.port = port

	addr := fmt.Sprintf("%s:%d", app.ip, app.port)

	dispatch := fasthttp.CompressHandler(app.dispatch)
	if err := fasthttp.ListenAndServe(addr, dispatch); err != nil {
		log.Fatalf("Error: %s", err)
	} else {
		log.Printf("Run on (%s) at (%d)", app.ip, app.port)
	}
}

func (app *Application) dispatch(ctx *fasthttp.RequestCtx) {
	uri, method := string(ctx.Path()), string(ctx.Method())
	action, errno := app.router.Match(uri, method)
	if errno == 0  {
		context := Context { fastHttpRequestCtx: ctx }
		context.SetContentType("text/plain; charset=utf8")
		action(&context)
		log.Printf("[%s] \n", uri)
		return
	}
	var httpCode int
	if errno == -1 {
		httpCode = fasthttp.StatusNotFound
	} else {
		httpCode = fasthttp.StatusMethodNotAllowed
	}
	ctx.Response.Reset()
	ctx.SetStatusCode(httpCode)
	ctx.SetBodyString(fasthttp.StatusMessage(httpCode))
}

func (app *Application) Amount(routes []Route) *Application {
	for _, r := range routes {
		app.router.Add(r)
	}
	return app
}