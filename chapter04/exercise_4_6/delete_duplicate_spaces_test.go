package delete_duplicate_spaces

import (
	"testing"
)

func TestDeleteDuplicateSpaces(t *testing.T) {
	var tests = []struct {
		s    []byte
		want []byte
	}{
		{[]byte(""), []byte("")},
		{[]byte("A"), []byte("A")},
		{[]byte("A B C"), []byte("A B C")},
		{[]byte("A  B  C"), []byte("A B C")},
		{[]byte("AB C  D   "), []byte("AB C D ")},
		{[]byte("   A  B CD"), []byte(" A B CD")},
		{[]byte("AB　C　　D　　　"), []byte("AB C D ")},
		{[]byte("　　　A　　B　CD"), []byte(" A B CD")},
	}

	for _, test := range tests {
		got := deleteDuplicateSpaces(test.s)
		if string(got) != string(test.want) {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
