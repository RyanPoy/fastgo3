package fastgo3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrifyUri(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("/abc", Purify("/abc"))
	assert.Equal("/abc", Purify("/abc/"))
	assert.Equal("/", Purify("/"))
}

func TestUpperMethod(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("GET", Upper("get"))
	assert.Equal("GET", Upper("Get"))
}

func TestUpprMethodShouldGiveErrorIfMethodUnsupported(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() { Upper("UnsupportedMethod") })
}
