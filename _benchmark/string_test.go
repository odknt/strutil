package main

import (
	"regexp"
	"testing"

	"github.com/odknt/strutil"
)

// --------------------------------------------------------------------------------
// IsAlnum
// --------------------------------------------------------------------------------

var regexpAlnum = regexp.MustCompile("[[:alnum:]]")

func BenchmarkIsAlnum_Strutil(b *testing.B) {
	rands := randBytes(b)
	b.ResetTimer()
	for _, v := range rands {
		strutil.IsAlnum(v)
	}
}

func BenchmarkIsAlnum_Regexp(b *testing.B) {
	rands := randBytes(b)
	b.ResetTimer()
	for _, v := range rands {
		regexpAlnum.Match(v)
	}
}

// --------------------------------------------------------------------------------
// ValidSubDomain
// --------------------------------------------------------------------------------

var regexpSubDomain = regexp.MustCompile(`^[a-z][a-z0-9\-]+[a-z0-9]`)

func BenchmarkValidSubDomain_Strutil(b *testing.B) {
	rands := randBytes(b)
	b.ResetTimer()
	for _, v := range rands {
		strutil.ValidSubDomain(v)
	}
}

func BenchmarkValidSubDomain_Regexp(b *testing.B) {
	rands := randBytes(b)
	b.ResetTimer()
	for _, v := range rands {
		regexpSubDomain.Match(v)
	}
}

// --------------------------------------------------------------------------------
// ToBytes
// --------------------------------------------------------------------------------

func BenchmarkToBytes_Strutil(b *testing.B) {
	rands := randStrings(b)
	b.ResetTimer()
	for _, v := range rands {
		// ToBytes returns must set to longer lifetime variable than here.
		escapes(strutil.ToBytes(v))
	}
}

func BenchmarkToBytes_Native(b *testing.B) {
	rands := randStrings(b)
	b.ResetTimer()
	for _, v := range rands {
		// escapes returns of string() for compare to ToBytes.
		escapes([]byte(v))
	}
	b.StopTimer()
}

// --------------------------------------------------------------------------------
// FromBytes
// --------------------------------------------------------------------------------

func BenchmarkFromBytes_Strutil(b *testing.B) {
	rands := randBytes(b)
	b.ResetTimer()
	for _, v := range rands {
		// FromBytes returns must set to longer lifetime variable than here.
		escapes(strutil.FromBytes(v))
	}
	b.StopTimer()
}

func BenchmarkFromBytes_Native(b *testing.B) {
	rands := randBytes(b)
	b.ResetTimer()
	for _, v := range rands {
		// escapes returns of string() for compare to FromBytes.
		escapes(string(v))
	}
	b.StopTimer()
}

// --------------------------------------------------------------------------------
// benchmark utility
// --------------------------------------------------------------------------------

func randStrings(b *testing.B) []string {
	rands := make([]string, b.N)
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.^|"

	for i := 0; i < b.N; i++ {
		rands[i] = strutil.RandString(10, letters)
	}
	return rands
}

func randBytes(b *testing.B) [][]byte {
	rands := make([][]byte, b.N)
	letters := strutil.ToBytes("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.^|")

	for i := 0; i < b.N; i++ {
		rands[i] = strutil.RandBytes(10, letters)
	}
	return rands
}

var dummy struct {
	b bool
	x interface{}
}

// escapes moves memory to heap.
func escapes(x interface{}) {
	if dummy.b {
		dummy.x = x
	}
}
