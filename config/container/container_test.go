package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainer(t *testing.T) {
	c := new(Container)
	assert.Nil(t, c.httpadapter)
	assert.Nil(t, c.service)
	assert.Nil(t, c.cache)
	assert.Nil(t, c.server)
	assert.NotNil(t, c.GetHttpAdapter())
	assert.NotNil(t, c.GetService())
	assert.NotNil(t, c.GetCache())
	assert.NotNil(t, c.GetServer(nil, nil))
	assert.NotNil(t, c.httpadapter)
	assert.NotNil(t, c.service)
	assert.NotNil(t, c.cache)
	assert.NotNil(t, c.server)
}