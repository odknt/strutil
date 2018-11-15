package strutil

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestFromBytes(t *testing.T) {
	bs := []byte("test string")
	str := FromBytes(bs)

	// check no-copy conversion.
	bsptr := *(*reflect.SliceHeader)(unsafe.Pointer(&bs))
	strptr := *(*reflect.SliceHeader)(unsafe.Pointer(&str))
	assert.Equal(t, bsptr.Data, strptr.Data)
}

func TestToBytes(t *testing.T) {
	str := "test string"
	bs := ToBytes(str)

	// check no-copy conversion.
	bsptr := *(*reflect.SliceHeader)(unsafe.Pointer(&bs))
	strptr := *(*reflect.SliceHeader)(unsafe.Pointer(&str))
	assert.Equal(t, bsptr.Data, strptr.Data)
}

func TestIsAlnum(t *testing.T) {
	strs := []struct {
		str []byte
		ret bool
	}{
		{[]byte("az"), true},
		{[]byte("abc"), true},
		{[]byte("a0"), true},
		{[]byte("AZ"), true},
		{[]byte("Zc0"), true},
		{[]byte("aa9"), true},
		{[]byte("0ba"), true},
		{[]byte("9ca"), true},
		{[]byte("9dZ"), true},
		{[]byte("0ea"), true},
		{[]byte("-"), false},
		{[]byte("a-"), false},
		{[]byte("-z"), false},
		{[]byte("a.z"), false},
	}
	for _, test := range strs {
		assert.Equal(t, test.ret, IsAlnum(test.str))
	}
}

func TestValidSubDomain(t *testing.T) {
	strs := []struct {
		str []byte
		ret bool
	}{
		{[]byte("azz"), true},
		{[]byte("a-b"), true},
		{[]byte("a0b"), true},
		{[]byte("zb9"), true},
		{[]byte("0ab"), false},
		{[]byte("a.b"), false},
		{[]byte("aZz"), false},
		{[]byte("aZ-"), false},
	}
	for _, test := range strs {
		assert.Equal(t, test.ret, ValidSubDomain(test.str))
	}
}
