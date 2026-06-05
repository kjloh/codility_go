package main

import "fmt"

/*
An array A consisting of N integers is given. The dominator of array A is the value that occurs in more than half of the elements of A.

For example, consider array A such that

 A[0] = 3    A[1] = 4    A[2] =  3
 A[3] = 2    A[4] = 3    A[5] = -1
 A[6] = 3    A[7] = 3
The dominator of A is 3 because it occurs in 5 out of 8 elements of A (namely in those with indices 0, 2, 4, 6 and 7) and 5 is more than a half of 8.

Write a function

func Solution(A []int) int

that, given an array A consisting of N integers, returns index of any element of array A in which the dominator of A occurs. The function should return −1 if array A does not have a dominator.

For example, given array A such that

 A[0] = 3    A[1] = 4    A[2] =  3
 A[3] = 2    A[4] = 3    A[5] = -1
 A[6] = 3    A[7] = 3
the function may return 0, 2, 4, 6 or 7, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].
*/

func Solution(A []int) int {
	if len(A) == 0 {
		return -1
	}

	size := 0
	candidate := 0

	for _, value := range A {
		if size == 0 {
			candidate = value
			size = 1
		} else if value == candidate {
			size++
		} else {
			size--
		}
	}

	count := 0
	index := -1

	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == candidate {
			count++
			index = i
		}

	}
	if count > len(A)/2 {
		return index
	}

	return -1
}

func main() {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{3, 4, 3, 2, 3, -1, 3, 3}, 0},
		{[]int{7, 8}, -1},
		{[]int{7}, 0},
	}

	for i, test := range tests {
		result := Solution(test.input)
		if result != test.expected {
			fmt.Println("Test", i+1, "FAILED: H=", test.input, "expected=", test.expected, "got=", result)
		} else {
			fmt.Println("Test", i+1, "PASSED: H=", test.input, "result=", result)
		}
	}
}
