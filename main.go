package fizzbuzz

import "strconv"

func FizzBuzz(last int) string {
	result := ""
	for number := 1; number <= last; number++ {
		result += strconv.Itoa(number) + "\n"
	}
	return result
}
