package fastgo3

type HandlerFunc func(ctx *Context)
type Route struct {
	Method  string
	Uri     string
	Handler HandlerFunc
}
func NewRoute(method string, uri string, handler HandlerFunc) Route {
	uri = Purify(uri)
	return Route{
		Method: Upper(method),
		Uri: uri,
		Handler: handler,
	}
}

type MethodHandler map[string]HandlerFunc
type Router struct {
	StaticRoutes map[string]MethodHandler
}

func newRouter() Router {
	return Router { StaticRoutes: make(map[string]MethodHandler) }
}

func (router *Router) Add(r Route) *Router {
	methodHandler, ok := router.StaticRoutes[r.Uri]
	if !ok { // 不存在
		methodHandler = make(MethodHandler)
		router.StaticRoutes[r.Uri] = methodHandler
	}
	method := r.Method
	if _, ok := methodHandler[method]; !ok {
		methodHandler[method] = r.Handler
	}
	return router
}

func (router *Router) Match(uri string, method string) (HandlerFunc, int) {
	uri = Purify(uri)
	methodHandler, ok := router.StaticRoutes[uri]
	if !ok { // 404
		return nil, -1
	}

	method = Upper(method)
	handler, ok := methodHandler[method]
	if !ok { // 504
		return nil, -2
	}

	return handler, 0
}
