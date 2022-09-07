package textutil_test

import (
	"testing"

	"github.com/c032/go-textutil"
)

func TestLex1(t *testing.T) {
	tokens := textutil.Lex1("3.14 Lorem ipsum. .. ... ....")
	expectedTokens := [][2]string{
		[2]string{"number", "3"},
		[2]string{"period", "."},
		[2]string{"number", "14"},
		[2]string{"space", " "},
		[2]string{"letters", "Lorem"},
		[2]string{"space", " "},
		[2]string{"letters", "ipsum"},
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
