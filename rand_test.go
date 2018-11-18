package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandString(t *testing.T) {
	s1 := RandString(100, "")
	assert.NotEmpty(t, s1)
	assert.Equal(t, 100, len(s1))

	s2 := RandString(100, "")
	assert.NotEmpty(t, s2)
	assert.Equal(t, 100, len(s2))
	assert.NotEqual(t, s1, s2)

	s3 := RandString(5, "a")
	assert.NotEmpty(t, s3)
	assert.Equal(t, "aaaaa", s3)
}
