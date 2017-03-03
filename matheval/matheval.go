package MathEval

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
	for _, token := range infixTokens {
		// Add operand to output...
		if isOperand(token) {
			if token == "Pi" {
				token = strconv.FormatFloat(math.Pi, 'E', -1, 64)
			}
			if token == "E" {
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
				if getOperatorPrecedence(operatorStack.Peek(), token) == Higher {
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

func UpdatePostfixValue(postfixTokens []string) float64 {
	var operandStack stacks.Stack
	for _, token := range postfixTokens {
		if isOperand(token) {
			operandStack.Push(token)
			continue
		}
		var evalResult float64
		if isOperator(token) {
			rightOperand := getOperandValue(operandStack.Pop())
			leftOperand := getOperandValue(operandStack.Pop())
			switch token {
			case "+":
				evalResult = leftOperand + rightOperand
			case "-":
				evalResult = leftOperand - rightOperand
			case "*":
				evalResult = leftOperand * rightOperand
			case "/":
				evalResult = leftOperand / rightOperand
			case "^":
				evalResult = math.Pow(leftOperand, rightOperand)
			}
			operandStack.Push(strconv.FormatFloat(evalResult, 'E', -1, 64))
			continue
		}
		if isFunction(token) {
			operand := getOperandValue(operandStack.Pop())
			switch token {
			case "cos":
				evalResult = math.Cos(getTrigOperand(operand))
			case "sin":
				evalResult = math.Sin(getTrigOperand(operand))
			case "tan":
				evalResult = math.Tan(getTrigOperand(operand))
			case "exp":
				evalResult = math.Exp(operand)
			case "log":
				evalResult = math.Log(operand)
			case "log10":
				evalResult = math.Log10(operand)
			case "sqrt":
				evalResult = math.Sqrt(operand)
			}
			operandStack.Push(strconv.FormatFloat(evalResult, 'E', -1, 64))
			continue
		}
	}
	if operandStack.Size() != 1 {
		panic("Operand stack count != 1")
	}
	return getOperandValue(operandStack.Pop())
}
