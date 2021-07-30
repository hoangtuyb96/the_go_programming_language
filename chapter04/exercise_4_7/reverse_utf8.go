package reverse_utf8

import "unicode/utf8"

func rev(b []byte) {
	for i := 0; i < len(b)/2; i++ {
		symmetry_of_i := len(b) - 1 - i
		b[i], b[symmetry_of_i] = b[symmetry_of_i], b[i]
	}
}

func reverseUTF8(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		rev(b[i : i+size])
		i += size
	}
	rev(b)
	return b
}
