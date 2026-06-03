package main

/*
Write a function:

func Solution(A int, B int, K int) int

that, given three integers A, B and K, returns the number of integers within the range [A..B] that are divisible by K, i.e.:

{ i : A ≤ i ≤ B, i mod K = 0 }

For example, for A = 6, B = 11 and K = 2, your function should return 3, because there are three numbers divisible by 2 within the range [6..11], namely 6, 8 and 10.

Write an efficient algorithm for the following assumptions:

A and B are integers within the range [0..2,000,000,000];
K is an integer within the range [1..2,000,000,000];
A ≤ B.
*/

func Solution(A int, B int, K int) int {
	if A > B || K == 0 {
		return 0
	}

	count := B/K - A/K
	if A%K == 0 {
		count++
	}
	return count
}

func assertEqual(expected, actual int) {
	if expected != actual {
		panic("assertEqual failed")
	}
}

func main() {
	println(Solution(6, 11, 2)) // 3
	assertEqual(3, Solution(6, 11, 2))
	println(Solution(7, 12, 2)) // 3
	assertEqual(3, Solution(7, 12, 2))
	println(Solution(6, 5, 2))
	assertEqual(0, Solution(6, 5, 2))
	println(Solution(6, 12, 2)) // 3
	assertEqual(4, Solution(6, 12, 2))
	println(Solution(1, 12, 2)) // 6
	assertEqual(6, Solution(1, 12, 2))
	println(Solution(6, 11, 0))
	assertEqual(0, Solution(6, 11, 0))
}
