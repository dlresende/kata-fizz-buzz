package fizzbuzz

import "testing"

func TestThatGivenANumberNItShouldPrintNumbersFrom1ToN(t *testing.T) {
	actual := FizzBuzz(2)
	expected := "1\n2\n"

	assertEquals(actual, expected, t)
}

func TestThatGivenANumberDivisibleBy3ItShouldPrintFizz(t *testing.T) {
	actual := FizzBuzz(3)
	expected := "1\n2\nFizz\n"

	assertEquals(actual, expected, t)
}

func TestThatGivenANumberDivisibleBy5ItShouldPrintBuzz(t *testing.T) {
	actual := FizzBuzz(5)
	expected := "1\n2\nFizz\n4\nBuzz\n"

	assertEquals(actual, expected, t)
}

func assertEquals(actual string, expected string, t *testing.T) {
	if actual != expected {
		t.Errorf("\nExpecting\n" + expected + "\nbut got\n" + actual)
	}
}
