package hamming

import "errors"

// Distance computes the Hamming difference between
// two DNA strands
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("DNA strands are not equal")
	}

	count := 0

	for pos := range a {
		if a[pos] != b[pos] {
			count++
		}
	}

	return count, nil
}
