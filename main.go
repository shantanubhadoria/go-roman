package main

import (
	"fmt"
	"os"

	"github.com/shantanubhadoria/go-roman/roman"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprint(os.Stderr, fmt.Sprintf("\n\n\t Usage: %s <roman numeral>\n\n", args[0]))
	} else {
		value, err := roman.ToIndoArabic(args[1])
		if err != nil {
			fmt.Fprint(os.Stderr, "Invalid roman numeral\n")
		} else {
			fmt.Println(value)
		}
	}
}
