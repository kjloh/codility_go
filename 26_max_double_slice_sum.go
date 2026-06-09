package main

import (
	"fmt"
)

/*
A non-empty array A consisting of N integers is given.

A triplet (X, Y, Z), such that 0 ≤ X < Y < Z < N, is called a double slice.

The sum of double slice (X, Y, Z) is the total of A[X + 1] + A[X + 2] + ... + A[Y − 1] + A[Y + 1] + A[Y + 2] + ... + A[Z − 1].

For example, array A such that:

    A[0] = 3
    A[1] = 2
    A[2] = 6
    A[3] = -1
    A[4] = 4
    A[5] = 5
    A[6] = -1
    A[7] = 2
contains the following example double slices:

double slice (0, 3, 6), sum is 2 + 6 + 4 + 5 = 17,
double slice (0, 3, 7), sum is 2 + 6 + 4 + 5 − 1 = 16,
double slice (3, 4, 5), sum is 0.
The goal is to find the maximal sum of any double slice.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the maximal sum of any double slice.

For example, given:

    A[0] = 3
    A[1] = 2
    A[2] = 6
    A[3] = -1
    A[4] = 4
    A[5] = 5
    A[6] = -1
    A[7] = 2
the function should return 17, because no double slice of array A has a sum of greater than 17.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [3..100,000];
each element of array A is an integer within the range [−10,000..10,000].
*/

func Solution(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}

	maxEndingAt := make([]int, n)
	for i := 1; i < n-2; i++ {
		tmp := maxEndingAt[i-1] + A[i]
		if tmp > 0 {
			maxEndingAt[i] = tmp
		} else {
			maxEndingAt[i] = 0
		}
	}

	maxStartingAt := make([]int, n)
	for i := n - 2; i > 1; i-- {
		tmp := maxStartingAt[i+1] + A[i]
		if tmp > 0 {
			maxStartingAt[i] = tmp
		} else {
			maxStartingAt[i] = 0
		}
	}

	maxSum := 0
	for i := 1; i < n-1; i++ {
		tmp := maxEndingAt[i-1] + maxStartingAt[i+1]
		if tmp > maxSum {
			maxSum = tmp
		}
	}

	return maxSum
}

func main() {
	tests := []struct {
		H        []int
		expected int
	}{
		// Edge cases
		{[]int{3, 2, 6, -1, 4, 5, -1, 2}, 17},
		{[]int{0, 1, 0, 1, 0}, 2},
		{[]int{10, 20, 30}, 0},
		{[]int{-1, -2, -3}, 0},
		{[]int{-1, -2, -3, -4, -5}, 0},
		{[]int{1, 1, 1, 1, 1}, 2},
		{[]int{1, 10, 1, 10, 1}, 20},
		{[]int{1, 2, 3, 4, 5, 6}, 12},
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
