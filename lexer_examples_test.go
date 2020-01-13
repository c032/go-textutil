package textutil_test

import (
	"testing"

	"git.wark.io/lib/textutil-go"
)

func TestLex1(t *testing.T) {
	tokens := textutil.Lex1("343 Guilty Spark. .. ... ....")
	expectedTokens := [][2]string{
		[2]string{"number", "343"},
		[2]string{"space", " "},
		[2]string{"letters", "Guilty"},
		[2]string{"space", " "},
		[2]string{"letters", "Spark"},
		[2]string{"period", "."},

		[2]string{"space", " "},
		[2]string{"period", "."},
		[2]string{"period", "."},

		[2]string{"space", " "},
		[2]string{"ellipsis", "..."},

		[2]string{"space", " "},
		[2]string{"period", "."},
		[2]string{"period", "."},
		[2]string{"period", "."},
		[2]string{"period", "."},
	}

	if len(tokens) != len(expectedTokens) {
		t.Fatalf("expecting %d tokens, got %d", len(expectedTokens), len(tokens))
	}

	for i, token := range tokens {
		expectedToken := expectedTokens[i]
		if len(token) != len(expectedToken) {
			t.Error("tokens have not the same length")
			continue
		}
		if token[0] != expectedToken[0] {
			t.Errorf("expecting %#v, got %#v", expectedToken[0], token[0])
		}
		if token[1] != expectedToken[1] {
			t.Errorf("expecting %#v, got %#v", expectedToken[1], token[1])
		}
	}
}
