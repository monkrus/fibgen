package main

func main() {
	
}

func fibo() func() int {
first, second := 0, 1
return func() int {
	ret := first
	first, second = second, first
	return ret
}



}
