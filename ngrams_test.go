package textutil_test

import (
	"testing"

	"git.wark.io/lib/textutil-go"
)

func TestNgrams(t *testing.T) {
	var (
		ngrams         []string
		expectedNgrams []string
	)

	ngrams = textutil.Ngrams("343", 3)
	if len(ngrams) != 1 || ngrams[0] != "343" {
		t.Fatal("Ngrams should work with an input of length `n`")
	}

	ngrams = textutil.Ngrams("343", 4)
	if ngrams != nil {
		t.Fatal("Ngrams should return nil if `n` is greater than the input length")
	}

	ngrams = textutil.Ngrams("Hatsune Miku", 3)
	expectedNgrams = []string{
		"Hat",
		"ats",
		"tsu",
		"sun",
		"une",
		"ne ",
		"e M",
		" Mi",
		"Mik",
		"iku",
	}

	if len(ngrams) != len(expectedNgrams) {
		t.FailNow()
	}

	for i := 0; i < len(ngrams); i++ {
		if ngrams[i] != expectedNgrams[i] {
			t.FailNow()
		}
	}

	ngrams = textutil.Ngrams("Hatsune Miku", 4)
	expectedNgrams = []string{
		"Hats",
		"atsu",
		"tsun",
		"sune",
		"une ",
		"ne M",
		"e Mi",
		" Mik",
		"Miku",
	}

	if len(ngrams) != len(expectedNgrams) {
		t.FailNow()
	}

	for i := 0; i < len(ngrams); i++ {
		if ngrams[i] != expectedNgrams[i] {
			t.FailNow()
		}
	}
}
