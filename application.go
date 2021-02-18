package fastgo3

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"runtime/debug"
)

type Application struct {
	ip     string
	port   int
	router *Router
}

func NewApplication() Application {
	router := newRouter()
	return Application{router: &router}
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

func (app *Application) writeError(ctx *fasthttp.RequestCtx) {
	if err := recover(); err != nil {
		httpCode := fasthttp.StatusInternalServerError
		ctx.Response.Reset()
		ctx.SetStatusCode(httpCode)
		errMsg := fmt.Sprintf("%s\n\n%s", fasthttp.StatusMessage(httpCode), string(debug.Stack()))
		ctx.SetBodyString(errMsg)
	}
}

func (app *Application) dispatch(ctx *fasthttp.RequestCtx) {
	defer app.writeError(ctx)

	uri, method := string(ctx.Path()), string(ctx.Method())
	handler, errno := app.router.Match(uri, method)
	if errno == 0 {
		context := NewContext(ctx)
		handler(&context)
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

func (app *Application) Get(uri string, handler HandlerFunc) *Application {
	return app.Route([]string{"GET"}, uri, handler)
}

func (app *Application) Post(uri string, handler HandlerFunc) *Application {
	return app.Route([]string{"POST"}, uri, handler)
}

func (app *Application) Put(uri string, handler HandlerFunc) *Application {
	return app.Route([]string{"PUT"}, uri, handler)
}

func (app *Application) Delete(uri string, handler HandlerFunc) *Application {
	return app.Route([]string{"DELETE"}, uri, handler)
}

func (app *Application) Patch(uri string, handler HandlerFunc) *Application {
	return app.Route([]string{"PATCH"}, uri, handler)
}

func (app *Application) Connect(uri string, handler HandlerFunc) *Application {
	return app.Route([]string{"CONNECT"}, uri, handler)
}

func (app *Application) Head(uri string, handler HandlerFunc) *Application {
	return app.Route([]string{"HEAD"}, uri, handler)
}

func (app *Application) Options(uri string, handler HandlerFunc) *Application {
	return app.Route([]string{"OPTIONS"}, uri, handler)
}

func (app *Application) Trace(uri string, handler HandlerFunc) *Application {
	return app.Route([]string{"TRACE"}, uri, handler)
}

func (app *Application) Route(methods []string, uri string, handler HandlerFunc) *Application {
	for _, method := range methods {
		r := NewRoute(method, uri, handler)
		app.router.Add(r)
	}
	return app
}

func (app *Application) GetRouter() *Router {
	return app.router
}
