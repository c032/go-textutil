package textutil

import (
	"strings"
	"unicode/utf8"
)

var parens = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'⁽': '⁾',
	'₍': '₎',
	'⌈': '⌉',
	'⌊': '⌋',
	'〈': '〉',
	'❨': '❩',
	'❪': '❫',
	'❬': '❭',
	'❮': '❯',
	'❰': '❱',
	'❲': '❳',
	'❴': '❵',
	'⟦': '⟧',
	'⟨': '⟩',
	'⟪': '⟫',
	'⦃': '⦄',
	'⦅': '⦆',
	'⦗': '⦘',
	'⸢': '⸣',
	'⸤': '⸥',
	'⸨': '⸩',
	'〈': '〉',
	'《': '》',
	'「': '」',
	'『': '』',
	'【': '】',
	'〔': '〕',
	'〖': '〗',
	'〘': '〙',
	'〚': '〛',
	'﹁': '﹂',
	'﹃': '﹄',
	'﹇': '﹈',
	'﹙': '﹚',
	'﹛': '﹜',
	'﹝': '﹞',
	'（': '）',
	'［': '］',
	'｛': '｝',
	'｟': '｠',
	'｢': '｣',
}

// SplitParens separates chunks inside parens from chunks outside them.
func SplitParens(text string) []string {
	if len(text) == 0 {
		return []string{text}
	}

	var (
		ok bool

		opening rune
		closing rune

		start int = 0
		end   int = len(text)
	)

	positions := make([][2]int, 0)

	// Read the text rune by rune.
	for i, char := range text {

		// Check whether we are inside parens.
		if opening != 0 {
			if char == opening {
				end = i

				positions = append(positions, [2]int{start, end})

				opening = 0
				start = i
				end = len(text)
			} else if char == closing {
				closing = 0
				end = i + utf8.RuneLen(char)

				positions = append(positions, [2]int{start, end})

				opening = 0
				start = end
				end = len(text)
			}

		} else if closing, ok = parens[char]; ok {
			// If `char` is an "opening paren", then `closing` will
			// be its "closing" version.

			end = i
			if start != end {
				positions = append(positions, [2]int{start, end})
			}

			start = i
			end = len(text)

			opening = char
		}
	}

	if start != end {
		positions = append(positions, [2]int{start, end})
	}

	groups := make([]string, len(positions))
	for i, position := range positions {
		start, end = position[0], position[1]
		groups[i] = text[start:end]
	}

	return groups
}

// IsParen returns true if r is one of the supported parens.
func IsParen(r rune) bool {
	for opening, closing := range parens {
		if r == opening || r == closing {
			return true
		}
	}

	return false
}

// StripParens returns text without any enclosing parens.
func StripParens(text string) string {
	for ropening, rclosing := range parens {
		opening := string(ropening)
		closing := string(rclosing)
		if strings.HasPrefix(text, opening) && strings.HasSuffix(text, closing) {
			text = strings.TrimPrefix(text, opening)
			text = strings.TrimSuffix(text, closing)
			break
		}
	}

	return text
}
