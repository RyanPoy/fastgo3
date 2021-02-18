package fastgo3

import (
	"fmt"
	"strings"
)

func Purify(uri string) string {
	l := len(uri)
	if l > 1 && uri[l-1] == '/' {
		return string(uri[:l-1])
	}
	return uri
}


var SUPPORT_METHODS = []string{"GET", "HEAD", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"}
var SUPPORT_METHOD_SET = func() map[string]int {
	m := make(map[string]int)
	for _, key := range SUPPORT_METHODS {
		m[key] = 0
	}
	return m
}()

func Upper(method string) string {
	r := strings.ToUpper(method)
	_, ok := SUPPORT_METHOD_SET[r]
	if ok {
		return r
	}
	panic(fmt.Sprintf("Just support %s, but got '%s'", SUPPORT_METHODS, method))
}
