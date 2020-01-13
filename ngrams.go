package textutil

// Ngrams return a list of all n-size tokens (including overlaps) found
// in the input.
func Ngrams(input string, n int) []string {

	if len(input) < n {
		return nil
	}

	howManyNgrams := len(input) - n + 1
	ngrams := make([]string, 0, howManyNgrams)

	var (
		cursor int
		chunk  string
	)

	for cursor = 0; cursor+n < len(input)+1; cursor++ {
		chunk = input[cursor : cursor+n]
		ngrams = append(ngrams, chunk)
	}

	return ngrams
}
