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
