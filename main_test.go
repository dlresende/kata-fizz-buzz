package main

import "testing"

func TestThatGivenANumberItShouldPrintTheNumber(t *testing.T) {
	assertEquals(Replace(1), "1", t)
	assertEquals(Replace(2), "2", t)
}

func TestThatGivenANumberDivisibleBy3ItShouldPrintFizz(t *testing.T) {
	actual := Replace(3)

	assertEquals(actual, "Fizz", t)
}

func TestThatGivenANumberDivisibleBy5ItShouldPrintBuzz(t *testing.T) {
	actual := Replace(5)

	assertEquals(actual, "Buzz", t)
}

func TestThatGivenANumberDivisibleBy3And5ItShouldPrintFizzBuzz(t *testing.T) {
	assertEquals(Replace(15), "FizzBuzz", t)
}

func assertEquals(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("\nExpecting\n" + expected + "\nbut got\n" + actual)
	}
}

func TestThatGivenANumberNItShouldPrintNumbersFrom1ToN(t *testing.T) {
	expected := `1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
`
	assertEquals(IterateUntil(15), expected, t)
}
