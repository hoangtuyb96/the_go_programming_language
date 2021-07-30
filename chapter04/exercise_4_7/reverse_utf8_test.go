package reverse_utf8

import "testing"

func TestReverseUTF8(t *testing.T) {
	var tests = []struct {
		s	 []byte
		want []byte
	}{
		{[]byte("hoangtuyb96 test"), []byte("tset 69byutgnaoh")},
	}

	for _, test := range tests {
		got := reverseUTF8(test.s)
		if string(got) != string(test.want) {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
