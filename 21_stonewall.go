package main

/*
You are going to build a stone wall. The wall should be straight and N meters long, and its thickness should be constant; however, it should have different heights in different places. The height of the wall is specified by an array H of N positive integers. H[I] is the height of the wall from I to I+1 meters to the right of its left end. In particular, H[0] is the height of the wall's left end and H[N−1] is the height of the wall's right end.

The wall should be built of cuboid stone blocks (that is, all sides of such blocks are rectangular). Your task is to compute the minimum number of blocks needed to build the wall.

Write a function:

func Solution(H []int) int

that, given an array H of N positive integers specifying the height of the wall, returns the minimum number of blocks needed to build it.

For example, given array H containing N = 9 integers:

  H[0] = 8    H[1] = 8    H[2] = 5
  H[3] = 7    H[4] = 9    H[5] = 8
  H[6] = 7    H[7] = 4    H[8] = 8
the function should return 7. The figure shows one possible arrangement of seven blocks.



Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array H is an integer within the range [1..1,000,000,000].
*/
import "fmt"

func Solution(H []int) int {
	if len(H) == 0 {
		return 0
	}

	stack := make([]int, 0)
	blocks := 0
	for _, height := range H {
		// Pop all blocks taller than current height
		for len(stack) > 0 && height < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		// If height is different from top of stack, add new block
		if len(stack) == 0 || height > stack[len(stack)-1] {
			blocks++
			stack = append(stack, height)
		}
	}
	return blocks
}

func main() {
	tests := []struct {
		H        []int
		expected int
	}{
		// Edge cases
		{[]int{}, 0},
		{[]int{5}, 1},

		// All same heights
		{[]int{5, 5, 5, 5}, 1},
		{[]int{1, 1, 1, 1, 1}, 1},

		// Strictly increasing
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 2, 3}, 3},

		// Strictly decreasing
		{[]int{5, 4, 3, 2, 1}, 1},
		{[]int{3, 2, 1}, 1},

		// Simple up and down
		{[]int{1, 3, 1}, 2},
		{[]int{3, 1, 3}, 2},

		// Original example
		{[]int{8, 8, 5, 7, 9, 8, 7, 4, 8}, 7},

		// More complex patterns
		{[]int{1, 1, 2, 2, 1, 1}, 2},
		{[]int{5, 3, 5, 3, 5}, 3},
		{[]int{1, 2, 1, 2, 1}, 3},

		// Large jump up then down
		{[]int{1, 10, 1}, 2},
		{[]int{1, 100, 50, 1}, 3},

		// Saw-tooth pattern
		{[]int{1, 5, 1, 5, 1, 5}, 3},

		// Pyramid shape
		{[]int{1, 2, 3, 2, 1}, 3},
		{[]int{1, 2, 3, 4, 3, 2, 1}, 4},
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
