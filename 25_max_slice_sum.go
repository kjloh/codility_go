package main

import "fmt"

/*
A non-empty array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P ≤ Q < N, is called a slice of array A. The sum of a slice (P, Q) is the total of A[P] + A[P+1] + ... + A[Q].

Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers, returns the maximum sum of any slice of A.

For example, given array A such that:

A[0] = 3  A[1] = 2  A[2] = -6
A[3] = 4  A[4] = 0
the function should return 5 because:

(3, 4) is a slice of A that has sum 4,
(2, 2) is a slice of A that has sum −6,
(0, 1) is a slice of A that has sum 5,
no other slice of A has sum greater than (0, 1).
Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..1,000,000];
each element of array A is an integer within the range [−1,000,000..1,000,000];
the result will be an integer within the range [−2,147,483,648..2,147,483,647].
*/

func Solution(A []int) int {
	maxSum := A[0]
	currentMax := A[0]

	for i := 1; i < len(A); i++ {
		if A[i] > (currentMax + A[i]) {
			currentMax = A[i]
		} else {
			currentMax = currentMax + A[i]
		}
		if maxSum < currentMax {
			maxSum = currentMax
		}
	}

	return maxSum
}

func main() {
	tests := []struct {
		H        []int
		expected int
	}{
		{[]int{1, 1, 1, 1, 1, 1}, 6},
		{[]int{1}, 1},
		{[]int{0}, 0},
		{[]int{-1, -1, -1, -1, -1, -1}, -1},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{10, 20, -100, 1, 2}, 30},
		{[]int{10, -1, 10, -1, 10}, 28},
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
