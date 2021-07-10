package main

import (
	"fmt"
    "log"
    "net/http"
)

func fibonacci(n int) int {
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
