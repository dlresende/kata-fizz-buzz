package fizzbuzz

import "testing"

func TestThatGivenANumberNItShouldPrintNumbersFrom1ToN(t *testing.T) {
	actual := FizzBuzz(2)
	expected := "1\n2\n"

	if actual != expected {
		t.Errorf("Expecting " + expected + ", but got " + actual)
	}
}
