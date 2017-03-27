package main

import (
	"testing"

	"github.com/badmuts/go-todo-rest/web"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	object := web.NewServer()

	assert := assert.New(t)

	assert.Equal(123, 123, "they should be equal")

	assert.NotEqual(123, 456, "they should not be equal")

	assert.Nil(object)

	if assert.NotNil(object) {
		assert.Equal("Something", object.Value)
	}
}
