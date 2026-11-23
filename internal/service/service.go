package service

import (
	"fmt"
	"strings"
)

func Convert(input string) (string, error) {
	var TextToMorse = map[string]string{
		"А": ".-", "Б": "-...", "В": ".--", "Г": "--.", "Д": "-..",
		"Е": ".", "Ё": ".", "Ж": "...-", "З": "--..", "И": "..",
		"Й": ".---", "К": "-.-", "Л": ".-..", "М": "--", "Н": "-.",
		"О": "---", "П": ".--.", "Р": ".-.", "С": "...", "Т": "-",
		"У": "..-", "Ф": "..-.", "Х": "....", "Ц": "-.-.", "Ч": "---.",
		"Ш": "----", "Щ": "--.-", "Ъ": "--.--", "Ы": "-.--", "Ь": "-..-",
		"Э": "..-..", "Ю": "..--", "Я": ".-.-",

		"0": "-----", "1": ".----", "2": "..---", "3": "...--", "4": "....-",
		"5": ".....", "6": "-....", "7": "--...", "8": "---..", "9": "----.",

		" ": "/",
	}
	var MorseToText = map[string]string{
		".-": "А", "-...": "Б", ".--": "В", "--.": "Г", "-..": "Д",
		".": "Е", "...-": "Ж", "--..": "З", "..": "И", ".---": "Й",
		"-.-": "К", ".-..": "Л", "--": "М", "-.": "Н", "---": "О",
		".--.": "П", ".-.": "Р", "...": "С", "-": "Т", "..-": "У",
		"..-.": "Ф", "....": "Х", "-.-.": "Ц", "---.": "Ч", "----": "Ш",
		"--.-": "Щ", "--.--": "Ъ", "-.--": "Ы", "-..-": "Ь", "..-..": "Э",
		"..--": "Ю", ".-.-": "Я",

		"-----": "0", ".----": "1", "..---": "2", "...--": "3", "....-": "4",
		".....": "5", "-....": "6", "--...": "7", "---..": "8", "----.": "9",

		"/": " ",
	}

	if strings.ContainsAny(input, ".-/") {
		output := ""
		parts := strings.Split(input, " ")

		for _, p := range parts {
			if p == "" {
				continue
			}

			if char, ok := MorseToText[p]; ok {
				output += char
			} else {
				return "", fmt.Errorf("неизвестный код морзе: %s", p)
			}
		}

		return output, nil
	} else {
		upper := strings.ToUpper(input)
		words := strings.Split(upper, " ")

		var morseWords []string

		for _, w := range words {
			if w == "" {
				continue
			}
			var codes []string

			for _, r := range w {
				code, ok := TextToMorse[string(r)]

				if !ok {
					return "", fmt.Errorf("неизвестный текст: %s", string(r))
				}

				codes = append(codes, code)
			}

			morseWords = append(morseWords, strings.Join(codes, " "))
		}

		return strings.Join(morseWords, " / "), nil
	}
}
