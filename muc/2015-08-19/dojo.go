package dojo

import (
	"errors"
)

const (
	UPPER_LIMIT int = 2000
)

func Withdraw(amount int) (map[int]int, error) {
	if amount < 0 {
		return map[int]int{}, errors.New("Are you retarded?")
	}

	if amount > UPPER_LIMIT {
		return map[int]int{}, errors.New("Upper limit exceeded!")
	}

	notes := [7]int{500, 200, 100, 50, 20, 10, 5}

	rest_amount := amount

	return_map := make(map[int]int)

	for _, value := range notes {
		count := rest_amount / value
		rest_amount = rest_amount % value
		if count != 0 {
			return_map[value] = count
		}
	}

	if rest_amount != 0 {
		return map[int]int{}, errors.New("Cannot return.")
	}

	return return_map, nil
}
