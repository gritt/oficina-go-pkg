package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	// given
	wantHello := "Hello World"
	wantLen := 11

	// when
	gotHello, gotLen := HelloWorld()

	// then
	assert.Equal(t, wantHello, gotHello)
	assert.Equal(t, wantLen, gotLen)
}
