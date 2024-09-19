package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleAdd(t *testing.T) {
	assert.Equal(t, 2, 1+1, "wrong calculation")
}
