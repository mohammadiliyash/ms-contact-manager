package utils

import (
	"bytes"
	"strings"
)

// AddDashes ...
func AddDashes(str string) string {
	return insertRune(str, '-')
}

// insertRune...
func insertRune(str string, runeToInsert rune) string {
	buf := bytes.NewBufferString("")
	for i, v := range str {
		if i > 0 && v >= 'A' && v <= 'Z' {
			buf.WriteRune(runeToInsert)
		}
		buf.WriteRune(v)
	}

	return strings.ToLower(buf.String())
}
