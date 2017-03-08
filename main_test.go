package main

import "testing"

func TestThatGivenANumberItShouldPrintTheNumber(t *testing.T) {
	assertEquals(Replace(1), "1", t)
	assertEquals(Replace(2), "2", t)
}

func TestThatGivenANumberNItShouldPrintNumbersFrom1ToN(t *testing.T) {
	actual := IterateUntil(5)
	expected := "1\n2\nFizz\n4\nBuzz\n"

	assertEquals(actual, expected, t)
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
