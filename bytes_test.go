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
