package main

import (
	"fmt"
)

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	} else {
		return (Fibonacci(n-1) + Fibonacci(n-2))
	}
}

func main() {

	var n, i, j int
	j = 0

	fmt.Print("Enter the number of terms: ")
	fmt.Scanln(&n)

	fmt.Println("Fibonacci series: ")
	for i = 1; i <= n; i++ {
		fmt.Print(Fibonacci(j), ",")
		j++
	}
	fmt.Println()
}

/* option 2
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
