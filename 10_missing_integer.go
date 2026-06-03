package main

/*
Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].
*/

func Solution(A []int) int {

	n := len(A)
	if n == 0 {
		return 1
	}

	present := make(map[int]bool)

	for _, v := range A {
		if v > 0 {
			present[v] = true
		}
	}

	for i := 1; i <= n; i++ {
		if !present[i] {
			return i
		}
	}
	return n + 1
}

func assertEqual(a, b int) {
	if a != b {
		panic("assertion failed")
	}
}

func main() {
	println(Solution([]int{1, 2, 3}))
	assertEqual(Solution([]int{1, 2, 3}), 4)
	println(Solution([]int{-1, -3}))
	assertEqual(Solution([]int{-1, -3}), 1)
	println(Solution([]int{1, 3, 6, 4, 1, 2}))
	assertEqual(Solution([]int{1, 3, 6, 4, 1, 2}), 5)
	println(Solution([]int{1, 3, 6, 4, 1, 2, 8, 9, 10}))
	assertEqual(Solution([]int{1, 3, 6, 4, 1, 2, 8, 9, 10}), 5)
}
