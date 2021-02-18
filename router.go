package fastgo3

import (
	"strings"
)

type HandlerFunc func(ctx *Context)
type Route struct {
	Method  string
	Uri     string
	Handler HandlerFunc
}


type MethodHandler map[string]HandlerFunc
type Router struct {
	StaticRoutes map[string]MethodHandler
}

func newRouter() Router {
	return Router { StaticRoutes: make(map[string]MethodHandler) }
}

func (router *Router) Add(route Route) *Router {
	methodHandler, ok := router.StaticRoutes[route.Uri]
	if !ok { // 不存在
		methodHandler = make(MethodHandler)
		router.StaticRoutes[route.Uri] = methodHandler
	}
	method := strings.ToLower(route.Method)
	if _, ok := methodHandler[method]; !ok {
		methodHandler[method] = route.Handler
	}
	return router
}

func (router *Router) Match(uri string, method string) (HandlerFunc, int) {
	methodHandler, ok := router.StaticRoutes[uri]
	if !ok { // 404
		return nil, -1
	}
	method = strings.ToLower(method)
	handler, ok := methodHandler[method]
	if !ok { // 504
		return nil, -2
	}

	return handler, 0
}
