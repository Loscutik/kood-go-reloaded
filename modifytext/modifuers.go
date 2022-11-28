/*
function used for converting string in modifystr.go
*/
package modifytext

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

/*
converts hex number (pesented by string) in decimal number (pesented by string)
*/
func convertHexToDec(nbr string) (string, error) {
	if v, _ := regexp.MatchString(`[^0-9AaBbCcDdEeFf]`, nbr); v {
		return "", fmt.Errorf("wrong hex number: %s", nbr)
	}
	return convertBase(strings.ToUpper(nbr), "0123456789ABCDEF", "0123456789"), nil
}

/*
converts bin number (pesented by string) in decimal number (pesented by string)
*/
func convertBinToDec(nbr string) (string, error) {
	if v, _ := regexp.MatchString(`[^01]`, nbr); v {
		return "", fmt.Errorf("wrong bin number: %s", nbr)
	}
	return convertBase(strings.ToUpper(nbr), "01", "0123456789"), nil
}

/*
converts a number given by string  `nbr` from base `baseFrom` to `baseTo`
*/
func convertBase(nbr, baseFrom, baseTo string) string {
	lenBaseFrom := len([]rune(baseFrom))
	lenBaseTo := len([]rune(baseTo))
	tmp := 0
	mapBaseFrom := make(map[rune]int, lenBaseFrom)
	arrBaseTo := []rune(baseTo)
	for i, letter := range []rune(baseFrom) {
		mapBaseFrom[letter] = i
	}

	// nbr to decimal number
	for _, char := range nbr {
		tmp = lenBaseFrom*tmp + mapBaseFrom[char]
	}

	//  convert decimal number tmp  to "number" in BaseTo
	if tmp == 0 {
		return "0"
	}
	res := ""
	for tmp != 0 {
		res = string(arrBaseTo[tmp%lenBaseTo]) + res
		tmp /= lenBaseTo
	}
	return res
}

/*
capitalizes first letters of each words in the given string
*/
func Capitalize(s string) string {
	words := strings.Split(s, " ")
	res := make([]rune, 0, len(s))
	// capitalise the first letter of each word and add the changed word to res
	for i, w := range words {
		letters := ([]rune(w))
		letters[0] = unicode.ToUpper(letters[0])
		if i == 0 {
			res = append(res, letters...)
		} else {
			res = append(res, ' ')
			res = append(res, letters...)
		}
	}
	return string(res)
}
