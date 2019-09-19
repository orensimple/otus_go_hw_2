package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	newText := stringConvert(text)
	fmt.Println(newText)
}

func stringConvert(text string) string {
	if unicode.IsDigit(rune(text[0])) {
		return ""
	}
	var symbolBuf string
	var newText strings.Builder
	var slashBuf bool
	for _, runeValue := range text {
		symbol := string(runeValue)
		if !unicode.IsDigit(runeValue) {
			if symbol != `\` || slashBuf {
				newText.WriteString(symbol)
			}
		} else {
			count, _ := strconv.Atoi(symbol)
			if count > 0 && !slashBuf {
				newText.WriteString(strings.Repeat(symbolBuf, count-1))
			}
			if slashBuf {
				newText.WriteString(symbol)
			}
		}
		symbolBuf = symbol
		if symbol == `\` && !slashBuf {
			slashBuf = true
		} else {
			slashBuf = false
		}
	}
	return newText.String()
}
