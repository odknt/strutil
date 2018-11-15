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

// IsAlnum returns true if given byte slice contains only alphabets or numbers.
func IsAlnum(bs []byte) bool {
	for _, b := range bs {
		if (b < 0x41 || 0x5A < b) && (b < 0x61 || 0x7A < b) && (b < 0x30 || 0x39 < b) {
			return false
		}
	}
	return true
}

// ValidSubDomain returns true if given byte slice meet the conditions below.
// * starts with a lowercase alphabet.
// * contains only lowercase alphabets or numbers.
// * ends with a lowecase alphanets or numbers.
func ValidSubDomain(bs []byte) bool {
	if bs[0] < 0x61 || 0x7A < bs[0] {
		return false
	}
	end := len(bs) - 1
	if bs[end] == 0x2D {
		return false
	}
	for _, b := range bs[1:end] {
		if (b < 0x61 || 0x7A < b) && (b < 0x30 || 0x39 < b) && (b != 0x2D) {
			return false
		}
	}
	return true
}
