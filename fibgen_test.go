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
	//func setupRouter starts the router
	router := setupRouter()
	// start recording the  http response
	w := httptest.NewRecorder()
	// start new  specific request (error value is skipped)
	req, _ := http.NewRequest("GET", "/ping", nil)
	//  replicate the handler utulizing the response recorder and new request
	router.ServeHTTP(w, req)
	// assert the equality of request and response
	// recorder HTTP response code
	assert.Equal(t, 200, w.Code)
	// response recorded in `Body` buffer
	assert.Equal(t, "pong", w.Body.String())
}

// testing incorrect values (negative, non-number and big number)
func TestNegNum(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get?number=-3", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "number should be a positive number, but it is negative", w.Body.String())
}

/*func TestNonNumber(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get?number=($)", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "strconv.Atoi: parsing \"($)\": invalid syntax", w.Body.String())

}*/

func TestBigNum(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get?number=9223372036854775807", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "number should be a number, but it is exceeds the maximum value if int", w.Body.String())
}

//testing correct values (0, 10)

func TestZeroNum(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get?number=0", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "number should be zero, but it is not", w.Body.String())
}

func TestTenNum(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get?number=10", nil)
	router.ServeHTTP(w, req)
	assert.NotEqual(t, http.StatusOK, w.Code)
	assert.Equal(t, "number should be ten, but it is not", w.Body.String())
}
