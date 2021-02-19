package fastgo3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStaticMatchGet(t *testing.T) {
	assert := assert.New(t)

	app := Default()
	app.Get("/hello", nil)

	router := app.GetRouter()
	_, errno := router.Match("/hello", "GET")
	assert.Equal(0, errno)

	_, errno = router.Match("/hello", "get")
	assert.Equal(0, errno)

	_, errno = router.Match("/hello/", "GET")
	assert.Equal(0, errno)
}

func TestStaticMatchGetChinese(t *testing.T) {
	assert := assert.New(t)

	app := Default()
	app.Get("/你好", nil)

	router := app.GetRouter()
	_, errno := router.Match("/你好", "GET")
	assert.Equal(0, errno)
}

func TestStaticMatchPost(t *testing.T) {
	assert := assert.New(t)

	app := Default()
	app.Post("/hello", nil)

	router := app.GetRouter()

	_, errno := router.Match("/hello", "Post")
	assert.Equal(0, errno)

	_, errno = router.Match("/hello", "PoSt")
	assert.Equal(0, errno)
}

func TestStaticMatchGetAndPost(t *testing.T) {
	assert := assert.New(t)

	app := Default()
	app.Get("/hello", nil)
	app.Post("/hello", nil)

	router := app.GetRouter()
	_, errno := router.Match("/hello", "GET")
	assert.Equal(0, errno)

	_, errno = router.Match("/hello", "PoSt")
	assert.Equal(0, errno)
}

func TestStaticMatchOtherHttpMethods(t *testing.T) {
	assert := assert.New(t)

	app := Default()
	app.Route([]string{"PUT", "Delete"}, "/upload", nil)

	router := app.GetRouter()
	_, errno := router.Match("/upload", "put")
	assert.Equal(0, errno)

	_, errno = router.Match("/upload", "delete")
	assert.Equal(0, errno)
}

func TestStaticMiss(t *testing.T) {
	assert := assert.New(t)

	app := Default()
	app.Get("/hello", nil)
	app.Post("/upload", nil)

	router := app.GetRouter()

	_, errno := router.Match("/hi", "GET")
	assert.Equal(-1, errno)

	_, errno = router.Match("/hello", "Post")
	assert.Equal(-2, errno)
}
