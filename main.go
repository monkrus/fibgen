package main

import (
	"fmt"
    "log"
    "net/http"
	"strings"
	"testing"
)

func TestConstruct(t *testing.T) {
	var tests = []struct {
		input  int
		expected int
		received int
	}{
		// input  0, received 0
		//  input 1, received 1
		//  input 10, received 55
		//  input -2, received "value needs to be equal or greater than zero"
	
	

	for _, test := range tests {
		if output := construct(test.input, test.received); output != test.expected {
			t.Errorf("Test Failed: {%s} inputted, {%s} expected, recieved: {%s}\n",
				fmt.Sprint(test.input, test.received), test.expected, output)
		}
	}

}

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return (fibonacci(n-1) + fibonacci(n-2))
	}
}

func main() {
        // Set routing rules
        http.HandleFunc("/", Tmp)

        //Use the default DefaultServeMux.
        err := http.ListenAndServe(":8080", nil)
        if err != nil {
                log.Fatal(err)
        }
}

func Tmp(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "test")
}


	var n, i, j int
	j = 0

	fmt.Print("Enter the number of terms: ")
	fmt.Scanln(&n)

	fmt.Println("Fibonacci series: ")
	for i = 1; i <= n; i++ {
		fmt.Print(fibonacci(j), ",")
		j++
	}
	fmt.Println()
}

/* option
func fibo() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}

}
func main() {

	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
*/
