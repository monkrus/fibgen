package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
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
	// where "if input is not equal to expected output, the error is created"
	for _, test := range tests {
		if output := Fibonacci(test.input); output != test.expected {
			t.Errorf("Test Failed: {%d} inputted, {%d} expected, recieved: {%d}\n",
				test.input, test.expected, output)
		}
	}
}

// test to confirm the server has been started
func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
func TestNum(t *testing.T) {
	// we create a struct with inputted and expected values
	 tests := []struct {
		input    int
		expected int
		&n int 
	}{
		// indicate abovementioned values
		{&n, 0},
		{&n, 1},
		{&n, 9223372036854775807},
	}
	// run range loop  for all tests
	// where "if input is not equal to expected output, the error is created"
	for _, test := range tests {
		if output := Fibonacci(test.input); output != test.expected {
			t.Errorf("Test Failed: {%d} inputted, {%d} expected, recieved: {%d}\n",
				test.input, test.expected, output)
		}
	}
}
