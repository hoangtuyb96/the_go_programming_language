package rotate

func rotate_ints_1(ints []int) []int {
	first := ints[0]
	copy(ints, ints[1:])
	ints[len(ints)-1] = first
	return ints
}

func rotate_ints_2(ints []int) []int {
	last := ints[len(ints)-1]
	ints = append(ints, 0)
	copy(ints[1:], ints)
	ints[0] = last
	return ints[:len(ints)-1]
}
