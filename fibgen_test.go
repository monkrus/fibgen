package main

import (
	"testing"
)

// function TestFibo is using functionality of the testing package
func TestFibo(t *testing.T) {
	// we create a struct with inputted and expected values
	var tests = []struct {
		input    int
		expected int
	}{
		// indicate abovementioned values
		{0, 0},
		{1, 1},
		{20, 6765},
	}
	// run range loop  for all tests
	// where "if input is not equalt to expected output, the error is created"
	for _, test := range tests {
		if output := Fibonacci(test.input); output != test.expected {
			t.Errorf("Test Failed: {%d} inputted, {%d} expected, recieved: {%d}\n",
				test.input, test.expected, output)
		}
	}
}
