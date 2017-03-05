package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	lastNumber, _ := toInt(os.Args[1])
	output := FizzBuzz(lastNumber)
	fmt.Print(output)
}

func FizzBuzz(last int) string {
	result := ""
	for number := 1; number <= last; number++ {
		if isDivisibleBy(number, 5) {
			result += "Buzz"
		} else if isDivisibleBy(number, 3) {
			result += "Fizz"
		} else {
			result += toString(number)
		}
		result += "\n"
	}
	return result
}

func isDivisibleBy(dividend int, divisor int) bool {
	return dividend%divisor == 0
}

func toString(number int) string {
	return strconv.Itoa(number)
}

func toInt(numberAsString string) (int, error) {
	return strconv.Atoi(numberAsString)
}
