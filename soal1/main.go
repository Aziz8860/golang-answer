package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculates the nth term of A000124 sequence (lazy caterer sequence)
// Formula: a(n) = n(n+1)/2 + 1 for n >= 1
func calculateA000124(n int) int {
	if n == 0 {
		return 1
	}
	return (n*(n+1))/2 + 1
}

func generateSequence(n int) []int {
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		sequence[i] = calculateA000124(i)
	}
	return sequence
}

func main() {
	fmt.Println("A000124 Sequence Generator")
	fmt.Print("Enter the number: ")
	
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	n, err := strconv.Atoi(input)
	if err != nil || n <= 0 {
		fmt.Println("Please enter a valid positive integer")
		return
	}
	
	sequence := generateSequence(n)
	
	// Format the output like this: 1-2-4-7-11-16-22
	result := make([]string, len(sequence))
	for i, num := range sequence {
		result[i] = strconv.Itoa(num)
	}
	
	// output result
	fmt.Println(strings.Join(result, "-"))
}
