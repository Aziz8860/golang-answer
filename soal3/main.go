package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isBalanced(s string) string {
	stack := []rune{}
	
	bracketPairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	
	for _, char := range s {
		if char == ' ' {
			continue
		}
		
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			
			if len(stack) == 0 {
				return "NO"
			}
			
			// Pop the last opening bracket from stack
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			
			// Check if it matches the current closing bracket
			if last != bracketPairs[char] {
				return "NO"
			}
		}
	}
	
	if len(stack) == 0 {
		return "YES"
	}
	
	return "NO"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Enter a string with brackets:")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	
	result := isBalanced(input)
	fmt.Println("Output:")
	fmt.Println(result)
}
