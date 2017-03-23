package matheval

import (
	"strings"
	"strconv"
)

func init() {
	SetTrigUnits(Degrees)
}

func isOperator(s string) bool {
	return strings.Contains("+-*/^'", s)
}

func isParen(s string) bool {
	return strings.Contains("()", s)
}

func isFunction(s string) bool {
	fn := " " + strings.ToLower(s) + " "
	return strings.Contains(" cos sin tan exp log log10 sqrt ", fn)
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

func getOperandValue(op string) float64 {
	val, _ := strconv.ParseFloat(op, 64)
	return val
}

const (
	Higher = 1 + iota
	Lower
	SameLeftAssoc
	SameRightAssoc
)

func getScore(op string) int {
	op = strings.ToLower(op) + " "
	switch {
	case strings.Contains("^ ", op):
		return 1
	case strings.Contains("* / ", op):
		return 2
	case strings.Contains("+ - ", op):
		return 3
	case strings.Contains("( ) ", op):
		return 4
	}
	return 0
}

func getOperatorPrecedence(op string, refOp string) int {
	opScore := getScore(op)
	refOpScore := getScore(refOp)
	switch {
	case opScore > refOpScore:
		return Higher
	case opScore < refOpScore:
		return Lower
	case opScore == 1:
		return SameRightAssoc
	default:
		return SameLeftAssoc
	}
}

var TrigUnits int

const (
	Degrees = 1 + iota
	Radians
)

const (
	DegToRad = 0.017453292519943295769236907684886127134428718885417 // N[Pi/180, 50]
)

func SetTrigUnits(units int) {
	if units == Degrees || units == Radians {
		TrigUnits = units
	}
}

func getTrigOperand(op float64) float64 {
	if TrigUnits == Degrees {
		return op * DegToRad
	} else {
		return op
	}
}