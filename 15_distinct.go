package main

import "fmt"

/*
rite a function

func Solution(A []int) int

that, given an array A consisting of N integers, returns the number of distinct values in array A.

For example, given array A consisting of six elements such that:

 A[0] = 2    A[1] = 1    A[2] = 1
 A[3] = 2    A[4] = 3    A[5] = 1
the function should return 3, because there are 3 distinct values appearing in array A, namely 1, 2 and 3.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].
*/

func Solution(A []int) int {
	distinctValues := make(map[int]int)

	for _, v := range A {
		distinctValues[v]++
	}

	return len(distinctValues)
}

func assertEqual(expected int, actual int) {
	if expected != actual {
		panic("Actual value is incorrect")
	}
}

func main() {
	// Test 1
	A := []int{2, 1, 1, 2, 3, 1}
	fmt.Println("Test 1: ", A)
	println(Solution(A))
	assertEqual(3, Solution(A))
}
