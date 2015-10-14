package dojo

import (
	. "strings"
)

func Problem(input string, offensiveWord string) bool {

	var working_input string = ""

	for _, c := range input {
		if Contains(offensiveWord, string(c)) {
			working_input += string(c)
		}
	}

	return working_input == offensiveWord

}
