package fastgo3

type ActionFunc func(ctx *Context)

type Route struct {
	Method string
	Uri string
	Action ActionFunc
}

func Get(uri string, action ActionFunc) Route {
	return Route {Method: "GET", Uri: uri, Action: action}
}

func Post(uri string, action ActionFunc) Route {
	return Route {Method: "POST", Uri: uri, Action: action}
}

func Put(uri string, action ActionFunc) Route {
	return Route {Method: "PUT", Uri: uri, Action: action}
}

func Delete(uri string, action ActionFunc) Route {
	return Route {Method: "DELETE", Uri: uri, Action: action}
}

func Patch(uri string, action ActionFunc) Route {
	return Route {Method: "PATCH", Uri: uri, Action: action}
}
