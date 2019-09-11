package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	newText := stringConvert(text)
	fmt.Println(newText)
}

func stringConvert(text string) string {
	_, error := strconv.Atoi(string(text[0]))
	if error == nil {
		return ""
	} else {
		var newText, symbolBuf string
		var slashBuf bool
		for _, runeValue := range text {
			symbol := string(runeValue)
			i, error := strconv.Atoi(symbol)
			if error != nil {
				if symbol != `\` {
					newText = newText + symbol
				} else if slashBuf == true {
					newText = newText + symbol
				}
			} else {
				if i > 0 {
					if slashBuf != true {
						for a := 2; a <= i; a++ {
							newText = newText + symbolBuf
						}
					}
				}
				if slashBuf == true {
					newText = newText + symbol
				}

			}
			symbolBuf = symbol
			if symbol == `\` {
				if slashBuf == true {
					slashBuf = false
				} else {
					slashBuf = true
				}
			} else {
				slashBuf = false
			}
		}
		return newText
	}
}
