package strutil

import (
	"unsafe"
)

// FromBytes provides no-copy conversion from []byte to string.
func FromBytes(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

// ToBytes provides no-copy conversion from string to []byte.
func ToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}
