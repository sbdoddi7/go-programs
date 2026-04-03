package main

import "fmt"

// Implement stacks
// It is linear datastructre that follows LIFO(Last if first output) meaning that last added element removed first
// It is useful in undoing operation, navigatin browser tabs, reversing strings and arrays

type stack struct {
	items []rune
}

func (s *stack) Push(item rune) {
	s.items = append(s.items, item)
}

func (s *stack) POP() {
	if s.IsEmpty() {
		fmt.Println("No elements to pop")
	}

	s.items = s.items[:len(s.items)-1]

}

func (s *stack) TOP() (rune, error) {
	if s.IsEmpty() {
		return -1, fmt.Errorf("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func (s *stack) IsEmpty() bool {
	return len(s.items) > 0
}

func checkBalance(exp string) bool {
	s := stack{}

	for _, char := range exp {
		switch char {
		case '(':
			s.Push(char)
		case '{':
			s.Push(char)
		case '[':
			s.Push(char)
		case ')':
			top, _ := s.TOP()
			if top == '(' {
				s.POP()
			} else {
				return false
			}
		case '}':
			top, _ := s.TOP()
			if top == '{' {
				s.POP()
			} else {
				return false
			}
		case ']':
			top, _ := s.TOP()
			if top == '[' {
				s.POP()
			} else {
				return false
			}
		}

	}

	if s.IsEmpty() {
		return true
	}
	return false
}

func areBracketsProperlyMatched(str string) bool {
	stacks := []rune{}

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range str {
		if char == '(' || char == '[' || char == '{' {
			stacks = append(stacks, char)
		} else if char == ')' || char == '}' || char == ']' {
			if len(stacks) == 0 {
				return false
			}

			top := stacks[len(stacks)-1]
			stacks = stacks[:len(stacks)-1]

			if top != pairs[char] {
				return false
			}
		}
	}

	if len(stacks) == 0 {
		return true
	}

	return false
}

func main() {

	str := "() {} []"
	res := checkBalance(str)

	fmt.Println("Result:", res)

	res1 := areBracketsProperlyMatched(str)

	fmt.Println("Result 1:", res1)

}
