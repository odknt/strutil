package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStringer struct {
	val string
}

func (t testStringer) String() string {
	return t.val
}

func TestDef(t *testing.T) {
	assert.Equal(t, "str", Def("str", "default"))
	assert.Equal(t, "default", Def("", "default"))
}

func TestDefStringer(t *testing.T) {
	assert.Equal(t, "default", DefStringer(nil, "default"))

	test := &testStringer{}
	assert.Equal(t, "default", DefStringer(test, "default"))

	test = &testStringer{val: "test"}
	assert.Equal(t, "test", DefStringer(test, "default"))
}
