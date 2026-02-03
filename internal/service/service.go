package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorseCode(s string) bool {
	morseChars := ".- "

	for _, r := range s {
		if !strings.ContainsRune(morseChars, r) {
			return false
		}
	}
	return true
}

// Функцию описать
func Convert(s string) (string, error) {
	if s == "" {
		return "", errors.New("Строка пустая")
	}
	if isMorseCode(s) {
		return morse.ToText(s), nil
	} else {
		return morse.ToMorse(s), nil
	}
}
