package fastgo3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func fakeAction(context *Context) {}

func TestStaticMatch(t *testing.T) {
	assert := assert.New(t)

	router := newRouter()
	router.Add(Get("/hello", fakeAction))
	router.Add(Post("/upload", fakeAction))

	action, _ := router.Match("/hello", "GET")
	assert.NotNil(action)

	action, _ = router.Match("/hello", "get")
	assert.NotNil(action)

	action, _ = router.Match("/upload", "Post")
	assert.NotNil(action)

	action, _ = router.Match("/upload", "PoSt")
	assert.NotNil(action)
}

func TestStaticMiss(t *testing.T) {
	assert := assert.New(t)

	router := newRouter()
	router.Add(Get("/hello", fakeAction))

	_, errno := router.Match("/hi", "GET")
	assert.Equal(-1, errno)

	_, errno = router.Match("/hello", "Post")
	assert.Equal(-2, errno)
}
