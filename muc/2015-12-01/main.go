package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func Calculate(input string) int {
	oper, digits := Transform(Tokenize(input))

	fmt.Println(digits)
	fmt.Println(oper)

	for i, _ := range digits {
		if i < len(oper) {
			switch oper[i] {
			case "+":
				if i < len(digits)-1 {
					digits[i+1] = digits[i] + digits[i+1]
				}
			case "-":
				digits[i+1] = digits[i] - digits[i+1]
			}
		}
	}

	return digits[len(digits)-1]
}

func Tokenize(input string) []string {
	r, err := regexp.Compile("(\\d+|\\+|\\-|\\*|/|\\(|\\))")

	if err != nil {
		panic("error")
	}

	digits := r.FindAllString(input, -1)

	return digits
}

func Transform(tokens []string) ([]string, []int) {
	operations := make([]string, 0)
	digits := make([]int, 0)

	for _, v := range tokens {
		i, err := strconv.Atoi(v)

		if err != nil {
			operations = append(operations, v)
		} else {
			digits = append(digits, i)
		}
	}

	return operations, digits
}
