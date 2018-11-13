package strutil

import (
	"net"
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

func TestDefExpr(t *testing.T) {
	assert.Equal(t, "str", DefExpr(1 == 1, "str", "default"))
	assert.Equal(t, "default", DefExpr(1 == 2, "str", "default"))
}

func TestDefStringer(t *testing.T) {
	assert.Equal(t, "default", DefStringer(nil, "default"))

	test := &testStringer{}
	assert.Equal(t, "default", DefStringer(test, "default"))

	test = &testStringer{val: "test"}
	assert.Equal(t, "test", DefStringer(test, "default"))
}

func TestDefExprStringer(t *testing.T) {
	// net.ParseIP returns zero-length net.IP instead of nil when parse failed.
	ip := net.ParseIP("invalid")
	assert.Equal(t, "127.0.0.1", DefExprStringer(len(ip) > 0, ip, "127.0.0.1"))
	ip = net.ParseIP("0.0.0.0")
	assert.Equal(t, "0.0.0.0", DefExprStringer(len(ip) > 0, ip, "127.0.0.1"))
}
