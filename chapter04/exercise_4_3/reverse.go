package reverse

func reverse(s *[6]int) {
	for i := 0; i < len(s)/2; i++ {
		symmetry_of_i := len(s) - 1 - i
		s[i], s[symmetry_of_i] = s[symmetry_of_i], s[i]
	}
}
