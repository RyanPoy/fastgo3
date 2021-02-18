package fastgo3

import (
	"strings"
)

type MethodAction map[string]ActionFunc

type Router struct {
	StaticRoutes map[string]MethodAction
}

func newRouter() Router {
	return Router { StaticRoutes: make(map[string]MethodAction) }
}

func (router *Router) Add(route Route) *Router {
	methodAction, ok := router.StaticRoutes[route.Uri]
	if !ok { // 不存在
		methodAction = make(MethodAction)
		router.StaticRoutes[route.Uri] = methodAction
	}
	method := strings.ToLower(route.Method)
	if _, ok := methodAction[method]; !ok {
		methodAction[method] = route.Action
	}
	return router
}

func (router *Router) Match(uri string, method string) (ActionFunc, int) {
	methodAction, ok := router.StaticRoutes[uri]
	if !ok {
		return nil, -1
	}
	method = strings.ToLower(method)
	action, ok := methodAction[method]
	if !ok {
		return nil, -2
	}

	return action, 0
}
