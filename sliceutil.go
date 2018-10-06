package stringutil

// IndexOf returns the index of the first occurence of `string`.
func IndexOf(strings []string, target string) int {
	for i, s := range strings {
		if s == target {
			return i
		}
	}
	return -1
}

// Permutations populates `dest` with all possible permutations of the input
// slice `xs`.
func Permutations(dest *[][]string, xs []string, unfixed int) {
	if len(xs)-1 == unfixed {
		*dest = append(*dest, xs)
	}

	for i := unfixed; i < len(xs); i++ {
		ys := CopySlice(xs)
		ys[unfixed], ys[i] = ys[i], ys[unfixed]
		Permutations(dest, ys, unfixed+1)
	}
}

// CopySlice returns a copy of the input slice `xs`.
func CopySlice(xs []string) []string {
	ys := make([]string, len(xs))
	for i := range xs {
		ys[i] = xs[i]
	}
	return ys
}
