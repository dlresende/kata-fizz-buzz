package fizzbuzz

import "strconv"

func FizzBuzz(last int) string {
	var result string
	result = ""
	for number := 1; number <= last; number++ {
		if number%5 == 0 {
			result += "Buzz"
		} else if number%3 == 0 {
			result += "Fizz"
		} else {
			result += strconv.Itoa(number)
		}
		result += "\n"
	}
	return result
}
