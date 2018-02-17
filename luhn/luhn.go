package luhn

import (
	"strconv"
	"unicode"
)

// Valid determine whether or not number is valid per the Luhn formula.
func Valid(input string) bool {
	code := parseCode(input)
	codeLength := len(code)
	if codeLength < 2 {
		return false
	}

	sum := 0
	for i := codeLength - 1; i >= 0; i-- {
		if (codeLength-i)%2 == 0 {
			sum += double(code[i])
		} else {
			sum += code[i]
		}
	}

	return sum%10 == 0
}

func parseCode(input string) []int {
	var code []int
	for _, s := range input {
		if unicode.IsSpace(s) {
			continue
		}
		if !unicode.IsDigit(s) {
			return []int{}
		}
		digit, _ := strconv.Atoi(string(s))
		code = append(code, digit)
	}

	return code
}

func double(i int) int {
	result := i * 2
	if result > 9 {
		result -= 9
	}
	return result
}
