package fastgo3

import "strings"

func Purify(uri string) string {
	l := len(uri)
	if uri[l-1] == '/' {
		return string(uri[:l-1])
	}
	return uri
}

func Upper(method string) string {
	return strings.ToUpper(method)
}
