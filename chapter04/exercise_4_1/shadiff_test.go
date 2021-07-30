package shadiff

import (
	"testing"
)

func TestBitDiff(t *testing.T) {
	tests := []struct {
		a, b []byte
		want int
	}{
		{[]byte{13}, []byte{9}, 121},
		{[]byte{1, 2, 3}, []byte{4, 5, 6}, 135},
	}
	for _, test := range tests {
		got := ShaBitDiff(test.a, test.b)
		if got != test.want {
			t.Errorf("bitDiff(%v, %v), got %d, want %d",
				test.a, test.b, got, test.want)
		}
	}
}
