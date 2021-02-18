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
