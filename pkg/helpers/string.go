package helpers

import "strings"

func Indent(n int, s string) string {
	return strings.Repeat(" ", n) + s
}
