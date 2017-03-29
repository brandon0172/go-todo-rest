package main

import (
	"testing"

	"github.com/brandon0172/go-todo-rest/go-todo-rest/web"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	object := web.NewServer()

	assert := assert.New(t)

	assert.Equal(123, 123, "they should be equal")

	assert.NotEqual(123, 456, "they should not be equal")

	assert.NotNil(object)

	if assert.NotNil(object) {
		assert.Equal("Something", "Something")
	}
}
