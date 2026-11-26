package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var convertMorseToText = morse.NewConverter(
	morse.DefaultMorse,
	morse.WithCharSeparator(" "),
	morse.WithWordSeparator("   "),
	morse.WithTrailingSeparator(false),
	morse.WithHandler(morse.IgnoreHandler),
)

var convertInputToMorse = morse.NewConverter(
	morse.DefaultMorse,
	morse.WithCharSeparator(" "),
	morse.WithLowercaseHandling(true),
	morse.WithTrailingSeparator(false),
	morse.WithHandler(morse.IgnoreHandler),
)

const wordSeparator = " / "
const morseWordSeparator = "   "

func Convert(input string) (string, error) {
	if isMorse(input) {
		return decodeFromMorse(input)
	}

	return encodeToMorse(input)
}

func isMorse(s string) bool {
	return strings.ContainsFunc(s, func(r rune) bool { return r == '.' || r == '-' }) &&
		!strings.ContainsFunc(s, func(r rune) bool {
			switch r {
			case '.', '-', '/', ' ', '\t', '\n', '\r':
				return false
			}
			return true
		})
}

func encodeToMorse(text string) (string, error) {
	upper := strings.ToUpper(text)

	for _, r := range upper {
		if r == ' ' {
			continue
		}

		if _, ok := morse.DefaultMorse[r]; !ok {
			return "", fmt.Errorf("неизвестный текст: %s", string(r))
		}
	}

	words := strings.Split(upper, " ")
	encoded := make([]string, 0, len(words))
	for _, w := range words {
		if w == "" {
			continue
		}
		encoded = append(encoded, convertInputToMorse.ToMorse(w))
	}

	return strings.Join(encoded, wordSeparator), nil
}

func decodeFromMorse(morseStr string) (string, error) {
	normalized := strings.ReplaceAll(morseStr, wordSeparator, morseWordSeparator)
	normalized = strings.ReplaceAll(normalized, "/", morseWordSeparator)

	for _, w := range strings.Split(normalized, morseWordSeparator) {
		for _, code := range strings.Fields(w) {
			if morse.MorseToRune(code) == 0 {
				return "", fmt.Errorf("неизвестный код морзе: %s", code)
			}
		}
	}

	return convertMorseToText.ToText(normalized), nil
}
