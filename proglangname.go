package main

import (
	"fmt"
	"strings"

	"github.com/hasokon/yahooapi"
)

// LangageNum is the number of Programing langages
const LangageNum = 29

var proglangages = [LangageNum]string{
	"C",
	"Lisp",
	"C++",
	"C#",
	"Java",
	"JavaScript",
	"Python",
	"Ruby",
	"BASIC",
	"COBOL",
	"Dart",
	"FORTRAN",
	"Groovy",
	"Haskell",
	"D",
	"Kotlin",
	"MATLAB",
	"Nim",
	"Pascal",
	"Go",
	"Perl",
	"PHP",
	"R言語",
	"Rust",
	"Scala",
	"Smalltalk",
	"Swift",
	"TeX",
	"VerilogHDL",
}

// CheckProgramingLangageName is a function that detect Programing Langage Name
func CheckProgramingLangageName(text string) (string, bool) {
	result, err := yahooapi.MorphologicalAnalysys(text)
	if err != nil {
		return err.Error(), false
	}

	progNames := make([]string, 0)
	wordlist := result.Ma.Wordlist.Wordlist
	fmt.Println(wordlist)
	for _, word := range wordlist {
		upperword := strings.ToUpper(word.Surface)
		for _, progname := range proglangages {
			upperprogname := strings.ToUpper(progname)
			fmt.Printf("%s %s", upperprogname, upperword)
			if upperword == upperprogname {
				progNames = append(progNames, progname)
			}
		}
	}

	for _, progname := range progNames {
		switch progname {
		case "Go":
			return "我が名を呼ぶ声がするンゴ", true
		case "Lisp":
			return "Lispはきもいンゴ", true
		case "Dart":
			return "Dartはワイの親友や！", true
		case "C":
			return "Cパイセンには頭が上がらないンゴ", true
		default:
			return fmt.Sprintf("%sはくたばるがいいンゴ!!", progname), true
		}
	}
	return "", false
}
