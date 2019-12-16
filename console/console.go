package console

import (
	"bufio"
	"os"
	"unicode"
)

var reader = bufio.NewReader(os.Stdin)

// GetChar : get one character from console
func GetChar() byte {
	line, err := reader.ReadString('\n')
	if err == nil {
		return line[0]
	}
	return ' '
}

// Ask a questions, return index of variant chosen by user
func Ask(question string, variants []rune) int {
	for {
		print(question)
		char := rune(GetChar())
		for index, variant := range variants {
			if unicode.ToLower(variant) == unicode.ToLower(char) {
				return index
			}
		}
	}
}
