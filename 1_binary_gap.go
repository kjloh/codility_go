package main

/*
A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.

For example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.

Write a function:

class Solution { public int solution(int N); }

that, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.

For example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5. Given N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..2,147,483,647].
*/

func Solution(N int) int {
	// Skip trailing zeros
	for N > 0 && N&1 == 0 {
		N >>= 1
	}

	var maxGap, gap int
	for N > 0 {
		if N&1 == 0 {
			gap++
		} else {
			if gap > maxGap {
				maxGap = gap
			}
			gap = 0
		}
		N >>= 1
	}

	return maxGap
}

func assert(ok bool) {
	if !ok {
		panic("assertion failed")
	}
}

func main() {
	println(Solution(9))
	assert(Solution(9) == 2)
	println(Solution(529))
	assert(Solution(529) == 4)
	println(Solution(20))
	assert(Solution(20) == 1)
	println(Solution(15))
	assert(Solution(15) == 0)
	println(Solution(32))
	assert(Solution(32) == 0)
	println(Solution(1041))
	assert(Solution(1041) == 5)
}
