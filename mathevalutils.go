package matheval

import (
	"strings"
)

func isOperator(s string) bool {
	return strings.Contains("+-*/^'", s)
}

func isParen(s string) bool {
	return strings.Contains("()", s)
}

func isFunction(s string) bool {
	fn := " " + strings.ToLower(s) + " "
	return strings.Contains(" cos exp int sin frac round trunc tan ln log10 sqrt ", fn)
}

func isOperand(s string) bool {
	result := false
	if len(s) > 0 {
		result = true
		if isOperator(s) || isParen(s) || isFunction(s) {
			result = false
		}
	}
	return result
}
