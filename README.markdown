# go-roman

## Installation

    $ go get github.com/shantanubhadoria/go-roman

## Running this on the command line

    $ go-roman IV
    4

## go-roman
--
    import "github.com/shantanubhadoria/go-roman/roman"

Package roman provides methods for verifying if a number is a valid roman
numeral and finding its indo-arabic equivalent

Author: Shantanu Bhadoria

## Usage

#### func  IsRoman

```go
func IsRoman(roman string) bool
```
IsRoman checks if a number is valid roman numeral. The function returns true for
valid roman numerals and false otherwise

Rules(from wikipedia):

Numbers like V,L,D fall under non repeatable category, Max unbroken repetition
is 3 Special cases for complex digits where only certain combinations are
allowed

#### func  ToIndoArabic

```go
func ToIndoArabic(roman string) (int, error)
```
ToIndoArabic converts roman numeral to Indo-Arabic numerals. The function
returns 0, error in case of invalid roman numeral and number,nil for valid input
