package strutil

import (
	"fmt"
)

// Def returns def if str is empty.
func Def(str, def string) string {
	if len(str) == 0 {
		return def
	}
	return str
}

// DefExpr returns def if expr is false else str.
func DefExpr(expr bool, str, def string) string {
	if expr {
		return str
	}
	return def
}

// DefStringer returns def if fmt.Stringer is nil.
func DefStringer(str fmt.Stringer, def string) string {
	if str == nil {
		return def
	}
	return Def(str.String(), def)
}

// DefExprStringer returns def if expr is false or Stringer is nil else call String().
func DefExprStringer(expr bool, str fmt.Stringer, def string) string {
	if expr {
		return DefStringer(str, def)
	}
	return def
}
