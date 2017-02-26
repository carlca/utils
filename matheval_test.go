package matheval

import (
	"fmt"
	"testing"
)

func TestTokens(t *testing.T) {
	expression := "-3 + -4 - 7 / 12 - 3 - 34 / 12"
	infixTokens := UpdateInfixTokens(expression)
	fmt.Print("Infix Tokens: ")
	fmt.Println(infixTokens)
	fmt.Println(len(infixTokens))
	if len(infixTokens) != 17 {
		t.Error("len(infixTokens) expected 17, got", len(infixTokens))
	}
	postfixTokens := UpdatePostfixTokens(infixTokens)
	fmt.Println(postfixTokens)
}
