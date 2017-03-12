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

	number, _ := toNumber(os.Args[1])
	output := CountTo(number)
	fmt.Print(output)
}

func CountTo(lastNumber int) string {
	output := ""
	for number := 1; number <= lastNumber; number++ {
		output += Guess(number) + "\n"
	}
	return output
}

func Guess(number int) string {
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

func toNumber(numberAsString string) (int, error) {
	return strconv.Atoi(numberAsString)
}
