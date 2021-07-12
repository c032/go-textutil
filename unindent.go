package textutil

import (
	"strings"
)

// Unindent removes all extra indentation to match the first non-whitespace
// line.
func Unindent(input string) string {
	lines := strings.Split(input, "\n")
	for len(lines) > 0 && strings.TrimSpace(lines[0]) == "" {
		lines = lines[1:]
	}
	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}

	var prefix string

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		var prefixChar rune
		if line[0] == ' ' {
			prefixChar = ' '
		} else if line[0] == '\t' {
			prefixChar = '\t'
		} else {
			return input
		}

		for _, c := range line {
			if c == prefixChar {
				prefix += string(prefixChar)
			}
		}

		break
	}

	if prefix == "" {
		return input
	}

	var output []string
	for _, line := range lines {
		noSpace := strings.TrimSpace(line)
		offset := strings.Index(line, noSpace)

		line = line[:offset] + noSpace
		line = strings.TrimPrefix(line, prefix)

		output = append(output, line)
	}

	return strings.Join(output, "\n")
}
