package main

/*
A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:

S is empty;
S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.
For example, the string "{[()()]}" is properly nested but "([)()]" is not.

Write a function:

func Solution(S string) int

that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.

For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..200,000];
string S is made only of the following characters: '(', '{', '[', ']', '}' and/or ')'.
*/

func Solution(S string) int {
	matches := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]rune, 0, len(S))

	for _, ch := range S {
		if opening, isClosing := matches[ch]; isClosing {
			// It's a closing bracket
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return 0
			}
			stack = stack[:len(stack)-1]
		} else if ch == '(' || ch == '{' || ch == '[' {
			// It's an opening bracket
			stack = append(stack, ch)
		} else {
			// Invalid character
			return 0
		}
	}

	// Stack should be empty if properly nested
	if len(stack) == 0 {
		return 1
	}
	return 0
}

func assertEqual(expected int, actual int) {
	if expected != actual {
		panic("Test failed")
	}
}

func main() {
	tests := []struct {
		input    string
		expected int
	}{
		{"{[()()]}", 1},
		{"{[(()])}", 0},
		{"", 1},
		{"()", 1},
		{"[]", 1},
		{"{}", 1},
		{"([{}])", 1},
		{"([)]", 0},
		{"{[}]", 0},
		{"(", 0},
		{")", 0},
		{"((()))", 1},
		{"[[[[]]]]", 1},
	}

	for i, test := range tests {
		result := Solution(test.input)
		if result != test.expected {
			println("Test", i+1, "FAILED: input=", test.input, "expected=", test.expected, "got=", result)
		} else {
			println("Test", i+1, "PASSED: input=", test.input)
		}
	}
}
