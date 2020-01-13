package textutil_test

import (
	"testing"

	"git.wark.io/lib/textutil-go"
)

func TestSplitParens(t *testing.T) {
	samples := map[string][]string{
		"[HorribleSubs] Himouto! Umaru-chan - 01 [720p].mkv": []string{
			"[HorribleSubs]",
			" Himouto! Umaru-chan - 01 ",
			"[720p]",
			".mkv",
		},
		"Himouto! Umaru-chan【01】【GB】【720P】【MP4】": []string{
			"Himouto! Umaru-chan",
			"【01】",
			"【GB】",
			"【720P】",
			"【MP4】",
		},
	}

	for sample, expected := range samples {
		result := textutil.SplitParens(sample)
		if len(result) != len(expected) {
			t.Fatalf("expecting %d chunks; got %d", len(expected), len(result))
		}
		for i := 0; i < len(expected); i++ {
			if result[i] != expected[i] {
				t.Fatalf("expecting %#v, got %#v", expected[i], result[i])
			}
		}
	}
}

func TestIsParen(t *testing.T) {
	sample := []rune{'[', ']', '【', '】'}
	for _, p := range sample {
		if !textutil.IsParen(p) {
			t.Fatalf("%#v should be a paren", p)
		}
	}
}

func TestStripParens(t *testing.T) {
	sample := "【GB】"
	expected := "GB"

	if result := textutil.StripParens(sample); result != expected {
		t.Fatalf("expecting %#v; got %#v", expected, result)
	}
}
