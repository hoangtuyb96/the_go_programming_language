package shadiff

import (
	"crypto/sha256"
	"fmt"
)

// 101 & 101-1
// 101 & 100
// 100
// 100 & 011
// 000
func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}

func bitDiff(a, b []byte) int {
	fmt.Printf("%x\n", a)
	fmt.Printf("%x\n", b)
	count := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		fmt.Printf("i = %d\n", i)
		switch {
		case i >= len(a):
			fmt.Printf("b[i] = %x\n", b[i])
			count += popCount(b[i])
		case i >= len(b):
			fmt.Printf("a[i] = %x\n", a[i])
			count += popCount(a[i])
		default:
			fmt.Printf("a[i] = %x\n", a[i])
			fmt.Printf("b[i] = %x\n", b[i])	
			fmt.Printf("a[i] ^ b[i] = %x\n", a[i] ^ b[i])
			count += popCount(a[i] ^ b[i])
		}
	}
	return count
}

// ShaBitDiff returns the number of bits that are different in the SHA256
// hashes of two buffers.
func ShaBitDiff(a, b []byte) int {
	shaA := sha256.Sum256(a)
	shaB := sha256.Sum256(b)
	return bitDiff(shaA[:], shaB[:])
}
