package strutil

// Def returns def if str is empty.
func Def(str string, def string) string {
	if len(str) == 0 {
		return def
	}
	return str
}
