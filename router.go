package fastgo3

import (
	"fmt"
	"regexp"
	"strings"
)

type MethodHandler map[string]HandlerFunc

type StaticRoute struct {
	Method  string
	Uri     string
	Handler HandlerFunc
}

type DynRouteNode struct {
	path           string         // origin path
	name           string		  // name of path
	rePath         *regexp.Regexp // regex path
	Handler        HandlerFunc    // HandlerFunc or nil
	isEnd          bool
	staticChildren map[string]*DynRouteNode
	dynChildren    []*DynRouteNode
	method         string
}

func NewDynRouteNode(path string, isEnd bool, handler HandlerFunc) DynRouteNode {
	node := DynRouteNode{
		path:           path,
		name: 			path,
		rePath:         nil,
		isEnd:          isEnd,
		Handler:        handler,
		staticChildren: make(map[string]*DynRouteNode),
		dynChildren:    make([]*DynRouteNode, 0),
	}
	l := len(node.path)
	if l >= 2 && node.path[0] == '<' && node.path[l-1] == '>' {
		vs := strings.SplitN(path[1 : l-1], ":", 2)
		node.name = vs[0]
		tp := "s"
		if len(vs) > 1 {
			tp = vs[1]
		}
		if tp != "s" && tp != "i" && tp != "f" {
			panic(fmt.Sprintf("Can't support type: %s. Just support dynamic path type [s, i, l, f]", tp))
		}
		if tp == "s" {
			node.rePath, _ = regexp.Compile(`[^/]+`)
		} else if tp == "i" {
			node.rePath, _ = regexp.Compile(`^\d+$`)
		} else if tp == "f" {
			node.rePath, _ = regexp.Compile(`^\d+(\.\d+)??$`)
		}
	}
	return node
}

//func (node *DynRouteNode) ToString() string {
//
//}
//
//func (node *DynRouteNode) iter

func (node *DynRouteNode) findForMatch(path string) *DynRouteNode {
	if child, ok := node.staticChildren[path]; ok {
		return child
	}
	for _, child := range node.dynChildren {
		if child != nil && child.rePath.Match([]byte(path)) {
			return child
		}
	}
	return nil
}

func (node *DynRouteNode) findForAdd(path string) *DynRouteNode {
	if child, ok := node.staticChildren[path]; ok {
		return child
	}
	for _, child := range node.dynChildren {
		if child != nil && child.path == path {
			return child
		}
	}
	return nil
}

func (node *DynRouteNode) Add(child *DynRouteNode) {
	if child.rePath != nil {
		node.dynChildren = append(node.dynChildren, child)
	} else {
		node.staticChildren[child.path] = child
	}
}

type Router struct {
	StaticRoutes    map[string]MethodHandler
	DynRoot         DynRouteNode
	DynRoutesHeight int
}

func newRouter() Router {
	return Router{
		StaticRoutes:    make(map[string]MethodHandler),
		DynRoot:         NewDynRouteNode("", false, nil),
		DynRoutesHeight: 0,
	}
}

func (router *Router) isStatic(uri string) bool {
	for _, path := range strings.Split(uri, "/") {
		if len(path) > 2 && path[0] == '<' && path[len(path)-1] == '>' {
			return false
		}
	}
	return true
}

func (router *Router) Add(method string, uri string, handler HandlerFunc) *Router {
	uri = Purify(uri)
	if router.isStatic(uri) {
		r := StaticRoute{
			Method:  Upper(method),
			Uri:     uri,
			Handler: handler,
		}
		router.addAStaticRoute(r)
	} else {
		node := &router.DynRoot
		vs := strings.Split(uri, "/")
		l := len(vs) - 1
		for idx, path := range vs {
			child := node.findForAdd(path)
			if child != nil {
				node = child
				if !node.isEnd && idx == l {
					node.isEnd = true
					node.method = method
				}
			} else if idx != l {
				newNode := NewDynRouteNode(path, false, nil)
				node.Add(&newNode)
				node = &newNode
			} else {
				newNode := NewDynRouteNode(path, true, handler)
				newNode.method = method
				node.Add(&newNode)
				node = &newNode
			}
		}
	}
	return router
}

func (router *Router) addAStaticRoute(r StaticRoute) *Router {
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

func (router *Router) Match(uri string, method string) (HandlerFunc, map[string]string, int) {
	uri = Purify(uri)
	method = Upper(method)
	handler, ok := router.matchAStaticRoute(uri, method)
	if ok == 504 || ok == 0 { // 504 or 200
		return handler, nil, ok
	}

	node := &router.DynRoot
	args := make(map[string]string)
	for _, path := range strings.Split(uri, "/") {
		child := node.findForMatch(path)
		if child == nil {
			return nil, nil, 404
		}
		node = child
		args[node.name] = path
	}
	if !node.isEnd {
		return nil, nil, 404
	}
	if node.method != method {
		return nil, nil, 504
	}
	return node.Handler, args, 0
}

func (router *Router) matchAStaticRoute(uri string, method string) (HandlerFunc, int) {
	methodHandler, ok := router.StaticRoutes[uri]
	if !ok { // 404
		return nil, 404
	}
	handler, ok := methodHandler[method]
	if !ok { // 504
		return nil, 504
	}

	return handler, 0
}
