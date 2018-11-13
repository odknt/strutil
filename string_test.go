package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDef(t *testing.T) {
	assert.Equal(t, "str", Def("str", "default"))
	assert.Equal(t, "default", Def("", "default"))
}
