package delete_duplicate_spaces

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

func deleteDuplicateSpaces(s []byte) []byte {
	var buf bytes.Buffer
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRune(s[i:])
		if unicode.IsSpace(r) {
			nextrune, _ := utf8.DecodeRune(s[i+size:])
			if !unicode.IsSpace(nextrune) {
				buf.WriteRune(' ')
			}
		} else {
			buf.WriteRune(r)
		}
		i += size
	}
	return buf.Bytes()
}
