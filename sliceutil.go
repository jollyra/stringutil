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

// Unique returns a slice of the unique strings in `xs`.
func Unique(xs []string) []string {
	xsUnique := make([]string, 0)
	for _, x := range xs {
		if IndexOf(xsUnique, x) == -1 {
			xsUnique = append(xsUnique, x)
		}
	}
	return xsUnique
}

// EqualSlice returns true if all xs and ys contain the same strings regardless
// of order.
func EqualSlice(xs, ys []string) bool {
	if len(xs) != len(ys) {
		return false
	}

	for _, x := range xs {
		if IndexOf(ys, x) == -1 {
			return false
		}
	}

	return true
}
