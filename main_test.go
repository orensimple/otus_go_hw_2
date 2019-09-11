package main

import (
	"testing"
)

const testResult1 = `aaaabccddddde`
const testResult2 = `abcd`
const testResult3 = ``
const testResult4 = `qwe45`
const testResult5 = `qwe44444`
const testResult6 = `qwe\\\\\`

func TestStringConvert(t *testing.T) {
	result := stringConvert("a4bc2d5e")
	if result != testResult1 {
		t.Errorf("%s %s", result, testResult1)
	}
	result = stringConvert("abcd")
	if result != testResult2 {
		t.Errorf("%s %s", result, testResult2)
	}
	result = stringConvert("45")
	if result != testResult3 {
		t.Errorf("%s %s", result, testResult4)
	}
	result = stringConvert(`qwe\4\5`)
	if result != testResult4 {
		t.Errorf("%s %s", result, testResult4)
	}
	result = stringConvert(`qwe\45`)
	if result != testResult5 {
		t.Errorf("%s %s", result, testResult5)
	}
	result = stringConvert(`qwe\\5`)
	if result != testResult6 {
		t.Errorf("%s %s", result, testResult6)
	}

}
