package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

var prevIsDigit = false

func Unpack(str string) (string, error) {
	// Place your code here.
	runes := []rune(str)
	var result strings.Builder
	for i, rune := range runes {
		if i == 0 && unicode.IsDigit(rune) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(rune) && prevIsDigit {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(rune) {
			prevIsDigit = true
			continue
		}

		if i+1 == len(runes) {
			result.WriteRune(rune)
			return result.String(), nil
		}

		if unicode.IsDigit(runes[i+1]) {
			countRepeat := runes[i+1] - 48
			prevIsDigit = false
			result.WriteString(strings.Repeat(string(rune), int(countRepeat)))
		} else {
			prevIsDigit = false
			result.WriteRune(rune)
		}
	}

	return result.String(), nil
}
