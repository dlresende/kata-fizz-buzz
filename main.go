package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Print("Wrong number of arguments.\nUsage: kata-fizz-buzz <number>")
		os.Exit(1)
	}

	lastNumber, _ := toInt(os.Args[1])
	output := IterateUntil(lastNumber)
	fmt.Print(output)
}

func IterateUntil(lastNumber int) string {
	output := ""
	for number := 1; number <= lastNumber; number++ {
		output += Replace(number) + "\n"
	}
	return output
}

func Replace(number int) string {
	if isDivisibleBy(number, 3*5) {
		return "FizzBuzz"
	}

	if isDivisibleBy(number, 5) {
		return "Buzz"
	}

	if isDivisibleBy(number, 3) {
		return "Fizz"
	}

	return toString(number)
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
