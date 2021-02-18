package fastgo3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func fakeAction(context *Context) {}

func TestStaticMatchGet(t *testing.T) {
	assert := assert.New(t)

	app := NewApplication()
	app.Get("/hello", fakeAction)

	router := app.GetRouter()
	handler, _ := router.Match("/hello", "GET")
	assert.NotNil(handler)

	handler, _ = router.Match("/hello", "get")
	assert.NotNil(handler)

	handler, _ = router.Match("/hello/", "GET")
	assert.NotNil(handler)
}

func TestStaticMatchPost(t *testing.T) {
	assert := assert.New(t)

	app := NewApplication()
	app.Post("/hello", fakeAction)

	router := app.GetRouter()

	handler, _ := router.Match("/hello", "Post")
	assert.NotNil(handler)

	handler, _ = router.Match("/hello", "PoSt")
	assert.NotNil(handler)
}

func TestStaticMatchGetAndPost(t *testing.T) {
	assert := assert.New(t)

	app := NewApplication()
	app.Get("/hello", fakeAction)
	app.Post("/hello", fakeAction)

	router := app.GetRouter()
	handler, _ := router.Match("/hello", "GET")
	assert.NotNil(handler)

	handler, _ = router.Match("/hello", "PoSt")
	assert.NotNil(handler)
}

func TestStaticMatchOtherHttpMethods(t *testing.T) {
	assert := assert.New(t)

	app := NewApplication()
	app.Route([]string{"PUT", "Delete"}, "/upload", fakeAction)

	router := app.GetRouter()
	handler, _ := router.Match("/upload", "put")
	assert.NotNil(handler)

	handler, _ = router.Match("/upload", "delete")
	assert.NotNil(handler)
}

func TestStaticMiss(t *testing.T) {
	assert := assert.New(t)

	app := NewApplication()
	app.Get("/hello", fakeAction)
	app.Post("/upload", fakeAction)

	router := app.GetRouter()

	_, errno := router.Match("/hi", "GET")
	assert.Equal(-1, errno)

	_, errno = router.Match("/hello", "Post")
	assert.Equal(-2, errno)
}
