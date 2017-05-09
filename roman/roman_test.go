package roman

import (
	"fmt"
	"testing"
)

type romanMap struct {
	roman      string
	indoarabic int
}

var nonRomanNumerals = []string{
	"IJKL",
	"IIII",
	"IVI",
	"VIV",
	"IXIII",
}

var romanNumerals = []romanMap{
	romanMap{"I", 1},
	romanMap{"V", 5},
	romanMap{"X", 10},
	romanMap{"L", 50},
	romanMap{"C", 100},
	romanMap{"D", 500},
	romanMap{"M", 1000},
	romanMap{"XIII", 13},
	romanMap{"XIV", 14},
}

func TestIsRoman(t *testing.T) {
	for _, nr := range nonRomanNumerals {
		assertEqual(t, IsRoman(nr), false, "Invalid roman numeral returned true for IsRoman")
	}
	for _, r := range romanNumerals {
		assertEqual(t, IsRoman(r.roman), true, "Valid roman numeral returned false for IsRoman")
	}
}

func TestToIndoArabic(t *testing.T) {
	for _, r := range romanNumerals {
		result, _ := ToIndoArabic(r.roman)
		assertEqual(t, result, r.indoarabic, "Valid roman numeral returned false for IsRoman")
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	message = fmt.Sprintf("%v != %v \n%s", a, b, message)
	t.Fatal(message)
}
