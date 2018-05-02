package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tts, err := Handler(Request{
		location: "CBUS",
	})
	assert.IsType(t, nil, err)
	assert.NotEqual(t, 0, tts)
}
