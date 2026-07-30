package headers

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestHeaderParse(t *testing.T) {
	// Test: Valid single header
	headers := NewHeaders()
	data := []byte("Host: localhost:9988\r\nFooFoo:     barbar      \r\n\r\n")
	n, done, err := headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	host, ok := headers.Get("Host")
	assert.True(t, ok)
	assert.Equal(t, "localhost:9988", host)

	foofoo, ok := headers.Get("FooFoo")
	assert.True(t, ok)
	assert.Equal(t, "barbar", foofoo)

	_, ok = headers.Get("missingKey")
	assert.False(t, ok)
	assert.Equal(t, 50, n)
	assert.True(t, done)

	// Test: Invalid spacing header
	headers = NewHeaders()
	data = []byte("			Host : localhost:9988		\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)

	headers = NewHeaders()
	data = []byte("H©st: localhost:9988\r\n\r\n")
	n, done, err = headers.Parse(data)
	require.Error(t, err)
	assert.Equal(t, 0, n)
	assert.False(t, done)
	
	headers = NewHeaders()
	data = []byte("Host: localhost:9988\r\nHost: localhost:9988\r\n")
	n, done, err = headers.Parse(data)
	require.NoError(t, err)
	require.NotNil(t, headers)
	assert.False(t, done)

	host, ok = headers.Get("HOST")
	assert.True(t, ok)
	assert.Equal(t, "localhost:9988,localhost:9988", host)
	assert.False(t, done)
}
