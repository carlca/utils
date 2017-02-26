package matheval

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/carlca/types/stacks"
)

func UpdateInfixTokens(exp string) []string {
	// strip exp of all whitespace...
	exp = strings.Replace(exp, " ", "", -1)
	// Make a slice of rune copy of exp - this will be the source of string operations
	r := []rune(exp)
	// Make an empty slice of rune - this will be the destination of string operations
	rr := []rune{}
	// Check for any '-' operators used for negation or '+' used as a prefix...
	for i := 0; i < len(r); i++ {
		if r[i] == '-' || r[i] == '+' {
			// If the '-' or '+' is the first token or is not preceded by an operand,
			// then insert a '0' before the '-' or '+'...
			if i == 0 {
				rr = append(rr, '0')
			} else {
				rPrev := string(r[i-1])
				if isOperator(rPrev) || isParen(rPrev) {
					rr = append(rr, '0')
				}
			}
		}
		rr = append(rr, r[i])
	}
	// Break expression into comma separated slice of strings
	infixTokens := make([]string, 0)
	// Use buffer for string concatenation efficiency
	buffer := bytes.NewBufferString("")
	for i := 0; i < len(rr); i++ {
		preDelimiter := ""
		postDelimiter := ""
		if isOperator(string(rr[i])) || isParen(string(rr[i])) {
			if i > 0 {
				preDelimiter = ","
			}
			postDelimiter = ","
		}
		buffer.WriteString(preDelimiter + string(rr[i]) + postDelimiter)
	}
	// Clean up and add to infixTokens slice of string...
	tokens := strings.Replace(buffer.String(), ",,", ",", -1)
	infixTokens = strings.Split(tokens, ",")
	return infixTokens
}

func UpdatePostfixTokens(infixTokens []string) []string {
	var postfixStack stacks.Stack
	var operatorStack stacks.Stack
	// Iterate through infixTokens...
	for i := 0; i < len(infixTokens); i++ {
		token := infixTokens[i]
		// Add operand to output...
		if isOperand(token) {
			if token == "math.Pi" {
				token = strconv.FormatFloat(math.Pi, 'E', -1, 64)
			}
			if token == "math.E" {
				token = strconv.FormatFloat(math.E, 'E', -1, 64)
			}
			postfixStack.Push(token)
			continue
		}
		// Push opening paren onto stack...
		if token == "(" {
			operatorStack.Push(token)
			continue
		}
		// Add pending operators to output...
		if token == ")" {
			// Pop operators from stack until an open bracket is found
			for operatorStack.Peek() != "(" {
				if operatorStack.Empty() {
					panic("Mismatched bracket error")
				}
				postfixStack.Push(operatorStack.Pop())
			}
			// Discard the open bracket...
			operatorStack.Pop()
			continue
		}
		// Push function onto the stack...
		if isFunction(token) {
			operatorStack.Push(token)
			continue
		}
		// Add pending operators to output...
		if isOperator(token) {
			for {
				if operatorStack.Size() == 0 {
					break
				}
				if getOperatorPrecedence(operatorStack.Peek(), token) == Lower {
					break
				}
				if getOperatorPrecedence(operatorStack.Peek(), token) == SameRightAssoc {
					break
				}
				postfixStack.Push(operatorStack.Pop())
			}
			operatorStack.Push(token)
			continue
		}
	}
	// Add remaining operators to output
	for operatorStack.Size() > 0 {
		if operatorStack.Peek() == "(" {
			panic("Mismatched bracket error")
		}
		postfixStack.Push(operatorStack.Pop())
	}
	fmt.Print("Postfix Tokens: ")
	fmt.Println(postfixStack)
	postfixTokens := []string(postfixStack)
	return postfixTokens
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
	case opScore < refOpScore:
		return Higher
	case opScore > refOpScore:
		return Lower
	case opScore == 1:
		return SameRightAssoc
	default:
		return SameLeftAssoc
	}
}
