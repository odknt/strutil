package strutil

import (
	"math/rand"
	"time"
)

var randSrc = rand.NewSource(time.Now().UnixNano())

const (
	randBits = 63

	defLetters  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	defNeedBits = 6
	defMask     = 1<<defNeedBits - 1
	defMaxIdx   = randBits / defNeedBits
)

// RandBytes returns randomized byte slice for given n length.
// If letters is not given then uses alphabets and numbers.
func RandBytes(n int, letters []byte) []byte {
	b := make([]byte, n)

	var needBits, maxIdx uint64
	var mask int64

	idxLen := len(letters)
	if idxLen == 0 {
		needBits = defNeedBits
		mask = defMask
		maxIdx = defMaxIdx
		idxLen = len(defLetters)
		letters = ToBytes(defLetters)
	} else {
		// calculate need bits length of max index.
		var i uint64 = 1
		for ; i < 64; i++ {
			if (idxLen >> i) == 0 {
				needBits = i
				break
			}
		}
		mask = 1<<needBits - 1
		maxIdx = randBits / needBits
	}
	cache, remain := randSrc.Int63(), maxIdx
	for i := 0; i < n; remain-- {
		if cache == 0 || remain == 0 {
			cache, remain = randSrc.Int63(), maxIdx
		}
		idx := int(cache & mask)
		if idx < idxLen {
			b[i] = letters[idx]
			i++
		}
		cache >>= needBits
	}
	return b
}

// RandString returns random string for given n length.
// If letters is not given then uses alphabets and numbers.
func RandString(n int, letters string) string {
	return FromBytes(RandBytes(n, ToBytes(letters)))
}
