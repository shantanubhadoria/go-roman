// Package roman provides methods for verifying if a number is a valid roman numeral and finding its indoarabic equivalent
//
// Author: Shantanu Bhadoria
package roman

import (
	// "fmt"
	"errors"
	"regexp"
	"strings"
)

var digitsMap = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

// IsRoman checks if a number is valid roman numeral.
// The function returns true for valid roman numerals and false otherwise
//
// Rules(from wikipedia):
//
// Numbers like V,L,D fall under non repeatable category,
// Max unbroken repetition is 3
// Special cases for complex digits where only certain combinations are allowed
func IsRoman(roman string) bool {
	if roman == "" {
		return false
	}

	romanb := []byte(strings.ToUpper(roman))
	check, _ := regexp.Match("^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", romanb)
	return check
}

// ToIndoArabic converts roman numeral to Indo-Arabic numerals.
// The function returns 0, error in case of invalid roman numeral and number,nil for valid input
func ToIndoArabic(roman string) (int, error) {
	if !IsRoman(roman) {
		return 0, errors.New("non roman number provided")
	}

	// Keep track of the last parsed digit to look for compound digits like IV when prevDigit < currentDigit
	previousValue := 1000 // Assign maxval among all digits to start with
	romanb := []byte(strings.ToUpper(roman))
	totalValue := 0
	for _, rDigit := range romanb {
		value := digitsMap[string(rDigit)]
		if previousValue < value {
			totalValue -= previousValue                     // Undo addition from previous step
			totalValue = totalValue + value - previousValue // Compound digit addition
		} else {
			totalValue += value
		}
		previousValue = value
	}

	return totalValue, nil
}
