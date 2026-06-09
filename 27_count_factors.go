package main

import (
	"fmt"
	"math"
)

/*
A positive integer D is a factor of a positive integer N if there exists an integer M such that N = D * M.

For example, 6 is a factor of 24, because M = 4 satisfies the above condition (24 = 6 * 4).

Write a function:

func Solution(N int) int

that, given a positive integer N, returns the number of its factors.

For example, given N = 24, the function should return 8, because 24 has 8 factors, namely 1, 2, 3, 4, 6, 8, 12, 24. There are no other factors of 24.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..2,147,483,647].
*/
func Solution(N int) int {

	if N <= 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	s := math.Sqrt(float64(N))

	count := 0
	for i := 2; i <= int(s); i++ {
		if N%i == 0 {
			count++
		}
	}

	count = count*2 + 2
	if s == float64(int(s)) {
		count--
	}
	return count
}

func main() {
	tests := []struct {
		H        int
		expected int
	}{
		// Edge cases
		{1, 1},
		{2, 2},
		{24, 8},
		{4, 3},
		{25, 3},
		{100, 9},
		{104729, 2},
		{2147483647, 2},
	}

	for i, test := range tests {
		result := Solution(test.H)
		if result != test.expected {
			fmt.Println("Test", i+1, "FAILED: H=", test.H, "expected=", test.expected, "got=", result)
		} else {
			fmt.Println("Test", i+1, "PASSED: H=", test.H, "result=", result)
		}
	}
}
