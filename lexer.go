package textutil

import (
	"regexp"
)

var wordRegexp = regexp.MustCompile(`[a-zA-ZáíúéóÁÍÚÉÓ]`)
var numberRegexp = regexp.MustCompile(`[0-9]`)

// Lex1 parses and tokenizes the given text an returns a slice
// of 2-index array. The first index contains the token type,
// and the second, the token.
func Lex1(text string) [][2]string {
	if len(text) == 0 {
		return nil
	}

	out := make([][2]string, 0)
	cursor := 0

	var push func(tokenType, token string)

	push = func(tokenType, token string) {
		if tokenType == "period" && len(token) > 1 {

			// Convert 3 periods into an ellipsis.
			if len(token) == 3 {
				push("ellipsis", token)
			} else {
				for i := 0; i < len(token); i++ {
					push(tokenType, ".")
				}
			}
			return
		}

		out = append(out, [2]string{tokenType, token})
	}

	token := ""
	tokenType := ""

	char := ""
	charType := ""

	for cursor < len(text) {
		char = string(text[cursor])
		cursor++

		if wordRegexp.MatchString(char) {
			charType = "letters"
		} else if numberRegexp.MatchString(char) {
			charType = "number"
		} else if char == " " {
			charType = "space"
		} else if char == "." {
			charType = "period"
		} else if char == "," {
			charType = "comma"
		} else if char == ";" {
			charType = "semicolon"
		} else if char == ":" {
			charType = "colon"
		} else if char == "'" {
			charType = "apostrophe"
		} else if char == "¿" {
			charType = "question_start"
		} else if char == "?" {
			charType = "question_end"
		} else if char == "¡" {
			charType = "exclamation_start"
		} else if char == "!" {
			charType = "exclamation_end"
		} else if char == "-" {
			charType = "hyphen"
		} else {
			charType = "other"
		}

		if len(token) == 0 {

			token = char
			tokenType = charType

			if tokenType != "letters" && tokenType != "number" && tokenType != "period" {
				push(tokenType, token)
				tokenType = ""
				token = ""
			}

		} else if tokenType == "letters" && charType == "letters" {
			token += char
		} else if tokenType == "number" && charType == "number" {
			token += char
		} else if tokenType == "period" && charType == "period" {
			token += char
		} else {
			push(tokenType, token)
			token = char
			tokenType = charType
		}
	}

	if len(token) > 0 {
		push(tokenType, token)
	}

	return out
}
