package strutil

import (
	"fmt"
)

// Def returns def if str is empty.
func Def(str string, def string) string {
	if len(str) == 0 {
		return def
	}
	return str
}

// DefStringer returns def if stringer is nil or empty.
func DefStringer(str fmt.Stringer, def string) string {
	if str == nil {
		return def
	}
	return Def(str.String(), def)
}
