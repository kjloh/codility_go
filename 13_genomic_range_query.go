package main

import "fmt"

/*
A DNA sequence can be represented as a string consisting of the letters A, C, G and T, which correspond to the types of successive nucleotides in the sequence. Each nucleotide has an impact factor, which is an integer. Nucleotides of types A, C, G and T have impact factors of 1, 2, 3 and 4, respectively. You are going to answer several queries of the form: What is the minimal impact factor of nucleotides contained in a particular part of the given DNA sequence?

The DNA sequence is given as a non-empty string S = S[0]S[1]...S[N-1] consisting of N characters. There are M queries, which are given in non-empty arrays P and Q, each consisting of M integers. The K-th query (0 ≤ K < M) requires you to find the minimal impact factor of nucleotides contained in the DNA sequence between positions P[K] and Q[K] (inclusive).

For example, consider string S = CAGCCTA and arrays P, Q such that:

    P[0] = 2    Q[0] = 4
    P[1] = 5    Q[1] = 5
    P[2] = 0    Q[2] = 6
The answers to these M = 3 queries are as follows:

The part of the DNA between positions 2 and 4 contains nucleotides G and C (twice), whose impact factors are 3 and 2 respectively, so the answer is 2.
The part between positions 5 and 5 contains a single nucleotide T, whose impact factor is 4, so the answer is 4.
The part between positions 0 and 6 (the whole string) contains all nucleotides, in particular nucleotide A whose impact factor is 1, so the answer is 1.
Write a function:

func Solution(S string, P []int, Q []int) []int

that, given a non-empty string S consisting of N characters and two non-empty arrays P and Q consisting of M integers, returns an array consisting of M integers specifying the consecutive answers to all queries.

Result array should be returned as an array of integers.

For example, given the string S = CAGCCTA and arrays P, Q such that:

    P[0] = 2    Q[0] = 4
    P[1] = 5    Q[1] = 5
    P[2] = 0    Q[2] = 6
the function should return the values [2, 4, 1], as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
M is an integer within the range [1..50,000];
each element of arrays P and Q is an integer within the range [0..N - 1];
P[K] ≤ Q[K], where 0 ≤ K < M;
string S consists only of upper-case English letters A, C, G, T.
*/

func Solution(S string, P []int, Q []int) []int {

	n := len(S)

	// Create prefix count arrays for each nucleotide
	prefixA := make([]int, n+1)
	prefixC := make([]int, n+1)
	prefixG := make([]int, n+1)
	prefixT := make([]int, n+1)

	for i := 0; i < n; i++ {
		prefixA[i+1] = prefixA[i]
		prefixC[i+1] = prefixC[i]
		prefixG[i+1] = prefixG[i]
		prefixT[i+1] = prefixT[i]

		switch S[i] {
		case 'A':
			prefixA[i+1]++
		case 'C':
			prefixC[i+1]++
		case 'G':
			prefixG[i+1]++
		case 'T':
			prefixT[i+1]++
		}
	}

	if len(P) < 1 || len(Q) < 1 {
		return []int{}
	}

	m := len(P)
	impact := make([]int, m)
	for i := 0; i < len(P); i++ {
		p, q := P[i], Q[i]
		if prefixA[q+1]-prefixA[p] > 0 {
			impact[i] = 1
		} else if prefixC[q+1]-prefixC[p] > 0 {
			impact[i] = 2
		} else if prefixG[q+1]-prefixG[p] > 0 {
			impact[i] = 3
		} else {
			impact[i] = 4
		}
	}

	return impact
}

func assertArrayEqual(expected []int, actual []int) {
	if len(expected) != len(actual) {
		panic("assertEqual failed")
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			panic("assertEqual failed")
		}
	}
}

func main() {
	fmt.Println(Solution("CAGCCTA", []int{}, []int{}))
	assertArrayEqual([]int{}, Solution("CAGCCTA", []int{}, []int{}))
	println("Test 1 passed: CAGCCTA with queries [] to []")

	// Test case 1: Original example
	fmt.Println(Solution("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}))
	assertArrayEqual([]int{2, 4, 1}, Solution("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}))
	println("Test 1 passed: CAGCCTA with queries [2,5,0] to [4,5,6]")

	// Test case 2: Single character queries
	fmt.Println(Solution("CAGCCTA", []int{0, 1, 2, 3}, []int{0, 1, 2, 3}))
	assertArrayEqual([]int{2, 1, 3, 2}, Solution("CAGCCTA", []int{0, 1, 2, 3}, []int{0, 1, 2, 3}))
	println("Test 2 passed: Single character queries")

	// Test case 3: All same nucleotide
	fmt.Println(Solution("AAAA", []int{0, 1, 2}, []int{1, 2, 3}))
	assertArrayEqual([]int{1, 1, 1}, Solution("AAAA", []int{0, 1, 2}, []int{1, 2, 3}))
	println("Test 3 passed: All same nucleotide (A)")

	// Test case 4: Query with T at start and A at end
	assertArrayEqual([]int{1, 4}, Solution("ATGC", []int{0, 1}, []int{3, 1}))
	println("Test 4 passed: Mixed nucleotides")

	// Test case 5: Entire string query
	assertArrayEqual([]int{1}, Solution("TGCA", []int{0}, []int{3}))
	println("Test 5 passed: Entire string contains all nucleotides")

	// Test case 6: Multiple queries on same string
	assertArrayEqual([]int{4, 3, 2, 1}, Solution("TGCA", []int{0, 1, 2, 3}, []int{0, 1, 2, 3}))
	println("Test 6 passed: Individual character queries")
}
