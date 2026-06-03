package main

import "fmt"

/*
You are given N counters, initially set to 0, and you have two possible operations on them:

increase(X) − counter X is increased by 1,
max counter − all counters are set to the maximum value of any counter.
A non-empty array A of M integers is given. This array represents consecutive operations:

if A[K] = X, such that 1 ≤ X ≤ N, then operation K is increase(X),
if A[K] = N + 1 then operation K is max counter.
For example, given integer N = 5 and array A such that:

    A[0] = 3
    A[1] = 4
    A[2] = 4
    A[3] = 6
    A[4] = 1
    A[5] = 4
    A[6] = 4
the values of the counters after each consecutive operation will be:

    (0, 0, 1, 0, 0)
    (0, 0, 1, 1, 0)
    (0, 0, 1, 2, 0)
    (2, 2, 2, 2, 2)
    (3, 2, 2, 2, 2)
    (3, 2, 2, 3, 2)
    (3, 2, 2, 4, 2)
The goal is to calculate the value of every counter after all operations.

Write a function:

func Solution(N int, A []int) []int

that, given an integer N and a non-empty array A consisting of M integers, returns a sequence of integers representing the values of the counters.

Result array should be returned as an array of integers.

For example, given:

    A[0] = 3
    A[1] = 4
    A[2] = 4
    A[3] = 6
    A[4] = 1
    A[5] = 4
    A[6] = 4
the function should return [3, 2, 2, 4, 2], as explained above.

Write an efficient algorithm for the following assumptions:

N and M are integers within the range [1..100,000];
each element of array A is an integer within the range [1..N + 1].
*/

func Solution(N int, A []int) []int {
	counters := make([]int, N)
	lastMaxValue := 0
	lastMaxTime := -1
	currentMax := 0

	// Process all operations
	for t, v := range A {
		if v == N+1 {
			// Lazy update: record the max, don't update all counters
			lastMaxValue = currentMax
			lastMaxTime = t
		} else if v >= 1 && v <= N {
			// If this counter hasn't been updated since last max, reset it
			if lastMaxTime > -1 && counters[v-1] < lastMaxValue {
				counters[v-1] = lastMaxValue
			}
			counters[v-1]++
			if counters[v-1] > currentMax {
				currentMax = counters[v-1]
			}
		}
	}

	// Apply final max to counters not updated since last max operation
	for i := 0; i < N; i++ {
		if counters[i] < lastMaxValue {
			counters[i] = lastMaxValue
		}
	}

	return counters
}

func assertArrayEqual(a, b []int) {
	if len(a) != len(b) {
		panic("assertion failed")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			panic("assertion failed")
		}
	}
}

func main() {
	fmt.Println(Solution(1, []int{5}))
	assertArrayEqual(Solution(1, []int{5}), []int{0})
	assertArrayEqual(Solution(4, []int{5}), []int{0, 0, 0, 0})
	fmt.Println(Solution(5, []int{5}))
	assertArrayEqual(Solution(5, []int{5}), []int{0, 0, 0, 0, 1})
	fmt.Println(Solution(4, []int{5}))
	assertArrayEqual(Solution(4, []int{5}), []int{0, 0, 0, 0})
	fmt.Println(Solution(5, []int{3, 4, 4, 6, 1, 4, 4}))
	assertArrayEqual(Solution(5, []int{3, 4, 4, 6, 1, 4, 4}), []int{3, 2, 2, 4, 2})
	fmt.Println(Solution(5, []int{6, 6, 6, 6, 6, 6, 6}))
	assertArrayEqual(Solution(5, []int{6, 6, 6, 6, 6, 6, 6}), []int{0, 0, 0, 0, 0})
}
