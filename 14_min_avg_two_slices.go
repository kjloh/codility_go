package main

import (
	"fmt"
)

/*
A non-empty array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P < Q < N, is called a slice of array A (notice that the slice contains at least two elements). The average of a slice (P, Q) is the sum of A[P] + A[P + 1] + ... + A[Q] divided by the length of the slice. To be precise, the average equals (A[P] + A[P + 1] + ... + A[Q]) / (Q − P + 1).

For example, array A such that:

    A[0] = 4
    A[1] = 2
    A[2] = 2
    A[3] = 5
    A[4] = 1
    A[5] = 5
    A[6] = 8
contains the following example slices:

slice (1, 2), whose average is (2 + 2) / 2 = 2;
slice (3, 4), whose average is (5 + 1) / 2 = 3;
slice (1, 4), whose average is (2 + 2 + 5 + 1) / 4 = 2.5.
The goal is to find the starting position of a slice whose average is minimal.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the starting position of the slice with the minimal average. If there is more than one slice with a minimal average, you should return the smallest starting position of such a slice.

For example, given array A such that:

    A[0] = 4
    A[1] = 2
    A[2] = 2
    A[3] = 5
    A[4] = 1
    A[5] = 5
    A[6] = 8
the function should return 1, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [2..100,000];
each element of array A is an integer within the range [−10,000..10,000].
*/

func Solution(A []int) int {
	n := len(A)

	minIndex := 0
	minSum := A[0] + A[1]
	minLen := 2

	for i := 0; i < n-1; i++ {
		// Check slice of length 2
		sum2 := A[i] + A[i+1]
		if sum2*minLen < minSum*2 {
			minSum = sum2
			minLen = 2
			minIndex = i
		}

		// Check slice of length 3
		if i < n-2 {
			sum3 := A[i] + A[i+1] + A[i+2]
			if sum3*minLen < minSum*3 {
				minSum = sum3
				minLen = 3
				minIndex = i
			}
		}
	}

	return minIndex
}

func assertEqual(expected int, actual int) {
	if expected != actual {
		panic("Actual value is incorrect")
	}
}

func main() {
	// Test case 1: Original example
	A := []int{4, 2, 2, 5, 1, 5, 6}
	fmt.Println("Test 1:", Solution(A))
	assertEqual(1, Solution(A))

	// Test case 2: Minimum at start
	B := []int{1, 2, 3, 4, 5}
	fmt.Println("Test 2:", Solution(B))
	assertEqual(0, Solution(B))

	// Test case 3: Minimum at end
	C := []int{5, 4, 3, 2, 1}
	fmt.Println("Test 3:", Solution(C))
	assertEqual(3, Solution(C))

	// Test case 4: All same numbers
	D := []int{5, 5, 5, 5, 5}
	fmt.Println("Test 4:", Solution(D))
	assertEqual(0, Solution(D))

	// Test case 5: Two elements only
	E := []int{10, 3}
	fmt.Println("Test 5:", Solution(E))
	assertEqual(0, Solution(E))

	// Test case 6: Negative numbers
	F := []int{-5, -2, -8, -1, 3}
	fmt.Println("Test 6:", Solution(F))
	assertEqual(0, Solution(F)) // (-2 + -8) / 2 = -5

	// Test case 7: Mixed positive and negative
	G := []int{100, -50, 25, -30, 10}
	fmt.Println("Test 7:", Solution(G))
	assertEqual(1, Solution(G)) // (-50 + 25) / 2 = -12.5

	// Test case 8: Large numbers
	H := []int{10000, 9999, 5000, 2000, 1000}
	fmt.Println("Test 8:", Solution(H))
	assertEqual(3, Solution(H)) // (2000 + 1000) / 2 = 1500

	// Test case 9: Duplicate minimums (should return smallest position)
	I := []int{1, 5, 2, 5, 1, 5, 10}
	fmt.Println("Test 9:", Solution(I))
	assertEqual(0, Solution(I)) // (1 + 5) / 2 and (1 + 5) / 2 at position 0 and 4, return 0

	// Test case 10: All negative
	J := []int{-1, -2, -3, -4, -5}
	fmt.Println("Test 10:", Solution(J))
	assertEqual(3, Solution(J)) // (-4 + -5) / 2 = -4.5
}
