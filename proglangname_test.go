package main

import (
	"fmt"
	"testing"
)

func TestProgName(t *testing.T) {
	str, ok := CheckProgramingLangageName("C言語")
	if !ok {
		message := fmt.Sprintf("CheckProgramingLangageName Error Input : C言語, Output : %s", str)
		t.Error(message)
	} else {
		fmt.Printf("ok : %s", str)
	}

	str, ok = CheckProgramingLangageName("c言語")
	if !ok {
		message := fmt.Sprintf("CheckProgramingLangageName Error Input : c言語, Output : %s", str)
		t.Error(message)
	} else {
		fmt.Printf("ok : %s", str)
	}
}
