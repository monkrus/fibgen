package main

import "fmt"

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
