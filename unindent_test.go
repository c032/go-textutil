package textutil_test

import (
	"testing"

	"github.com/c032/go-textutil"
)

func TestUnindent(t *testing.T) {
	tt := map[string]string{}

	var (
		input string
		want  string
	)

	input = `textutil`
	want = `textutil`

	tt[input] = want

	input = `
		textutil
	`
	want = `textutil`

	tt[input] = want

	input = `
		First.

		Second.
	`
	want = "First.\n\nSecond."

	tt[input] = want

	for input, want = range tt {
		got := textutil.Unindent(input)

		if got != want {
			t.Errorf("Unindent(%q) = %q; want %q", input, got, want)
		}
	}
}
