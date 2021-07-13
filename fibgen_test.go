package main

import (
	"testing"
)

func TestFibo(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{20, 6765},
	}

	for _, test := range tests {
		if output := Fibonacci(test.input); output != test.expected {
			t.Errorf("Test Failed: {%d} inputted, {%d} expected, recieved: {%d}\n",
				test.input, test.expected, output)
		}
	}
}
