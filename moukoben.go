package main

import "strings"

// CheckMoukoben is a function that detects any moukoben and return correct reply messages
func CheckMoukoben(text string) (string, bool) {
	switch {
	case strings.Contains(text, "334"):
		return "なんでや！阪神関係ないやろ！", true
	case strings.HasSuffix(text, "ンゴ"):
		return "はえ〜", true
	case strings.HasSuffix(text, "や！"):
		return "(*^◯^*)", true
	}

	return "", false
}
