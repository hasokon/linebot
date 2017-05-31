package main

import (
	"strings"

	"github.com/hasokon/yahooapi"
)

var countNgo = 0

// CheckMoukoben is a function that detects any moukoben and return correct reply messages
func CheckMoukoben(text string) (string, bool) {
	result, err := yahooapi.MorphologicalAnalysys(text)
	if err != nil {
		return "", false
	}

	isNgo := false
	wordlist := result.Ma.Wordlist.Wordlist
	for _, word := range wordlist {
		if strings.HasPrefix(word.Surface, "ンゴ") {
			isNgo = true
			break
		}
	}

	switch {
	case strings.Contains(text, "334"):
		return "なんでや！阪神関係ないやろ！", true
	case isNgo:
		countNgo++
		if countNgo >= 15 {
			countNgo = 0
			return "ンゴンゴうるさいンゴォ！！\n何回「はえ～」言えばええんや！", true
		}
		return "はえ〜", true
	case strings.HasSuffix(text, "や！"):
		return "(*^◯^*)", true
	}

	return "", false
}
