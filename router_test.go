package fastgo3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestStaticMatchGet(t *testing.T) {
	assert := assert.New(t)

	app := New()
	app.Get("/hello", nil)

	router := app.GetRouter()
	_, args, errno := router.Match("/hello", "GET")
	assert.Equal(0, errno)
	assert.Nil(args)

	_, _, errno = router.Match("/hello", "get")
	assert.Equal(0, errno)

	_, _, errno = router.Match("/hello/", "GET")
	assert.Equal(0, errno)
}

func TestStaticMatchGetChinese(t *testing.T) {
	assert := assert.New(t)

	app := New()
	app.Get("/你好", nil)

	router := app.GetRouter()
	_, _, errno := router.Match("/你好", "GET")
	assert.Equal(0, errno)
}

func TestStaticMatchPost(t *testing.T) {
	assert := assert.New(t)

	app := New()
	app.Post("/hello", nil)

	router := app.GetRouter()

	_, _, errno := router.Match("/hello", "Post")
	assert.Equal(0, errno)

	_, _, errno = router.Match("/hello", "PoSt")
	assert.Equal(0, errno)
}

func TestStaticMatchGetAndPost(t *testing.T) {
	assert := assert.New(t)

	app := New()
	app.Get("/hello", nil)
	app.Post("/hello", nil)

	router := app.GetRouter()
	_, _, errno := router.Match("/hello", "GET")
	assert.Equal(0, errno)

	_, _, errno = router.Match("/hello", "PoSt")
	assert.Equal(0, errno)
}

func TestStaticMatchOtherHttpMethods(t *testing.T) {
	assert := assert.New(t)

	app := New()
	app.Route([]string{"PUT", "Delete"}, "/upload", nil)

	router := app.GetRouter()
	_, _, errno := router.Match("/upload", "put")
	assert.Equal(0, errno)

	_, _, errno = router.Match("/upload", "delete")
	assert.Equal(0, errno)
}

func TestStaticMiss(t *testing.T) {
	assert := assert.New(t)

	app := New()
	app.Get("/hello", nil)
	app.Post("/upload", nil)

	router := app.GetRouter()

	_, _, errno := router.Match("/hi", "GET")
	assert.Equal(404, errno)

	_, _, errno = router.Match("/hello", "Post")
	assert.Equal(504, errno)
}

func TestDynMathString(t *testing.T) {
	assert := assert.New(t)
	app := New()
	app.Get("/hello/<name>", nil)

	router := app.GetRouter()
	_, args, errno := router.Match("/hello/abc", "GET")
	assert.Equal(0, errno)

	v, _ := args["name"]
	assert.Equal("abc", v)
}

func TestDynMathString2(t *testing.T) {
	assert := assert.New(t)
	app := New()
	app.Get("/hello/<name:s>", nil)

	router := app.GetRouter()
	_, args, errno := router.Match("/hello/abc", "GET")
	assert.Equal(0, errno)

	v, _ := args["name"]
	assert.Equal("abc", v)
}

func TestDynMathInt(t *testing.T) {
	assert := assert.New(t)
	app := New()
	app.Get("/users/<id:i>/hello", nil)

	router := app.GetRouter()
	_, args, errno := router.Match("/users/10/hello", "GET")
	assert.Equal(0, errno)

	v, _ := args["id"]
	assert.Equal("10", v)
}

func TestDynMissingInt(t *testing.T) {
	assert := assert.New(t)
	app := New()
	app.Get("/users/<id:i>/hello", nil)

	router := app.GetRouter()
	_, _, errno := router.Match("/users/10", "GET")
	assert.Equal(404, errno)

	_, _, errno = router.Match("/users/ryan10/hello", "GET")
	assert.Equal(404, errno)

	_, _, errno = router.Match("/users/10.1/hello", "GET")
	assert.Equal(404, errno)
}

func TestDynMathFloat(t *testing.T) {
	assert := assert.New(t)
	app := New()
	app.Get("/users/<score:f>/sort", nil)

	router := app.GetRouter()
	_, args, errno := router.Match("/users/10.1/sort", "GET")
	assert.Equal(0, errno)
	v, _ := args["score"]
	assert.Equal("10.1", v)

	_, args, errno = router.Match("/users/10.1.1/sort", "GET")
	assert.Equal(404, errno)

	_, args, errno = router.Match("/users/10/sort", "GET")
	assert.Equal(0, errno)

	v, _ = args["score"]
	assert.Equal("10", v)
}

func TestComplexMatch(t *testing.T) {
	assert := assert.New(t)
	app := New()
	app.Get("/users/<id:i>/score/<name>/<value:f>/end", nil)
	app.Get("/users/<id:i>/score", nil)
	app.Get("/users/<name:s>/hello", nil)
	app.Get("/users/123/hello", nil)

	router := app.GetRouter()
	_, _, errno := router.Match("/users/1/score/math/1.2/end", "GET")
	assert.Equal(0, errno)

	_, _, errno = router.Match("/users/1/score", "GET")
	assert.Equal(0, errno)

	_, _, errno = router.Match("/users/Jim/hello", "GET")
	assert.Equal(0, errno)

	_, _, errno = router.Match("/users/123/hello", "GET")
	assert.Equal(0, errno)
}

func TestComplexMissing(t *testing.T) {
	assert := assert.New(t)
	app := New()
	app.Get("/users/<id:i>/score/<name>/<value:f>/end", nil)
	router := app.GetRouter()

	_, _, errno := router.Match("/users/1/score/math/1.2/end", "POST")
	assert.Equal(504, errno)

	_, _, errno = router.Match("/users/1/score/math/1.2.0/end", "GET")
	assert.Equal(404, errno)

	_, _, errno = router.Match("/users/1.2/score/math/1.2/end", "GET")
	assert.Equal(404, errno)
}
