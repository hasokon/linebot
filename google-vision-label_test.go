package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFindLabels(t *testing.T) {
	f, err := os.Open("./testdata/cat.jpg")
	if err != nil {
		t.Error("File Open err : " + err.Error())
	}

	result, err := FindLabels(f)
	if err != nil {
		t.Error("Label get error : " + err.Error())
	}

	fmt.Println(result)
}
